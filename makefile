all:
  protoc --go_out=adapters/grpc --proto_path=adapters/grpc/proto adapters/grpc/proto/number_msg.proto
  protoc --go-grpc_out=adapters/grpc --proto_path=adapters/grpc/proto adapters/grpc/proto/arithmetics_srv.proto
