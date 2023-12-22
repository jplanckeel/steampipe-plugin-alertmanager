select 
    am.created_by as "CreatedBy", 
    am.ends_at as "EndAt" , 
    am.comment as "Comment",
    am._ctx ->> 'connection_name' AS "Zone"   
from 
    alertmanager_silences as am
