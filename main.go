package main

import (
	"bitbucket.org/integraltech/terraform-provider-uptime/uptime"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return uptime.Provider()
		},
	})
}
