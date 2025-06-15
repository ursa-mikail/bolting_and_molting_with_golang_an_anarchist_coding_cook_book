# Simple Makefile for gRPC project

# Variables
SERVER_PID_FILE = .server.pid

# Default target
all: clean proto build test

# Clean generated files
clean:
	@echo "🧹 Cleaning..."
	rm -rf api/ proto/*.pb.go proto/api/
	rm -rf bin/
	rm -f $(SERVER_PID_FILE)
	@echo "✅ Clean complete!"

# Generate protobuf files
proto:
	@echo "🔨 Generating protobuf files..."
	# Create directory
	mkdir -p api
	
	# Generate protobuf files
	protoc --proto_path=proto --go_out=. --go_opt=module=rpc-tutorial \
	       --go-grpc_out=. --go-grpc_opt=module=rpc-tutorial \
	       item.proto
	
	# Check files are generated correctly
	@echo "📁 Generated files:"
	ls -la api/
	@echo "✅ Protobuf generation complete!"

# Build applications
build:
	@echo "🏗️  Building..."
	mkdir -p bin
	go build -o bin/server ./server
	go build -o bin/client ./client
	@echo "✅ Build complete!"

# Kill any existing processes
kill-processes:
	@echo "💀 Killing existing processes..."
	-pkill -f "bin/server" 2>/dev/null || true
	-pkill -f "bin/client" 2>/dev/null || true
	-lsof -ti:4040 | xargs kill -9 2>/dev/null || true
	rm -f $(SERVER_PID_FILE)
	@echo "✅ Processes killed!"

# Test the application
test: kill-processes
	@echo "🧪 Testing..."
	# Start server in background and save PID
	./bin/server & echo $$! > $(SERVER_PID_FILE)
	
	# Wait for server to start
	sleep 2
	
	# Run client
	./bin/client
	
	# Clean up: kill server and remove PID file
	@echo "🧹 Cleaning up processes..."
	-kill $$(cat $(SERVER_PID_FILE)) 2>/dev/null || true
	-pkill -f "bin/server" 2>/dev/null || true
	-lsof -ti:4040 | xargs kill -9 2>/dev/null || true
	rm -f $(SERVER_PID_FILE)
	@echo "✅ Test complete and processes cleaned up!"

# Run server only
run-server: kill-processes
	@echo "🚀 Starting server..."
	./bin/server & echo $$! > $(SERVER_PID_FILE)
	@echo "✅ Server started (PID: $$(cat $(SERVER_PID_FILE)))"
	@echo "Use 'make stop-server' to stop it"

# Run client only
run-client:
	@echo "📡 Running client..."
	./bin/client

# Stop server
stop-server:
	@echo "🛑 Stopping server..."
	-kill $$(cat $(SERVER_PID_FILE)) 2>/dev/null || true
	-pkill -f "bin/server" 2>/dev/null || true
	-lsof -ti:4040 | xargs kill -9 2>/dev/null || true
	rm -f $(SERVER_PID_FILE)
	@echo "✅ Server stopped!"

# Quick build and test
quick: proto build test

# Check status
status:
	@echo "📊 Status:"
	@echo "Generated files:"
	@ls -la api/ 2>/dev/null || echo "  No API files found"
	@echo "Binaries:"
	@ls -la bin/ 2>/dev/null || echo "  No binaries found"
	@echo "Processes on port 4040:"
	@lsof -i:4040 2>/dev/null || echo "  Nothing running on port 4040"

# Help
help:
	@echo "📚 Available commands:"
	@echo "  make all         - Full clean, build, and test"
	@echo "  make clean       - Remove generated files"
	@echo "  make proto       - Generate protobuf files"
	@echo "  make build       - Build server and client"
	@echo "  make test        - Run test (with cleanup)"
	@echo "  make quick       - Proto + build + test"
	@echo "  make run-server  - Start server in background"
	@echo "  make run-client  - Run client"
	@echo "  make stop-server - Stop background server"
	@echo "  make status      - Show current status"

.PHONY: all clean proto build test kill-processes run-server run-client stop-server quick status help