all:
	protoc --go-grpc_out=require_unimplemented_servers=false:adapters/grpc --proto_path=adapters/grpc/proto adapters/grpc/proto/arithmetics_svc.proto
	protoc --go_out=adapters/grpc --proto_path=adapters/grpc/proto adapters/grpc/proto/number_msg.proto
