#gRPC service 
protoc:
	protoc  --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./proto/*.proto

#生成Gormgen CURD代码
generate-gorm-gen-code:
	@echo "Start Generating"
	go run ./cmd/gormgen
.PHONY: generate-gorm-gen-code