package main

import (
	"banking_with_grpc/netxd_customer_connectors/constants"
	"context"
	"fmt"

	"log"

	pb "banking_with_grpc/netxd_customer"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.Details{
		CustomerId: 320,
		Firstname:  "Jsri",
		Lastname:   "R",
		BankId:     12345,
		Balance:    100000,
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}

	fmt.Printf("CustomerID: %d\nCreatedTime:%v\n", response.CustomerId,response.CreatedAt)
}