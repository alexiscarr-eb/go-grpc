package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alexiscarr-eb/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting calcualtor...")

	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to calculator service: %v", err)
	}

	defer client.Close()

	sumStub := calcpb.NewSumServiceClient(client)

	getSum(sumStub)
}

func getSum(sumStub calcpb.SumServiceClient) {
	fmt.Println("Calculating sum...")

	req := &calcpb.SumRequest{
		Num_1: 123,
		Num_2: 321,
	}

	res, err := sumStub.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error occurred while calling Sum action: %v", err)
	}

	log.Printf("Sum: %v", res.Result)
}
