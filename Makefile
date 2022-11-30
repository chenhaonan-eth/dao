#生成Gormgen CURD代码
generate-gorm-gen-code:
	@echo "Start Generating"
	go run ./cmd/gormgen
.PHONY: generate-gorm-gen-code

#gRPC service 
generate-buf:
	buf generate --path ./proto/server/server.proto
.PHONY: generate-buf

clean-buf:
	rm proto/server/*.go
	rm proto/server/*.json
.PHONY: clean-buf

install:
	go install \
		github.com/bufbuild/buf/cmd/buf@lastest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@lastest \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@lastest \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@lastest \
		google.golang.org/protobuf/cmd/protoc-gen-go@lastest

run:
	go run main.go server 