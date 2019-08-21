package cfssl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jitesoft/terraform-provider-cfssl/cfssl/client"
	"github.com/jitesoft/terraform-provider-cfssl/cfssl/service"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CFSSL_URI", nil),
				Description: "URI to the cfssl remote server.",
			},
			"headers": {
				Type:        schema.TypeList,
				Required:    false,
				Description: "Extra headers to send with any API request. This can be used to add extra authentication headers in case your server is behind an API gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the header.",
							Required:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Value of the header.",
							Required:    false,
						},
					},
				},
			},
		},
		ResourcesMap:  {},
		ConfigureFunc: providerConfig,
	}
}

// Struct containing the configuration parameters.
type Config struct {
	Uri     string
	Headers map[string]string
}

// Initializes the provider.
func providerConfig(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Uri:     d.Get("uri").(string),
		Headers: d.Get("headers").(map[string]string),
	}

	c := client.New(config.Uri, config.Headers)
	s := service.New(c)

	err := c.Ping()
	return s, err
}
