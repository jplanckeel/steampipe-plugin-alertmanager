# Table: alertmanager_receivers

List all reicevers in the Alertmanager server.

## Examples

### Basic group info

```sql
select
  name
from
  alertmanager_receivers;
```


### Count receivers

```sql
select
  count(*)
from
  alertmanager_receivers;
```
