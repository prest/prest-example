# Run Custom Middlewares

To run extending framework example with local database configure the following environment variables

In this example we desabled the JWT with the `PREST_JWT_DEFAULT=false` environment variable.

```console
export PREST_PG_USER=user_name
export PREST_PG_PASS=user_password
export PREST_PG_DATABASE=database_name
export PREST_JWT_DEFAULT=false
```

Start example with following command:

```console
go run main.go
```

Then you can access the url http://localhost:3000/ping and in this case the server will return `Calling custom middleware` and then `Pong!` as response.

```console
curl http://localhost:3000/ping
```
