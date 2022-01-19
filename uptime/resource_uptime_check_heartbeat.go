package uptime

import (
	uptime "github.com/uptime-com/rest-api-clients/golang/uptime"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUptimeCheckHeartbeat() *schema.Resource {
	return &schema.Resource{
		Create: checkCreateFunc(httpCheck),
		Read: checkReadFunc(httpCheck),
		Update: checkUpdateFunc(httpCheck),
		Delete: checkDeleteFunc(httpCheck),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// Required attributes: Common
			"contact_groups": {
				Type: schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Required attributes: Specific
			"interval": {
				Type: schema.TypeInt,
				Required: true,
			},

			// Optional attributes: Common
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_in_global_metrics": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"is_paused": {
				Type: schema.TypeBool,
				Optional: true,
				Computed: false,
			},
			"notes": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Managed by Terraform",
			},
			"uptime_sla": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"response_time_sla": {
				Type: schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

// HeartbeatCheck implements the CheckType interface for Uptime.com Heartbeat checks.
type HeartbeatCheck struct{}

func (HeartbeatCheck) typeStr() string {return "Heartbeat"}

func (HeartbeatCheck) getSpecificAttrs(d *schema.ResourceData, c *uptime.Check) {
	if attr, ok := d.GetOk("interval"); ok {
		c.Interval = attr.(int)
	}
}

func (HeartbeatCheck) setSpecificAttrs(d *schema.ResourceData, c *uptime.Check) {
	d.Set("interval", c.Interval)
}

var heartbeatCheck HeartbeatCheck
