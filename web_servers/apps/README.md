## In-memory database

This is an example of how something more complex, like a Redis, memcached or beanstalkd database is based upon.

Run the in-memory database, listening on :8080

```
go run main.go
```

You can start using the in-memory database like this:

```
IN-MEMORY DATABASE

USE:
        SET key value 
        GET key 
        DEL key 

SET dog maca
GET dog
maca
DEL dog
GET dog
```