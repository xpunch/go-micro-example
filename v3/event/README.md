# Event

This example demonstrates how to handle event between client and server.

> 1. Client will publish access event when handle http request.
> 2. Server will subscribe access event and update access statistics.

## Protobuf

```
protoc --go_out=proto --micro_out=proto proto/statistics.proto
```

## Run Server

```
go run event/server/main.go event/server/handler.go event/server/subscriber.go
```
### Tips
> Panic recover is required in subscribe method to prevent crash when panic occur in subscribe goroutine.

## Run Client

```
go run event/client/main.go
```

## Testing
```
curl -X POST http://localhost:80/helloworld?user=test
{"message":"Hello test"}
```
```
curl -X GET http://localhost:80/statistics?method=POST
{"access_count":1}
```
```
curl -X GET http://localhost:80/statistics?method=GET
{"access_count":1}
```
```
curl -X GET http://localhost:80/statistics
{"access_count":3}
```