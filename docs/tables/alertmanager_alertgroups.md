# Table: alertmanager_alertgroups

List all reicevers in the Alertmanager server.

## Examples

### Basic group info

```sql
select
  alerts,
  labels,
  receiver
from
  alertmanager_alertgroups;
```


### Count receivers

```sql
select
  count(*)
from
  alertmanager_alertgroups;
```
