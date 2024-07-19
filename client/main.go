package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/charafzellou/grpc-golang-template/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	method := flag.String("method", "register", "You can use the following RPC methods : register, subscribe, getlastblock, addtransaction, bakeblock, confirmbake")
	uuid := flag.String("uuid", "", "UUID of the client")
	flag.Parse()
	log.Printf("Calling %s on '35.241.224.46:50051'", *method)

	conn, err := grpc.NewClient("35.241.224.46:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlockchainClient(conn)
	log.Printf("Connected to server '35.241.224.46:50051'")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	switch *method {
	case "register":
		res, err := client.Register(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not register: %v", err)
		}
		log.Printf("Registered with UUID: %s, Reputation: %d", res.GetUuid(), res.GetReputation())
	case "subscribe":
		res, err := client.Subscribe(ctx, &pb.SubscribeRequest{Uuid: *uuid})
		if err != nil {
			log.Fatalf("could not subscribe: %v", err)
		}
		log.Printf("Subscription response: %s", res.GetMessage())
	case "getlastblock":
		res, err := client.GetLastBlock(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not get last block: %v", err)
		}
		log.Printf("Last block info: %v", res)
	case "addtransaction":
		res, err := client.AddTransaction(ctx, &pb.Transaction{
			Sender:   "sender_uuid",
			Receiver: "receiver_uuid",
			Amount:   10,
			Data:     "transaction_data",
		})
		if err != nil {
			log.Fatalf("could not add transaction: %v", err)
		}
		log.Println("Transaction added successfully: ", res)
	case "bakeblock":
		res, err := client.BakeBlock(ctx, &pb.BakeRequest{Uuid: *uuid})
		if err != nil {
			log.Fatalf("could not bake block: %v", err)
		}
		log.Printf("Bake block response: %s", res.GetMessage())
	case "confirmbake":
		_, err := client.ConfirmBake(ctx, &pb.ConfirmRequest{Uuid: *uuid})
		if err != nil {
			log.Fatalf("could not confirm bake: %v", err)
		}
		log.Println("Block baking confirmed")
	default:
		log.Fatalf("Unknown method: %s", *method)
	}
}
