package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tdmalone/terraform-provider-dme/dme"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dme.Provider})
}
