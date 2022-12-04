PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
CONFIG=$(GOBASE)/config.yaml

# Redirect error output to a file, so we can show it in development mode.
STDERR=~/.$(PROJECTNAME)-stderr.txt
# PID file will store the server process id when it's running on development mode
PID=~/.$(PROJECTNAME)-api-server.pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent
.PHONY: help
all: help

## start: Start server.
start: start-server

## stop: Stop server.
stop: stop-server

## restart: restart server.
restart: restart-server

## install: Install missing dependencies. Runs `go get` internally.
install: go-get

## build: build server
build: go-build

## generate: generate Gormgen CURD and gRPC code
generate: generate-gorm-gen-code generate-buf

#生成Gormgen CURD代码
generate-gorm-gen-code:
	@echo "Start Generating"
	go run $(GOBASE)/cmd/gormgen

#gRPC service 
generate-buf:
	buf generate --path $(GOBASE)/proto/server/server.proto

clean-buf:
	rm $(GOBASE)/proto/server/*.go
	rm $(GOBASE)/proto/server/*.json

install-tools:
	go install \
		github.com/bufbuild/buf/cmd/buf@lastest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@lastest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@lastest \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@lastest \
		google.golang.org/protobuf/cmd/protoc-gen-go@lastest

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

go-build:
	@echo "  >  Building binary..."
	@go build -o $(GOBIN)/$(PROJECTNAME)

start-server:
	@echo " > Start $(PROJECTNAME)"
	@$(GOBIN)/$(PROJECTNAME) server -c $(CONFIG) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@echo " > Stop server"
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

restart-server: stop-server start-server

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo