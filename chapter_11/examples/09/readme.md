```

# Install tools (once)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

```
# Terminal 1
go run server/main.go

# Terminal 2
go run client/main.go
# Output: Hello, World!
```

```
Key pieces are:

.proto file defines your service and messages
protoc generates the Go boilerplate
Server implements the interface (SayHello)
Client dials and calls it like a normal function
```