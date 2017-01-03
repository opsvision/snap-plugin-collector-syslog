/*

Boilerplate licensing info goes here

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
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 34)
		})
	})
}
