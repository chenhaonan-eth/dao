protoc:
	protoc  --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./proto/*.proto