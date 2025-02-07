package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareNotificationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"alert_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"filters": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: schema.TypeString,
			},
		},
		"created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"conditions": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"email_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
		"webhooks_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
		"pagerduty_integration": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     mechanismData,
		},
	}
}

var mechanismData = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}
