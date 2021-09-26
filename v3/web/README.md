# Web Service

Use go-micro to handle http request with gin.

## Usage

```
1. create micro service with http server
2. create gin router
3. register server handler with router
```

## Run

web service

```
go run web/main.go
```

helloworld service

```
go run helloworld/server/main.go
```

## Testing

```
curl -X POST http://localhost:80/helloworld -d "{"""user""":"""test"""}"
```

```
{"message":"Hello test"}
```
