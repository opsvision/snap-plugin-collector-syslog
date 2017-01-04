/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 OpsVision Solutions

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package syslog

import (
	"encoding/json"
	"fmt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"gopkg.in/mcuadros/go-syslog.v2"
	"log"
	"os"
	"time"
)

// Constants
const (
	NS_VENDOR = "opsvision"
	NS_PLUGIN = "syslog"
	VERSION   = 1
	PORT      = 1514
	BUFSIZE   = 1024
	FILENAME  = "/tmp/syslog-collector.log"
)

// Types
type SyslogCollector struct {
	initialized bool
	incoming    syslog.LogPartsChannel
	logCounter  uint64
	logger      *log.Logger
}

type SyslogMessage struct {
	app_name        string
	client          string
	facility        string
	hostname        string
	message         string
	msg_id          string
	priority        string
	proc_id         string
	severity        string
	structured_data string
	timestamp       string
	tls_peer        string
	version         string
}

// Constructor
func New() *SyslogCollector {
	return new(SyslogCollector)
}

// Initialization
func (p *SyslogCollector) init(cfg plugin.Config) error {
	if p.initialized {
		return nil
	}

	// Setup logging (optional/debugging)
	if file, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		p.logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
		p.logger.Println("Snap Plugin Collector for Syslog, Version ", VERSION)
		p.logger.Println("Logging system online")
	}

	// Initialize the log counter
	p.logCounter = 1

	// Initialize the buffered channel
	bufsize := BUFSIZE // set initial to default
	if value, err := cfg.GetInt("bufsize"); err == nil {
		bufsize = int(value)
	}
	p.incoming = make(syslog.LogPartsChannel, bufsize)
	p.logger.Println("Using Bufsize: ", bufsize)

	// Start listening for Syslog messages
	go func() {
		port := PORT // set initial to default
		if value, err := cfg.GetInt("port"); err == nil {
			port = int(value)
		}

		network := fmt.Sprintf("0.0.0.0:%v", port)
		handler := syslog.NewChannelHandler(p.incoming)
		server := syslog.NewServer()
		server.SetFormat(syslog.RFC5424)
		server.SetHandler(handler)
		server.ListenUDP(network)
		server.Boot()

		p.logger.Println("Listening on ", network)

		server.Wait()
	}()

	p.initialized = true
	return nil
}

/*
	CollectMetrics collects metrics.

	CollectMetrics() will be called by Snap when a task that collects one of the
	metrics returned from this plugins GetMetricTypes() is started. The input will
	include a slice of all the metric types being collected.

	The output is the collected metrics as plugin.Metric and an error.
*/
func (p *SyslogCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	// Check init
	if len(mts) > 0 {
		if err := p.init(mts[0].Config); err != nil {
			return nil, err
		}
	} else {
		return mts, nil
	}

	metrics := []plugin.Metric{}
	time := time.Now()

	/*
		Non-Blocking read from the syslog channel results in NULL entries in the
		log file every time this function is called.
	*/
	select {
	case logParts := <-p.incoming:
		if len(logParts["hostname"].(string)) == 0 {
			break
		}

		// Iterate over the requested metrics
		for _, mt := range mts {
			metricName := mt.Namespace[len(mt.Namespace)-1].Value

			switch metricName {
			case "counter":
				setCounterMetric(p.logCounter, mt, &metrics, time)

				// Increment the counter
				p.logCounter += 1

			case "testing":
				setStaticMetric("testing", mt, &metrics, time)

			default:
				setLogPartMetrics(logParts, mt, &metrics, time)
			}
		}

	default:
		// NOOP - there is no messages to process
	}

	return metrics, nil
}

/*
	GetMetricTypes returns metric types.

	GetMetricTypes() will be called when your plugin is loaded in order to populate
	the metric catalog (where snaps stores all available metrics).

	Config info is passed in. This config information would come from global config
	snap settings.

	The metrics returned will be advertised to users who list all the metrics and
	will become targetable by tasks.
*/
func (p *SyslogCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	// Counter metric
	// namespace: /NS_VENDOR/NS_PLUGIN/counter
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "counter"),
		Version:   VERSION,
	})

	// Message metric
	// namespace: /NS_VENDOR/NS_PLUGIN/event/*/message
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "event").
			AddDynamicElement("source", "the source hostname or IP address").
			AddStaticElement("message"),
		Version: VERSION,
	})

	// Summary metric
	// namespace: /NS_VENDOR/NS_PLUGIN/event/*/summary
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "event").
			AddDynamicElement("source", "the source hostname or IP address").
			AddStaticElement("summary"),
		Version: VERSION,
	})

	return metrics, nil
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.

	A config policy is how users can provide configuration info to plugin. Here
	you define what sorts of config info your plugin needs and/or requires.
*/
func (p *SyslogCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	// The UDP and TCP port the syslog server should listen
	policy.AddNewIntRule([]string{NS_VENDOR, NS_PLUGIN},
		"port",
		false,
		plugin.SetMaxInt(65535),
		plugin.SetMinInt(1))

	// The size of the buffered channel
	policy.AddNewIntRule([]string{NS_VENDOR, NS_PLUGIN},
		"bufsize",
		false)

	return *policy, nil
}

/*
	setCounterMetric is used to create the counter metric
*/
func setCounterMetric(counter uint64, mt plugin.Metric, metrics *[]plugin.Metric, time time.Time) {
	metric := plugin.Metric{
		Timestamp: time,
		Namespace: mt.Namespace,
		Data:      counter,
		Tags:      mt.Tags,
		Config:    mt.Config,
		Version:   VERSION,
	}

	// Store our metric
	*metrics = append(*metrics, metric)
}

/*
	setStaticMetric is used to set a static value for a metric
*/
func setStaticMetric(data string, mt plugin.Metric, metrics *[]plugin.Metric, time time.Time) {
	metric := plugin.Metric{
		Timestamp: time,
		Namespace: mt.Namespace,
		Data:      data,
		Tags:      mt.Tags,
		Config:    mt.Config,
		Version:   VERSION,
	}

	// Store our metric
	*metrics = append(*metrics, metric)
}

/*
	setLogPartMetrics is used to set the syslog fields
*/
func setLogPartMetrics(logParts map[string]interface{}, mt plugin.Metric, metrics *[]plugin.Metric, time time.Time) {
	// Extract hostname and message payload
	hostname := fmt.Sprintf("%v", logParts["hostname"])
	summary := fmt.Sprintf("%v", logParts["message"])
	message, _ := json.Marshal(logParts)

	// Setup the namespace (/opsvision/syslog/event/[source]/summary
	ns := plugin.NewNamespace(NS_VENDOR, NS_PLUGIN,
		"event", hostname, "summary")

	// Create the summary metric
	metric := plugin.Metric{
		Timestamp: time,
		Namespace: ns,
		Data:      summary,
		Tags:      mt.Tags,
		Config:    mt.Config,
		Version:   VERSION,
	}

	// Store our metric
	*metrics = append(*metrics, metric)

	// Setup the namespace (/opsvision/syslog/event/[source]/message
	ns = plugin.NewNamespace(NS_VENDOR, NS_PLUGIN,
		"event", hostname, "message")

	// Create the summary metric
	metric = plugin.Metric{
		Timestamp: time,
		Namespace: ns,
		Data:      string(message),
		Tags:      mt.Tags,
		Config:    mt.Config,
		Version:   VERSION,
	}

	// Store our metric
	*metrics = append(*metrics, metric)
}
