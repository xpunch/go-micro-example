# Go-Micro V4
```
go get go-micro.dev/v4
```

### Protobuf
```
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
```

### Micro CLI
```
go install github.com/asim/go-micro/cmd/micro@latest
```

### Dashboard
```
go install github.com/xpunch/go-micro-dashboard@latest
```

### Protobuf
```
protoc --go_out=proto --micro_out=proto proto/helloworld.proto
protoc --go_out=proto --micro_out=proto proto/message.proto
protoc --go_out=proto --micro_out=proto proto/route_guide.proto
protoc --go_out=proto --micro_out=proto proto/statistics.proto
```