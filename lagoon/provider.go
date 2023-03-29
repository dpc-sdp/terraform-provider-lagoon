package lagoon

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	lagoon_client "github.com/uselagoon/machinery/api/lagoon/client"
)

const (
	CLIENT_USER_AGENT string = "github.com/dpc-sdp/terraform-provider-lagoon"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//DefaultFunc: // @todo see if we can use a function to load the yaml
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//DefaultFunc: // @todo see if we can use a function to load the yaml
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"lagoon_project_key": dataSourceProjectKey(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	endpoint := d.Get("endpoint").(string)

	client := lagoon_client.New(endpoint, CLIENT_USER_AGENT, &token, false)
	var diags diag.Diagnostics
	return client, diags
}
