OS_NAME := $(shell uname -s)

all: deps fmt vet build

deps: FORCE
# install ragel if not exist
ifeq (, $(shell which ragel))
ifeq ($(OS_NAME),Darwin)
		@echo "brew install antlr"
		@brew install antlr
else
	@echo "dnf install -y antlr"
	@dnf install -y antlr
endif
endif
	@echo "go mod tidy"
	@go mod tidy

antlr-go: FORCE
	@echo "antlr -Dlanguage=Go GraphObj.g4"
	@antlr -Dlanguage=Go GraphObj.g4 -o parser/

fmt: FORCE
	@echo "Formatting code"
	@go list -f '{{.Dir}}' ./  | xargs -L1 gofmt -l

vet:
	@echo "Vetting code"
	@go vet ./.

build: fmt
	go build -o ./bin/graphobj

clean: FORCE
	@echo "Cleaning"
	rm -rf ./bin

FORCE: