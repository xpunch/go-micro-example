# MQTT

### Protobuf

```
protoc --go_out=proto proto/message.proto
```

### Run
```
cd client
go run . --broker mqtt --broker_address 127.0.0.1:1883
```

```
cd server
go run . --broker mqtt --broker_address 127.0.0.1:1883
```

### Result
```
2021-11-18 13:29:06  file=v4@v4.4.0/service.go:206 level=info Starting [service] mqtt-server
2021-11-18 13:29:06  file=server/rpc_server.go:820 level=info Transport [http] Listening on [::]:64282
2021-11-18 13:29:06  file=server/rpc_server.go:840 level=info Broker [mqtt] Connected to tcp://127.0.0.1:1883
2021-11-18 13:29:06  file=server/rpc_server.go:654 level=info Registry [mdns] Registering node: mqtt-server-286497dc-f8d4-4f0b-afc6-e3181b053fc5
2021-11-18 13:29:06  file=server/rpc_server.go:706 level=info Subscribing to topic: micro.broker.mqtt.json
2021-11-18 13:29:06  file=server/rpc_server.go:706 level=info Subscribing to topic: micro.broker.mqtt.protobuf
2021-11-18 13:29:23  file=server/main.go:26 level=info &{92314b4f-d3dd-483c-a793-d711db01a791 No. 1 1637213363}
2021-11-18 13:29:23  file=server/main.go:21 level=info id:"cae41692-79d8-4a54-b049-6d51a50ba12e" message:"No. 1" timestamp:1637213363
```