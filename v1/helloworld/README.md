# Hello World
This example demonstrates how to write go micro server and client.
> 1. Generate protobuf files.
> 2. Implement your service, and register with micro service.

## Protobuf
```
protoc --go_out=proto --micro_out=proto proto/helloworld.proto
```

## Run Server
```
go run server/main.go --server_name helloworld
```

## Run Client
```
go run client/main.go
```