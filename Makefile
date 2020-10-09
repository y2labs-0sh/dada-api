.PHONY: generate

generate:
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build: generate security
	@go build -o ./build/server
	@echo "[OK] App binary was created!"

build-testnet: generate security
	@go build -tag testnet -o ./build/server-testnet
	@echo "[OK] App binary(testnet) was created!"

run:
	@./build/server
