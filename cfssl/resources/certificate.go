package resources

import "github.com/hashicorp/terraform/helper/schema"

func CertificateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema {
		},
	}
}
