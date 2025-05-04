Golang sample_scaffold_working_structure

### 🚀 Running Instructions
```
chmod +x run_demo.sh
./run_demo.sh
```

```
make         # sets up certs and runs server
make run-client  # in another terminal, runs the client
```

<pre>
chmod +x generate_certs.sh
./generate_certs.sh

Ensure you create a file to send:
echo "Hello, secure world!" > file_to_send

Then start the server:
go run server.go

Then start the client:
go run client.go

You should see output like:
✅ File integrity verified
</pre>

```
.
├── certs
│   ├── ca.crt
│   ├── ca.key
│   ├── ca.srl
│   ├── client.crt
│   ├── client.csr
│   ├── client.ext
│   ├── client.key
│   ├── server.crt
│   ├── server.csr
│   ├── server.ext
│   └── server.key
├── client.go
├── file_to_send
├── generate_certs.sh
├── go.mod
├── readme.md
├── received_file
├── server.go

```