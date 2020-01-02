package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/spotmaxtech/terraform-provider-alicloud/alicloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: alicloud.Provider})
}
