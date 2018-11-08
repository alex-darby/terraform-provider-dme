package main

import (
	"github.com/alex-darby/terraform-provider-dme/dme"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dme.Provider})
}
