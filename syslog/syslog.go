/*

Boilerplate licensing info goes here

*/

package syslog

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"time"
)

var (
	logCounter uint64
)

// Initialization
func init() {
	logCounter = 0
}

// SyslogCollector implementation
type SyslogCollector struct{}

/*
CollectMetrics collects metrics.

CollectMetrics() will be called by Snap when a task that collects one of the
metrics returned from this plugins GetMetricTypes() is started. The input will
include a slice of all the metric types being collected.

The output is the collected metrics as plugin.Metric and an error.
*/
func (SyslogCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	// increment our counter
	logCounter += 1

	for idx, mt := range mts {
		mts[idx].Timestamp = time.Now()

		if mt.Namespace[len(mt.Namespace)-1].Value == "counter" {
			mts[idx].Data = logCounter
			metrics = append(metrics, mts[idx])
		}
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

	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace("syslog", "counter"), Version: 1,
	})

	//metrics = append(metrics, plugin.Metric{
	//	Namespace: plugin.NewNamespace("syslog", "message"), Version: 1,
	//})

	return metrics, nil
}

/*
GetConfigPolicy() returns the configPolicy for your plugin.

A config policy is how users can provide configuration info to plugin. Here
you define what sorts of config info your plugin needs and/or requires.
*/
func (SyslogCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	policy.AddNewIntRule([]string{"port", "integer"},
		"port",
		false,
		plugin.SetMaxInt(65535),
		plugin.SetMinInt(1))

	return *policy, nil
}
