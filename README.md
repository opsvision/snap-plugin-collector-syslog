# snap-plugin-collector-syslog
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
