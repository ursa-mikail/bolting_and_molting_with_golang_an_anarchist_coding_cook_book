# RPC

## Version Protobuf
Protocol Buffers is a language-neutral, platform-neutral, extensible way of serializing structured data — think of it like a faster, smaller, and simpler alternative to JSON or XML.

It defines a schema (.proto files) describing your data structures (messages) and services (RPC calls).


1. Defining Data and RPC Service
Write a .proto file that describes:
Your messages (e.g., Item with name and description fields).
Your service API (e.g., AddItem, GetDB, EditItem, etc.) with their input/output messages.

```
syntax = "proto3";

message Item {
  string name = 1;
  string description = 2;
}

service API {
  rpc AddItem(Item) returns (Item);
  rpc GetDB(Empty) returns (ItemList);
  rpc EditItem(Item) returns (Item);
  // etc.
}
```

###  Code Generation
Run the protobuf compiler (protoc) with language plugins.
- It generates code in your target language (Go, Python, Java, etc.) for:
- The message types (e.g., Item struct/class).
- The client and server stubs to serialize/deserialize data and handle RPC calls.

### Transport
The client and server communicate by sending protobuf-encoded binary data over the wire.

Protobuf makes the communication efficient and ensures that both sides use the same data schema.

### In Go RPC Server-Client
If rpc.DialHTTP and API calls were implemented with protobuf support:
Your Item structs would be generated from .proto files.
RPC calls like "API.AddItem" would serialize the Item into protobuf binary before sending.
The server would deserialize it, process it, then serialize the response back.

| Step                      | Without protobuf               | With protobuf                             |
| ------------------------- | ------------------------------ | ----------------------------------------- |
| Define messages & service | Manually written structs & RPC | Write `.proto` schema & generate code     |
| Serialize data            | JSON or manual encoding        | Protobuf binary serialization             |
| Communication             | Text or manual binary format   | Protobuf binary over HTTP/gRPC or others  |
| Safety & compatibility    | Handled manually               | Built-in schema evolution & strict typing |

## For python rpc access
```
pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. item.proto
```

``` python client.py
import grpc
import item_pb2
import item_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:4040')
    stub = item_pb2_grpc.APIStub(channel)

    print("Adding item...")
    stub.AddItem(item_pb2.Item(name="Hello", description="From Python"))

    print("Getting DB...")
    db = stub.GetDB(item_pb2.Empty())
    for item in db.items:
        print(f"{item.name} - {item.description}")

if __name__ == '__main__':
    run()
```

| What          | How                                      |
| ------------- | ---------------------------------------- |
| Schema        | `.proto` file defines it once            |
| Go server     | Use `protoc-gen-go` to generate server   |
| Python client | Use `grpcio-tools` to generate client    |
| Transport     | gRPC with protobuf                       |
| Result        | Cross-language RPC with full type safety |

✅ Steps:
Run server/main.go (Go gRPC server).
Call it using client.py (Python gRPC client).
Communicate using shared item.proto.

```
.
.
├── proto/
│   └── item.proto       ← Shared schema
├── server/
│   └── main.go          ← gRPC server in Go
├── client/
│   └── client.py        ← gRPC client in Python
```

```
syntax = "proto3";

package item;

message Item {
  string name = 1;
  string description = 2;
}

message Empty {}

message ItemList {
  repeated Item items = 1;
}

service API {
  rpc AddItem(Item) returns (Item);
  rpc GetDB(Empty) returns (ItemList);
}
```

 Step: Install tools
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
```
pip install grpcio grpcio-tools
```

✅ Step : Generate protobuf code
From root of your project:
```
### For Go (generates into ./server)
protoc --go_out=server --go-grpc_out=server -Iproto proto/item.proto
```
### For Python (generates into ./client)
```
python -m grpc_tools.protoc -Iproto --python_out=client --grpc_python_out=client proto/item.proto
```

You should now have:
```
server/
├── item.pb.go
├── item_grpc.pb.go

client/
├── item_pb2.py
├── item_pb2_grpc.py
```


