---
organization: jplanckeel
category: ["alert management"]
icon_url: ""
brand_color: "#e0653b"
display_name: "AlertManager"
short_name: "alertmanager"
description: "Steampipe plugin to query alerts, silences and more from AlertManager."
og_description: "Query AlertManager with SQL! Open source CLI. No DB required."
og_image: 
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# AlertManager + Steampipe

[AlertManager](https://prometheus.io/docs/alerting/latest/alertmanager/) handles alerts sent by client applications such as the Prometheus serverservers and other infrastructure services.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List dashboards in your AlertManager Alerts:

```sql
select
  id,
  status,
  ends_at,
  created_by,a
  comment
from
  alertmanager_silences
```

```
+--------------------------------------+--------------------+--------------------------+------------+------------------------------------------------+
| id                                   | status             | ends_at                  | created_by | comment                                        |
+--------------------------------------+--------------------+--------------------------+------------+------------------------------------------------+
| 7bd63388-c0b6-475f-9dc3-0180d84dd8a2 | {"state":"active"} | 2033-05-29T17:08:23.298Z | jplanckeel |  silence alert ID34                            |
+--------------------------------------+--------------------+--------------------------+------------+------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/alertmanager/tables)**

## Get started

### Install

Download and install the latest AlertManager plugin:

```bash
steampipe plugin install jplanckeel/alertmanager
```

### Configuration

Installing the latest alertmanager plugin will create a config file (`~/.steampipe/config/alertmanager.spc`) with a single connection named `alertmanager`:

```hcl
connection "steampipe-plugin-alertmanager" {
  plugin  = "jplanckeel/alertmanager"

  # address to access the API (required).
  #address = "alertmanager.localhost.com:8080"
}
```

- `path` - To change path to access the api of the AlertManager.
- `ca_cert` - Certificate CA bundle to use to verify the AlertManager server's certificate. May alternatively be set via the `ALERTMANAGER_CA_CERT` environment variable.
- `tls_cert` - Client TLS certificate file to use to authenticate to the AlertManager server. May alternatively be set via the `ALERTMANAGER_TLS_CERT` environment variable.
- `tls_key` - Client TLS key file to use to authenticate to the AlertManager server. May alternatively be set via the `ALERTMANAGER_TLS_KEY` environment variable.
