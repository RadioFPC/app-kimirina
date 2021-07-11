# gRPC Example

This guide gets you started with gRPC in Go with a simple working example.

## Prerequisites

Install the protocol compiler plugins for Go using the following commands:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your `P