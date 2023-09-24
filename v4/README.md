# Go-Micro V4
```
go get go-micro.dev/v4@latest
```

### Protobuf
- Download protoc
```
https://github.com/protocolbuffers/protobuf/releases
```
- Generator
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
```

### Micro CLI
```
go install github.com/asim/go-micro/cmd/micro@latest
```

### Dashboard
```
go install github.com/go-micro/dashboard@latest
```

### Protobuf
```
protoc --go_out=proto --micro_out=proto proto/helloworld.proto
protoc --go_out=proto --micro_out=proto proto/message.proto
protoc --go_out=proto --micro_out=proto proto/route_guide.proto
protoc --go_out=proto --micro_out=proto proto/statistics.proto
```