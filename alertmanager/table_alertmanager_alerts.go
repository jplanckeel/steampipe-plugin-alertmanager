package alertmanager

import (
	"context"

	"github.com/prometheus/alertmanager/api/v2/client/alert"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAlertManagerAlerts() *plugin.Table {
	return &plugin.Table{
		Name:        "alertmanager_alerts",
		Description: "AlertManager Alerts.",
		List: &plugin.ListConfig{
			Hydrate: listAlert,
		},
		Columns: []*plugin.Column{
			{Name: "alert", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "status", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "starts_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ends_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "receivers", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

func listAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlertManager(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := alert.NewGetAlertsParams()

	alerts, err := conn.Alert.GetAlerts(opts)
	if err != nil {
		return nil, err
	}
	for _, a := range alerts.Payload {
		d.StreamListItem(ctx, a)
	}
	return nil, nil
}
