# http-rest-api

## default configs/apiserver.toml
```toml
bind_addr = ":8080"
log_level = "debug"
database_url = "host=localhost user=postgres password=userpass dbname=restapi_dev sslmode=disable"
session_key = "yoursecretsessionkey"
```

## default configs/apiserver_test.toml
```toml
database_url = "host=localhost user=postgres password=userpass dbname=restapi_test sslmode=disable"
```
