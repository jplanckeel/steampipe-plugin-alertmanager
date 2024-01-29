package alertmanager

import (
	"context"

	"github.com/prometheus/alertmanager/api/v2/client/silence"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAlertManagerSilences() *plugin.Table {
	return &plugin.Table{
		Name: "alertmanager_silences",
		//TODO: change description
		Description: "AlertManager Silences.",
		List: &plugin.ListConfig{
			Hydrate: listSilence,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "created_by", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "starts_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ends_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "matchers", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listSilence(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlertManager(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := silence.NewGetSilencesParams()

	silences, err := conn.Silence.GetSilences(opts)
	if err != nil {
		return nil, err
	}
	for _, s := range silences.Payload {
		d.StreamListItem(ctx, s)
	}
	return nil, nil
}
