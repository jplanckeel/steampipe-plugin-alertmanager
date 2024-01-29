package main

import (
	"github.com/turbot/steampipe-plugin-alertmanager/alertmanager"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: alertmanager.Plugin})
}
