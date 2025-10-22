## Using grpcurl (Alternative to Client)
You can also test with grpcurl:

```bash
# Install grpcurl if not already installed
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List services
grpcurl -plaintext localhost:50051 list

# Call the method
grpcurl -plaintext -d '{"data": "Hello from grpcurl"}' localhost:50051 example.SecureService/ProcessData
```

# Generate protobuf:
```
protoc --go_out=. --go-grpc_out=. proto/service.proto
```

## Running the Example
Start the server:

```bash
cd server
go run main.go
```

Run the client:

```bash
cd client
go run main.go
```

