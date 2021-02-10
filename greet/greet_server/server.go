package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/alexiscarr-eb/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

/*
func Greet on pointer to server
args: ctx context.Context, req *greetpb.GreetRequest
return type *greetpb.GreetResponse, error
implementing GreetServiceServer interface from greet.pb.go

	type GreetServiceServer interface {
		// unary API
		Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	}
*/
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)

	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
