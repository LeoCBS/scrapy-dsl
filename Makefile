.PHONY: repl
check:
	go test ./...

repl:
	go run main.go
