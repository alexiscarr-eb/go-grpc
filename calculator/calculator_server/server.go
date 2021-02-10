package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alexiscarr-eb/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum action was invoked with %v", req)

	num1 := req.GetNum_1()
	num2 := req.GetNum_2()

	sum := num1 + num2

	result := fmt.Sprintf("%v + %v = %v", num1, num2, sum)

	res := &calcpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Calculator service running...")

	lis, err := net.Listen("tcp", "0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
