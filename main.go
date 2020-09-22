package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/fastly/terraform-provider-fastly/fastly"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fastly.Provider})
}
