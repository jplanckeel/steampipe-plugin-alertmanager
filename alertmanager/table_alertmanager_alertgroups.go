package alertmanager

import (
	"context"

	"github.com/prometheus/alertmanager/api/v2/client/alertgroup"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAlertManagerAlertGroups() *plugin.Table {
	return &plugin.Table{
		Name:        "alertmanager_alertgroups",
		Description: "AlertManager Alert Groups.",
		List: &plugin.ListConfig{
			Hydrate: listAlertGroup,
		},
		Columns: []*plugin.Column{
			{Name: "alerts", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "receiver", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func listAlertGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlertManager(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := alertgroup.NewGetAlertGroupsParams()

	alerts, err := conn.Alertgroup.GetAlertGroups(opts)
	if err != nil {
		return nil, err
	}
	for _, ag := range alerts.Payload {
		d.StreamListItem(ctx, ag)
	}
	return nil, nil
}
