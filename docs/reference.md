# Resource Reference

## uptime\_check\_ntp
Example:

```go
resource "uptime_check_ntp" "google" {
    name = "Google Public NTP"
    address = "time.google.com"
    contact_groups = ["Default", "NTP"]
    interval = 1
    locations = ["US-East", "GBR"]
    tags = ["terraform"]
}
```

Required attributes:

* **address**, *string*: address of the server under test

* **contact_groups**, *list(string)*: contact groups to alert

* **interval**, *number*: time interval between checks, in minutes

* **locations**, *list(string)*: probe server locations

Optional attributes:

* **name**, *string*: human-readable/friendly name

* **tags**, *list(string)*: tags to attach to the check

* **notes**, *string*: arbitrary notes for check

* **include_in_global_metrics**, *bool*: whether to include this check in global uptime metrics

* **ip_version**, *string, limited to "IPV4" or "IPV6"*: IP version to use

* **sensitivity**, *number*: number of probe servers that should detect a failure before an alert is triggered

* **threshold**, *number*: timeout threshold for server response, in seconds

* **port**, *number*: port where service is running