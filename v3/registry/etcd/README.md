# ETCD
This example demonstrates how to use plugins, and it shows us how to use etcd as registry instead of default mdns.

## Run
```
go run main.go plugins.go --registry etcd --server grpc --server_name helloworld
```
```
Registry [etcd] Registering node: helloworld-cfeccaf2-28cc-4d4e-8eea-55c60a9d33ce
```

## Metadata
```
curl http://localhost:8080/helloworld/nodes
```
```json
[{
	"id": "helloworld-7779ad77-8836-46e8-9b4d-e7993e764fd5",
	"address": "127.0.0.1:54834",
	"metadata": {
		"broker": "http",
		"protocol": "grpc",
		"registry": "etcd",
		"server": "grpc",
		"transport": "grpc"
	}
}]
```