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

This Snap-Telemetry plugin collects events from Syslog.

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.7+](https://golang.org/dl/) (needed only for building)

Note: This plugin does not require Python rather it depends on the go library [gopsutil](https://github.com/shirou/gopsutil).  

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download
TBD

#### Building
TBD

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
### Collected Metrics
This plugin has the ability to gather the following metrics:

| Namespace | Description (optional) |
| ----------|----------------------- |
| /opsvision/syslog/counter | a 64bit counter representing the number of log messages processed since the start of the collector |
| /opsvision/syslog/events/[source]/message | the syslog event in JSON format |
| /opsvision/syslog/testing | a metric used for testing - will be omitted later |

Note: The [source] will be either the source hostname or IP address.

### Examples

```
/opsvision/syslog/counter: 1
/opsvision/syslog/events/foo/message: {}
/opsvision/syslog/testing: test message
```

### Roadmap
TBD

## Acknowledgements
* Author: [@dishmael](https://github.com/dishmael/)
