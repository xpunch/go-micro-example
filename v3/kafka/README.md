# Kafka
Use kafka as broker to publish and subscribe event messages.

## Plugin
```
import _ "github.com/asim/go-micro/plugins/broker/kafka/v3"
```
Or my custom plugin:
```
import _ "github.com/x-punch/micro-kafka/v3"
```

## Run
### Client
```
go run kafka/client/main.go --broker kafka --broker_address 127.0.0.1:9092
```
### Server
```
go run kafka/server/main.go --broker kafka --broker_address 127.0.0.1:9092
```

## Testing
```
2021-09-26 17:04:03  file=server/rpc_server.go:840 level=info Broker [kafka] Connected to 127.0.0.1:9092
2021-09-26 17:04:03  file=server/rpc_server.go:706 level=info Subscribing to topic: kafka-topic
2021-09-26 17:04:04  file=server/main.go:30 level=info &{f2e1a241-a67d-484d-9026-2bd682736347 1632647044}
2021-09-26 17:04:05  file=server/main.go:30 level=info &{85684f0d-0e6d-432c-9369-c2a8385d9065 1632647045}
```
