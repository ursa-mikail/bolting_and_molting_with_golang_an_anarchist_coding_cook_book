
A SerDe (Serializer/Deserializer) refers to a user-defined way to convert data to and from bytes—a necessary step for transmitting data over systems like Kafka. While SerDes are commonly associated with Apache Kafka, custom SerDes can be used outside of Kafka in any system that needs to serialize and deserialize data, such as:

1. Message queues (e.g. RabbitMQ, ActiveMQ)
2. Custom socket-based communication protocols
3. Database blob storage
4. File-based data exchange (e.g., over S3 or local files)
5. RPC frameworks or microservices using HTTP/gRPC

Default SerDes in Kafka (for String, JSON, Avro, Protobuf, etc.) are prebuilt and tied to Kafka’s message processing system. However, if your data doesn’t conform to those formats or has unique requirements (e.g., custom encryption, schema evolution, compression), you need a custom SerDe.

1. A customized SerDe typically involves implementing:
```
serialize(obj) -> bytes
deserialize(bytes) -> obj
```

When to Use a Custom SerDe Outside Kafka
1. Custom network protocol: You're sending/receiving structured data over TCP/UDP.
2. Binary format like FlatBuffers/Cap’n Proto: You need efficiency and schema evolution.
3. Legacy systems or proprietary formats: Need to interoperate with older systems or specific standards.
4. Data transformation pipelines: Reading/writing compressed, encoded formats (e.g., GZIP + Base64 + JSON).
5. Security: You need encrypted serialization (e.g., sensitive data with embedded keys or tokens).

### ascii_string_to_byte_stream.go
```
Demonstrates a custom serializer (streaming a string to bytes) and deserializer (rebuilding the ASCII string from bytes)
```

### file_to_byte_stream.go
```
Written to a file (simulating serialization to a file stream),
Then read back (simulating deserialization),
Converted back to the original ASCII/Unicode string.
```

### file_to_byte_stream_with_gzip_and_base64.go
```
Compresses the input string using GZIP,
Encodes the compressed data using Base64,
Writes the Base64 string to a file,
Reads the file back,
Decodes the Base64,
Decompresses the GZIP,
And reconstructs the original string.
```
