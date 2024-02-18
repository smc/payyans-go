GOROOT=$(shell go env GOROOT)

CLI_BIN := payyanscli

.PHONY: cli
cli:
	go build -o ${CLI_BIN} -ldflags "-s -w" ./cli.go

test:
	go test payyans/*.go

wasm:
	GOOS=js GOARCH=wasm go build -o web/payyans.wasm web/wasm.go
	cp "${GOROOT}/misc/wasm/wasm_exec.js" web/

.PHONY: web
web:
	$(MAKE) wasm
	cd web && go run build_index_html.go
