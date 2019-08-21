package main // 	"github.com/jitesoft/terraform-provider-cfssl"
import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jitesoft/terraform-provider-cfssl/cfssl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return cfssl.Provider()
		},
	})
}