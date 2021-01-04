package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rfalias/terraform-provider-pypwsh/pypwsh"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pypwsh.Provider,
	})
}
