# Description

This example is translate [go grpc example](https://grpc.io/docs/languages/go/basics/) into go-micro.
You can also find the orignal codes in [github.com/grpc/grpc-go](https://github.com/grpc/grpc-go/tree/master/examples/route_guide).

# Notes

- Not support client.SendClose(), details in [issue #2212](https://github.com/asim/go-micro/issues/2212).

  ```go
  // used to present stream.CloseAndRecv() in grpc
  if err := stream.Send(&pb.Point{Latitude: INT32_MAX, Longitude: INT32_MAX}); err != nil {
     logger.Fatal(err)
  }
  ```
  ```
  if err == io.EOF || proto.Equal(point, invalidPoint) {
	break
  }
  ```

- Do not forgot to close grpc stream connection.

  ```go
  stream, err := client.RouteChat(ctx)
  if err != nil {
        logger.Fatal(err)
  }
  // IMPORTANT: do not forgot to close stream
  defer stream.Close()
  ```

# Run the sample code

## Protobuf

```
protoc --go_out=proto --micro_out=proto proto/route_guide.proto
```

## Server

```
cd stream/server
go run .
```

## Client

```
cd stream/client
go run main.go
```
