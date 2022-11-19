# Go-Micro V3

## Tools

### Protobuf

```
go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest
```

```
protoc --go_out=event/proto --micro_out=event/proto event/proto/statistics.proto
protoc --go_out=helloworld/proto --micro_out=helloworld/proto helloworld/proto/helloworld.proto
protoc --go_out=stream/proto --micro_out=stream/proto stream/proto/route_guide.proto
```
