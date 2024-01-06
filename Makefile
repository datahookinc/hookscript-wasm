.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o hookscript.wasm

.PHONY: build-wasi
build-wasi:
	GOOS=wasip1 GOARCH=wasm go build -ldflags="-s -w" -o hookscript.wasm

.PHONY: prod-build
prod-build:
	go mod edit -dropreplace github.com/datahookinc/hookscript
	go mod tidy
	make build
	
.PHONY: dev-build
dev-build:
	go mod edit -replace github.com/datahookinc/hookscript=../hookscript
	go mod tidy
	make build