# Create api directory
mkdir -p api

# Generate with module-aware paths
protoc --proto_path=proto --go_out=. --go_opt=module=rpc-tutorial \
       --go-grpc_out=. --go-grpc_opt=module=rpc-tutorial \
       item.proto

# Build
go build -o bin/server ./server
go build -o bin/client ./client

```
# 1. Clean
rm -rf api/ proto/*.pb.go proto/api/

# 2. Create directory
mkdir -p api

# 3. Generate protobuf files
protoc --proto_path=proto --go_out=. --go_opt=module=rpc-tutorial \
       --go-grpc_out=. --go-grpc_opt=module=rpc-tutorial \
       item.proto

# 4. Check files are generated correctly
ls -la api/
# Should show: item.pb.go and item_grpc.pb.go

# 5. Build
mkdir -p bin
go build -o bin/server ./server
go build -o bin/client ./client

# 6. Test
./bin/server &
sleep 2
./bin/client

```

"""
2025/06/15 13:48:56 gRPC server listening on :4040
Database: [title:"First"  body:"A first item" title:"Second"  body:"A second item" title:"Third"  body:"A third item"]
Database after edits: [title:"First"  body:"A first item" title:"Second"  body:"A new second item"]
First item: title:"First"  body:"A first item"
"""

```
.
├── api/
│   ├── item.pb.go
│   └── item_grpc.pb.go
├── bin/
│   ├── client
│   └── server
├── client/
│   └── main.go
├── go.mod
├── server/
│   └── main.go
└── proto/
    └── item.proto
```

Usage:
```
# Full workflow (your requested sequence)
make all

# Or individual steps
make clean
make proto
make build  
make test

# Quick rebuild and test
make quick

# Manual server/client control
make run-server  # starts server in background
make run-client  # runs client
make stop-server # stops server

# Check what's happening
make status
```