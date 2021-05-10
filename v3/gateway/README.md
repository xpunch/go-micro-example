# Gateway
This example demonstrates how to implement an gateway forward http request into grpc services.

# Run
```
go run main.go
```

## Access grpc services through http
### Endpoint: helloworld
### Method: Helloworld.Call
```
curl -X POST http://localhost:8080/helloworld/Helloworld.Call -H "Content-Type:application/json" -d "{"""name""":"""Joe"""}"
```
```
{"message":"Hello Joe"}
```

## View service nodes
```
curl http://localhost:8080/helloworld/nodes
```
```
[{"id":"helloworld-109f6d75-42cc-43ff-96df-f63260cf5372","address":"127.0.0.1:50000","metadata":{"broker":"http","protocol":"grpc","registry":"mdns","server":"grpc","transport":"grpc"}}]
```