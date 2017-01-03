<!--
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
-->
# **Snap-Telemetry Collector for Syslog** [![Build Status](https://travis-ci.org/dishmael/snap-plugin-collector-syslog.svg?branch=master)](https://travis-ci.org/dishmael/snap-plugin-collector-syslog) [![Go Report Card](https://goreportcard.com/badge/github.com/dishmael/snap-plugin-collector-syslog)](https://goreportcard.com/report/github.com/dishmael/snap-plugin-collector-syslog)
Snap-Telemetry collector plugin for Syslog messages

##Collected Metrics
The following metrics are collected.  The <source> will be either the source hostname or IP address.
```
/opsvision/syslog/counter
/opsvision/syslog/events/<source>/app_name
/opsvision/syslog/events/<source>/client
/opsvision/syslog/events/<source>/facility
/opsvision/syslog/events/<source>/message
/opsvision/syslog/events/<source>/msg_id
/opsvision/syslog/events/<source>/priority
/opsvision/syslog/events/<source>/proc_id
/opsvision/syslog/events/<source>/severity
/opsvision/syslog/events/<source>/structured_data
/opsvision/syslog/events/<source>/timestamp
/opsvision/syslog/events/<source>/tls_peer
/opsvision/syslog/events/<source>/version
/opsvision/syslog/testing
```
