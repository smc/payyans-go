CLI_BIN := payyanscli

.PHONY: cli
cli:
	go build -o ${CLI_BIN} -ldflags "-s -w" ./cli.go

test:
	go test payyans/*.go
