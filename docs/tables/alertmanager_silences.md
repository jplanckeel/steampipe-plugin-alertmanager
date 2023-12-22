# Table: alertmanager_silences

List all silences in the Alertmanager server.

## Examples

### Basic group info

```sql
select
  id,
  status,
  ends_at,
  created_by,
  comment
from
  alertmanager_silences;
```

### Silence with ends more 30 days

```sql
select
   id,
   created_by,
   comment,
   ends_at
from
   alertmanager_silences
where
   ends_at :: timestamptz >= now() + '30 days' :: interval
```

### Silence with ends more 30 days

```sql
select
   id,
   created_by,
   comment,
   ends_at
from
   alertmanager_silences
where
   ends_at :: timestamptz >= now() + '365 days' :: interval
```

### Count silences

```sql
select
   count(*) AS NumberOfSilences
from
   alertmanager_silences
```

### Silence with ends more 30 days

```sql
select
   count(*) AS NumberOfSilences
from
   alertmanager_silences
where
   ends_at :: timestamptz >= now() + '365 days' :: interval
```
