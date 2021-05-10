# Hello World

## Protobuf
```
protoc --go_out=proto --micro_out=proto proto/helloworld.proto
```

## Run Server
```
cd server
go run main.go --server_name helloworld
```

## Run Client
```
cd client
go run main.go
```