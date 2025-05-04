# Makefile

CERT_DIR=certs
SEND_FILE=file_to_send
RECEIVED_FILE=received_file

all: setup build run-server

setup:
	@echo "🚧 Setting up certs and files..."
	chmod +x generate_certs.sh
	./generate_certs.sh
	echo "Hello, secure world!" > $(SEND_FILE)

build:
	@echo "🔧 Checking Go code for build..."
	#go vet ./...
	#go vet server.go client.go common.go
	#go vet server.go common.go
	#go vet client.go common.go
	go vet server.go client.go 
	@echo "✅ Build checks passed."

run-server:
	@echo "🚀 Starting the server..."
	go run server.go 

run-client:
	@echo "⏳ Waiting for server to start..."
	sleep 1	
	@echo "📤 Running the client..."
	go run client.go 

clean:
	@echo "🧹 Cleaning up files..."
	rm -f $(SEND_FILE) $(RECEIVED_FILE)
	rm -rf $(CERT_DIR)

reset: clean setup

help:
	@echo "🛠️  Available commands:"
	@echo "  make               → Setup, build, and run the server"
	@echo "  make run-client    → Run the client"
	@echo "  make clean         → Remove all generated files"
	@echo "  make reset         → Clean and regenerate everything"

