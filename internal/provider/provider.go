package provider

/*
	Usage:
	```
	provider "uptimerobot" {
	  api_key = "[YOUR MAIN API KEY]"
	  retries = 10
	}
	```
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	uptimerobotapi "github.com/vexxhost/terraform-provider-uptimerobot/internal/provider/api"
)

// Provider returns a schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTIMEROBOT_API_KEY", nil),
			},
			"retries": {
				Type:         schema.TypeInt,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("UPTIMEROBOT_RETRIES", 10),
				Description:  "Maximum number of retries to perform when an API request fails. This can also be provided as an environment variable `UPTIMEROBOT_RETRIES`",
				ValidateFunc: validation.IntBetween(0, 10000),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"uptimerobot_account":       dataSourceAccount(),
			"uptimerobot_alert_contact": dataSourceAlertContact(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_alert_contact": resourceAlertContact(),
			"uptimerobot_monitor":       resourceMonitor(),
			"uptimerobot_status_page":   resourceStatusPage(),
		},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := uptimerobotapi.New(r.Get("api_key").(string), r.Get("retries").(int))
			return config, nil
		},
	}
}
