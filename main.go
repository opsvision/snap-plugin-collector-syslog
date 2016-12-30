/*

Boilerplate licensing info goes here

*/

package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"opsvision.com/dishmael/snap-plugin-collector-syslog/syslog"
)

const (
	pluginName    = "syslog"
	pluginVersion = 1
)

func main() {
	plugin.StartCollector(syslog.SyslogCollector{}, pluginName, pluginVersion)
}
