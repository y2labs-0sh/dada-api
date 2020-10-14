.PHONY: generate security run

Branch=$(shell git rev-parse --abbrev-ref HEAD)
Commit=$(shell git rev-parse HEAD)
BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

security:
	@gosec -exclude -nosec=true ./...
	@echo "[OK] Go security check was completed!"

generate:
	@go generate ./...
	@echo "[OK] Files added to embed box!"

build: generate security
	@go build \
		-ldflags "-X main.Branch=$(Branch) -X main.Commit=$(Commit) -X main.BuildTime=$(BuildTime)" \
	    -o ./build/server
	@echo "[OK] App binary was created!"

build-testnet: generate security
	@go build \
	    -tags testnet \
		-ldflags "-X main.Branch=$(Branch) -X main.Commit=$(Commit) -X main.BuildTime=$(BuildTime)" \
		-o ./build/server-testnet
	@echo "[OK] App binary(testnet) was created!"

run:
	@./build/server
