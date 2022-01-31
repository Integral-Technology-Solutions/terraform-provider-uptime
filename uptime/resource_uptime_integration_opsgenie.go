package uptime

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	uptime "bitbucket.org/integraltech/uptime-rest-api-clients/golang/uptime"
)

func resourceUptimeIntegrationOpsgenie() *schema.Resource {
	return &schema.Resource{
		Create: resourceUptimeIntegrationOpsgenieCreate,
		Read: resourceUptimeIntegrationOpsgenieRead,
		Update: resourceUptimeIntegrationOpsgenieUpdate,
		Delete: resourceUptimeIntegrationOpsgenieDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"contact_groups": {
				Type: schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"api_endpoint": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"teams": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
			},
			"autoresolve": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func buildUptimeIntegrationOpsgenie(d *schema.ResourceData) (*uptime.Integration) {
	integrationOpsgenie := &uptime.Integration{
		Module: "Opsgenie",
		Name: d.Get("name").(string),
		ContactGroups: expandSetAttr(d.Get("contact_groups")),
		APIEndpoint: d.Get("api_endpoint").(string),
		APIKey: d.Get("api_key").(string),
		Teams: d.Get("teams").(string),
		Tags: d.Get("tags").(string),
		AutoResolve: d.Get("autoresolve").(bool),
	}
	return integrationOpsgenie
}

func resourceUptimeIntegrationOpsgenieCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*uptime.Client)
	ctx := context.Background()

	t := buildUptimeIntegrationOpsgenie(d)

	log.Printf("[INFO] Creating Uptime.com Opsgenie integration: %s", t.Name)
	t, _, err := client.Integrations.Create(ctx, t)
	if err != nil {
		return err
	}
	setResourceIDFromIntegrationOpsgenie(d, t)

	return resourceUptimeIntegrationOpsgenieRead(d, meta)
}

func resourceUptimeIntegrationOpsgenieRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*uptime.Client)
	ctx := context.Background()

	log.Printf("[INFO] Reading Uptime.com check Opsgenie integration: %s", d.Id())

	pk := pkFromResourceData(d)
	t, _, err := client.Integrations.Get(ctx, pk)
	if err != nil {
		if uptErr, ok := err.(*uptime.Error); ok {
			if uptErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing Opsgenie integration %d from state because it no longer exists in Uptime.com", pk)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", t.Name)
	d.Set("contact_groups", t.ContactGroups)
	d.Set("api_endpoint", t.APIEndpoint)
	d.Set("api_key", t.APIKey)
	d.Set("teams", t.Teams)
	d.Set("tags", t.Tags)
	d.Set("autoresolve", t.AutoResolve)
	d.Set("url", t.URL)
	return nil
}

func resourceUptimeIntegrationOpsgenieUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*uptime.Client)
	ctx := context.Background()

	t := buildUptimeIntegrationOpsgenie(d)
	pk := pkFromResourceData(d)
	t.PK = pk

	log.Printf("[DEBUG] Updating Opsgenie integration: %s", d.Id())

	newIntegration, _, err := client.Integrations.Update(ctx, t)
	if err != nil {
		return err
	}

	setResourceIDFromIntegrationOpsgenie(d, newIntegration)

	return resourceUptimeIntegrationOpsgenieRead(d, meta)
}

func resourceUptimeIntegrationOpsgenieDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*uptime.Client)
	ctx := context.Background()

	log.Printf("[INFO] Deleting Uptime Opsgenie integration: %s", d.Id())

	pk := pkFromResourceData(d)
	if _, err := client.Integrations.Delete(ctx, pk); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func setResourceIDFromIntegrationOpsgenie(d *schema.ResourceData, t *uptime.Integration) {
	id := strconv.Itoa(t.PK)
	d.SetId(id)
}
