package alertmanager

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-opsgenie",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"alertmanager_alerts":      tableAlertManagerAlerts(),
			"alertmanager_alertgroups": tableAlertManagerAlertGroups(),
			"alertmanager_receivers":   tableAlertManagerReceivers(),
			"alertmanager_silences":    tableAlertManagerSilences(),
		},
	}
	return p
}
