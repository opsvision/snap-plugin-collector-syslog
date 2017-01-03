// +build small

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
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSyslogCollector(t *testing.T) {
	syslog := SyslogCollector{}
	
	// TODO: Write better testing
	
	Convey("Test SyslogCollector", t, func() {
		Convey("Collect Integer", func() {
			metrics := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("opsvision", "syslog", "counter"),
					Config:    map[string]interface{}{"port": int64(1514)},
					Data:      34,
					Tags:      map[string]string{"hello": "world"},
					Unit:      "int",
					Timestamp: time.Now(),
				},
			}
			mts, err := syslog.CollectMetrics(metrics)
			So(mts, ShoulBeEmpty)
		})
	})
}
