package alertmanager

import (
	"context"

	"github.com/prometheus/alertmanager/api/v2/client/receiver"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAlertManagerReceivers() *plugin.Table {
	return &plugin.Table{
		Name:        "alertmanager_receivers",
		Description: "AlertManager Receiver.",
		List: &plugin.ListConfig{
			Hydrate: listReceiver,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listReceiver(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlertManager(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := receiver.NewGetReceiversParams()

	alerts, err := conn.Receiver.GetReceivers(opts)
	if err != nil {
		return nil, err
	}
	for _, r := range alerts.Payload {
		d.StreamListItem(ctx, r)
	}
	return nil, nil
}
