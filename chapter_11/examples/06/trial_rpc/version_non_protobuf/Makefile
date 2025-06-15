.PHONY: all server client run clean

.PHONY: all run clean

all: run

run:
	@echo "Starting server..."
	@go run server/main.go &
	@sleep 1
	@echo "Running client (===== 1st time =====)..."
	@go run client/main.go
	@echo "Running client (=====2nd time =====)..."
	@go run client/main.go

clean:
	@echo "Killing Go processes..."
	@pkill -f "/var/folders/.*/exe/main" || true
	@pkill -f "go run" || true

#🔧 Usage:
#To run server then client:
#make
#
#or
#make run
#
#To kill lingering Go processes:
#make clean
#