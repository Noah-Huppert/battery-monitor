.PHONY: run imports fmt

# Source files
SRC=main.go auth config db models api

# Runs the server
run:
	go run main.go

# Adds all required go imports
imports:
	goimports -l -w ${SRC}

# Formats go source
fmt:
	gofmt -w ${SRC}
