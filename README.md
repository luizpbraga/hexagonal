# hexagonal

I hated this. Fuck me
 
> "Keep it simple. Definitely don't add lots of abstraction before you have a good reason to need it."
>  — Randon guy in reddit

### BUILD

using Docker
```
    some docker command
```
or locally, after config the database (see init.sql)
```
    go run .
```

### Request

using grpcurl
```
    grpcurl -plaintext -d '{"a":5,"b":3}' -import-path hexagonal/adapters/grpc/proto/ -proto arithmetics_svc.proto localhost:9000 pb.ArithmeticService/GetAddition
```
or using an client (see client/client.go)
