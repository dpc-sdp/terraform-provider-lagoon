package lagoon

import (
	"context"
	"github.com/dpc-sdp/terraform-provider-lagoon/internal/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lagoon_client "github.com/uselagoon/machinery/api/lagoon/client"
	lagoon_schema "github.com/uselagoon/machinery/api/schema"
)

const (
	RESOURCE_KEY_PUBLIC  string = "public_key"
	RESOURCE_KEY_PRIVATE string = "private_key"
)

func dataSourceProjectKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectKeyRead,
		Schema: map[string]*schema.Schema{
			RESOURCE_KEY_PUBLIC: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			RESOURCE_KEY_PRIVATE: &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceProjectKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*lagoon_client.Client)

	name := d.Get("project").(string)
	project := &lagoon_schema.Project{}
	err := client.ProjectByName(ctx, name, project)
	if err != nil {
		return diag.FromErr(err)
	}

	publicKey, err := helpers.ConvertPrivateKeyToPublic(project.PrivateKey)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(RESOURCE_KEY_PUBLIC, publicKey)
	d.Set(RESOURCE_KEY_PRIVATE, project.PrivateKey)

	return diags
}
