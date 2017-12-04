.PHONY: run imports fmt

# Source files
SRC=main.go models

# Runs the server
run:
	go run main.go

# Adds all required go imports
imports:
	goimports -l -w ${SRC}

# Formats go source
fmt:
	gofmt -w ${SRC}
