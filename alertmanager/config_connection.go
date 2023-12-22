package alertmanager

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type alertManagerConfig struct {
	Address *string `cty:"address"`
	Path    *string `cty:"path"`
	CaCert  *string `cty:"ca_cert"`
	TlsKey  *string `cty:"tls_key"`
	TlsCert *string `cty:"tls_cert"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"path": {
		Type: schema.TypeString,
	},
	"ca_cert": {
		Type: schema.TypeString,
	},
	"tls_key": {
		Type: schema.TypeString,
	},
	"tls_cert": {
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
