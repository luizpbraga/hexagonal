package client

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/luizpbraga/hexagonal/adapters/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String("addr", "localhost:9000", "The server address in the format of host:port")

func clientTest() {
	// auth credentials (for example, TLS, GCE credentials, or JWT credentials
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("[DIAL FAIL] did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addRequest := &pb.OperationParameters{A: 1, B: 0}
	addResponse, err := client.GetDivision(ctx, addRequest)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Addition: %d + %d = %d\n", addRequest.A, addRequest.B, addResponse.Value)

	divRequest := &pb.OperationParameters{A: 10, B: 2}
	divResponse, err := client.GetDivision(ctx, divRequest)
	if err != nil {
		log.Fatalf("could not get division result: %v", err)
	}
	fmt.Printf("Division: %d / %d = %d\n", divRequest.A, divRequest.B, divResponse.Value)

	subRequest := &pb.OperationParameters{A: 10, B: 5}
	subResponse, err := client.GetSubtraction(ctx, subRequest)
	if err != nil {
		log.Fatalf("could not get subtraction result: %v", err)
	}
	fmt.Printf("Subtraction: %d - %d = %d\n", subRequest.A, subRequest.B, subResponse.Value)

	mulRequest := &pb.OperationParameters{A: 4, B: 6}
	mulResponse, err := client.GetMultiplication(ctx, mulRequest)
	if err != nil {
		log.Fatalf("could not get multiplication result: %v", err)
	}
	fmt.Printf("Multiplication: %d * %d = %d\n", mulRequest.A, mulRequest.B, mulResponse.Value)
}
