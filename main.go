package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"bitbucket.org/integraltech/terraform-provider-uptime/uptime"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func () terraform.ResourceProvider {
			return uptime.Provider()
		},
	})
}
