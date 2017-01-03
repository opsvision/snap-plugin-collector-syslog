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
	"fmt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"gopkg.in/mcuadros/go-syslog.v2"
	"time"
)

// Constants
const (
	NS_VENDOR = "opsvision"
	NS_PLUGIN = "syslog"
	VERSION   = 1
)

// Types
type SyslogCollector struct{}

// Global variables
var channel = make(syslog.LogPartsChannel, 1024)
var logCounter uint64

// Initialization
func init() {
	logCounter = 1

	go func() {
		handler := syslog.NewChannelHandler(channel)

		server := syslog.NewServer()
		server.SetFormat(syslog.RFC5424)
		server.SetHandler(handler)
		server.ListenUDP("0.0.0.0:1514")
		server.Boot()
		server.Wait()
	}()
}

/*
	CollectMetrics collects metrics.

	CollectMetrics() will be called by Snap when a task that collects one of the
	metrics returned from this plugins GetMetricTypes() is started. The input will
	include a slice of all the metric types being collected.

	The output is the collected metrics as plugin.Metric and an error.
*/
func (SyslogCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	time := time.Now()

	/*
		Non-Blocking read from the syslog channel results in NULL entries in the
		log file every time this function is called.
	*/
	select {
	case logParts := <-channel:
		if len(logParts["hostname"].(string)) == 0 {
			break
		}

		// Iterate over the requested metrics
		for _, mt := range mts {
			metricName := mt.Namespace[len(mt.Namespace)-1].Value

			switch metricName {
			case "counter":
				setCounterMetric(mt, &metrics, time)

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
func (SyslogCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	// namespace: /NS_VENDOR/NS_PLUGIN/counter
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "counter"),
		Version:   VERSION,
	})

	// namespace: /NS_VENDOR/NS_PLUGIN/testing
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "testing"),
		Version:   VERSION,
	})

	// namespace: /NS_VENDOR/NS_PLUGIN/*/message
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NS_VENDOR, NS_PLUGIN, "events").
			AddDynamicElement("source", "the source hostname or IP address").
			AddStaticElement("message"),
		Version: VERSION,
	})

	return metrics, nil
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.

	A config policy is how users can provide configuration info to plugin. Here
	you define what sorts of config info your plugin needs and/or requires.
*/
func (SyslogCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	// The UDP and TCP port the syslog server should listen
	policy.AddNewIntRule([]string{"port", "integer"},
		"port",
		false,
		plugin.SetMaxInt(65535),
		plugin.SetMinInt(1))

	return *policy, nil
}

/*
	setCounterMetric is used to create the counter metric
*/
func setCounterMetric(mt plugin.Metric, metrics *[]plugin.Metric, time time.Time) {
	metric := plugin.Metric{
		Timestamp: time,
		Namespace: mt.Namespace,
		Data:      logCounter,
		Tags:      mt.Tags,
		Config:    mt.Config,
		Version:   VERSION,
	}

	// Store our metric
	*metrics = append(*metrics, metric)

	// Increment the counter
	logCounter += 1
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
	fields := []string{
		"app_name",
		"client",
		"facility",
		"message",
		"msg_id",
		"priority",
		"proc_id",
		"severity",
		"structured_data",
		"timestamp",
		"tls_peer",
		"version",
	}

	for i := 0; i < len(fields); i++ {
		metricName := fields[i]
		hostname := fmt.Sprintf("%v", logParts["hostname"])
		value := fmt.Sprintf("%v", logParts[metricName])

		ns := plugin.NewNamespace(NS_VENDOR, NS_PLUGIN,
			"events", hostname, metricName)

		metric := plugin.Metric{
			Timestamp: time,
			Namespace: ns,
			Data:      value,
			Tags:      mt.Tags,
			Config:    mt.Config,
			Version:   VERSION,
		}

		// Store our metric
		*metrics = append(*metrics, metric)
	}
}
