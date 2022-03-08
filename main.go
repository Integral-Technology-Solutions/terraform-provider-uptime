package main

import (
	"bitbucket.org/integraltech/terraform-provider-uptime/uptime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: uptime.Provider,
	})
}
