
# AlertManager Plugin for Steampipe

Use SQL to query infrastructure including alerts, receivers, silences and more from AlertManager.

- **[Get started →](https://hub.steampipe.io/plugins/jplanckeel/alertmanager)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/jplanckeel/alertmanager/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/jplanckeel/alertmanager/issues)

## Quick start

### Install

Download and install the latest AlertManager plugin:

```bash
steampipe plugin install jplanckeel/alertmanager
```

Configure your [credentials](https://hub.steampipe.io/plugins/jplanckeel/alertmanager#credentials) and [config file](https://hub.steampipe.io/plugins/jplanckeel/alertmanager#configuration).

Configure your account details in `~/.steampipe/config/alertmanager.spc`:

```hcl
connection "steampipe-plugin-alertmanager" {
  plugin  = "jplanckeel/alertmanager"

  # address to access the API (required).
  #address = "alertmanager.localhost.com:8080"
}
```

Run steampipe:

```shell
steampipe query
```

List teams in your AlertManager account:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/jplanckeel/steampipe-plugin-alertmanager.git
cd steampipe-plugin-alertmanager
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/alertmanager.spc
```

Try it!

```
steampipe query
> .inspect alertmanager
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/jplanckeel/steampipe-plugin-alertmanager/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [AlertManager Plugin](https://github.com/jplanckeel/steampipe-plugin-alertmanager/labels/help%20wanted)
