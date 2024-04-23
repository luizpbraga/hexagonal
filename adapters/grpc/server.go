package rpc

import (
	"context"
	"log"
	"net"

	"github.com/luizpbraga/hexagonal/adapters/grpc/pb"
	"github.com/luizpbraga/hexagonal/ports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// THE GRPCService implements pb.ArithmeticServiceServer
type GRPCService struct {
	api ports.API
}

func NewGRPCService(api ports.API) *GRPCService {
	return &GRPCService{api: api}
}

// Registers the grpc Server
func (service *GRPCService) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, service)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (service *GRPCService) GetAddition(ctx context.Context, params *pb.OperationParameters) (*pb.Answer, error) {
	answe, err := service.api.GetAddition(params.A, params.B)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.Answer{Value: answe}, err
}

func (service *GRPCService) GetMultiplication(ctx context.Context, params *pb.OperationParameters) (*pb.Answer, error) {
	answe, err := service.api.GetMultiplication(params.A, params.B)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.Answer{Value: answe}, err
}

func (service *GRPCService) GetSubtraction(ctx context.Context, params *pb.OperationParameters) (*pb.Answer, error) {
	answe, err := service.api.GetSubtraction(params.A, params.B)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.Answer{Value: answe}, err
}

func (service *GRPCService) GetDivision(ctx context.Context, params *pb.OperationParameters) (*pb.Answer, error) {
	answe, err := service.api.GetDivision(params.A, params.B)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.Answer{Value: answe}, err
}
