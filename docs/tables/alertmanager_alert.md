# Table: alertmanager_alertgroups

List all reicevers in the Alertmanager server.

## Examples

### Basic group info

```sql
select
  alert,
  annotations,
  status,
  starts_at
from
  alertmanager_alerts;
```
### Alerts with start date more 30 days

```sql
select
   count(*) AS NumberOfSilences
from
   alertmanager_alerts
where
   starts_at :: timestamptz >= now() - '30 days' :: interval
```

### Count receivers

```sql
select
  count(*)
from
  alertmanager_alerts;
```
