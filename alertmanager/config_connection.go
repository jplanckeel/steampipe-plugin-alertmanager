package alertmanager

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type alertManagerConfig struct {
	Address *string `cty:"address"`
	Schemes *string `cty:"schemes"`
	Path    *string `cty:"path"`
	CertCA  *string `cty:"cert_ca"`
	CertKey *string `cty:"cert_key"`
	Cert    *string `cty:"cert"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"schemes": {
		Type: schema.TypeString,
	},
	"path": {
		Type: schema.TypeString,
	},
	"cert_ca": {
		Type: schema.TypeString,
	},
	"cert_key": {
		Type: schema.TypeString,
	},
	"cert": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &alertManagerConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) alertManagerConfig {
	if connection == nil || connection.Config == nil {
		return alertManagerConfig{}
	}
	config, _ := connection.Config.(alertManagerConfig)

	return config
}
