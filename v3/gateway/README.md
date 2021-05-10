# Gateway
This example demonstrates how to implement an gateway forward http request into grpc services.

# Run
```
go run main.go
```

# Access grpc services through http
## Endpoint: helloworld
## Method: Helloworld.Call
```
curl -v -X POST http://localhost:8080/helloworld/Helloworld.Call -H "Content-Type:application/json" -d "{"""name""":"""Joe"""}"
```