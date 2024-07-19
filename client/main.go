package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/charafzellou/grpc-golang-template/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("server_addr", "localhost:50051", "The server address in the format of host:port")
	action     = flag.String("action", "", "Action to perform: register, subscribe, getlastblock, addtransaction")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlockchainServiceClient(conn)

	ctx := context.Background()

	switch *action {
	case "register":
		registerClient(ctx, client)
	case "subscribe":
		subscribeForBaking(ctx, client)
	case "getlastblock":
		getLastBlockInfo(ctx, client)
	case "addtransaction":
		addTransaction(ctx, client)
	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}

func registerClient(ctx context.Context, client pb.BlockchainServiceClient) {
	resp, err := client.RegisterClient(ctx, &pb.RegisterRequest{})
	if err != nil {
		log.Fatalf("Failed to register: %v", err)
	}
	fmt.Printf("Registered with UUID: %s, Reputation Score: %.2f\n", resp.Uuid, resp.ReputationScore)
}

func subscribeForBaking(ctx context.Context, client pb.BlockchainServiceClient) {
	stream, err := client.SubscribeForBaking(ctx, &pb.SubscribeRequest{ClientUuid: "your-uuid-here"})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	for {
		update, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive update: %v", err)
		}
		if update.ChosenAsBaker {
			fmt.Println("Chosen as baker! Confirming...")
			// Implement confirmation logic here
		}
	}
}

func getLastBlockInfo(ctx context.Context, client pb.BlockchainServiceClient) {
	resp, err := client.GetLastBlockInfo(ctx, &pb.LastBlockRequest{})
	if err != nil {
		log.Fatalf("Failed to get last block info: %v", err)
	}
	fmt.Printf("Last Block - Hash: %s, Height: %d\n", resp.BlockHash, resp.BlockHeight)
}

func addTransaction(ctx context.Context, client pb.BlockchainServiceClient) {
	resp, err := client.AddTransaction(ctx, &pb.TransactionData{
		Sender:    "sender-address",
		Recipient: "recipient-address",
		Amount:    100.0,
	})
	if err != nil {
		log.Fatalf("Failed to add transaction: %v", err)
	}
	fmt.Printf("Transaction added - Success: %v, Hash: %s\n", resp.Success, resp.TransactionHash)
}
