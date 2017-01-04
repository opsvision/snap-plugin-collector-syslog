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

This Snap-Telemetry plugin collects events from Syslog in response to the [wishlist request](https://github.com/intelsdi-x/snap/issues/1117) for this feature.

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

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download
The simplest approach is to use ```go get``` to fetch and build the plugin. The following command will place the binary in your ```$GOPATH/bin``` folder where you can load it into snap.
```
$ go get github.com/dishmael/snap-plugin-collector-syslog
```

#### Building
You can also download the source and build it manually. The repository utilizes [glide](https://github.com/Masterminds/glide) for library management. Much like the previous method, executing ```go install``` will place the binary in your ```$GOPATH/bin``` folder.
```
$ git clone http://github.com/dishmael/snap-plugin-collector-syslog
$ glide up
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for ...
[INFO]	Resolving imports
[INFO]	--> Fetching updates for ...
[INFO]	Downloading dependencies. Please wait...
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting ...
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on ... dependencies.
$ go install
```

#### Source structure
The following file structure provides an overview of where the files exist in the source tree. The [syslog.go](https://github.com/dishmael/snap-plugin-collector-syslog/blob/master/syslog/syslog.go) file does all the work.
```
snap-plugin-collector-syslog
├── glide.yaml
├── LICENSE
├── main.go
├── README.md
└── syslog
    ├── syslog.go
    └── syslog_test.go
```

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
### Collected Metrics
This plugin has the ability to gather the following metrics:

| Namespace | Description (optional) |
| ----------|----------------------- |
| /opsvision/syslog/counter | a 64bit counter representing the number of log messages processed since the start of the collector |
| /opsvision/syslog/events/[source]/message | the entire syslog payload in JSON format |
| /opsvision/syslog/events/[source]/summary | the syslog message |

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
