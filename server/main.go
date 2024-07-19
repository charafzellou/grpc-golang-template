package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/charafzellou/grpc-golang-template/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBlockchainServiceServer
	clients     map[string]*Client
	clientsLock sync.Mutex
}

type Client struct {
	UUID            string
	ReputationScore float64
	BakingStream    pb.BlockchainService_SubscribeForBakingServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	blockchainServer := &server{
		clients: make(map[string]*Client),
	}
	pb.RegisterBlockchainServiceServer(s, blockchainServer)

	go selectBakerPeriodically(blockchainServer)

	fmt.Println("Server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) RegisterClient(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	uuid := generateUUID()
	s.clientsLock.Lock()
	s.clients[uuid] = &Client{UUID: uuid, ReputationScore: 0}
	s.clientsLock.Unlock()
	return &pb.RegisterResponse{Uuid: uuid, ReputationScore: 0}, nil
}

func (s *server) SubscribeForBaking(req *pb.SubscribeRequest, stream pb.BlockchainService_SubscribeForBakingServer) error {
	s.clientsLock.Lock()
	client, exists := s.clients[req.ClientUuid]
	if !exists {
		s.clientsLock.Unlock()
		return fmt.Errorf("client not registered")
	}
	client.BakingStream = stream
	s.clientsLock.Unlock()

	// Keep the stream open
	<-stream.Context().Done()
	return nil
}

func (s *server) GetLastBlockInfo(ctx context.Context, req *pb.LastBlockRequest) (*pb.LastBlockResponse, error) {
	// Implement fetching last block info
	return &pb.LastBlockResponse{BlockHash: "sample-hash", BlockHeight: 100}, nil
}

func (s *server) AddTransaction(ctx context.Context, req *pb.TransactionData) (*pb.TransactionResponse, error) {
	// Implement adding transaction to mempool
	return &pb.TransactionResponse{Success: true, TransactionHash: "sample-tx-hash"}, nil
}

func selectBakerPeriodically(s *server) {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		s.clientsLock.Lock()
		if len(s.clients) == 0 {
			s.clientsLock.Unlock()
			continue
		}
		randomIndex := rand.Intn(len(s.clients))
		i := 0
		var chosenClient *Client
		for _, client := range s.clients {
			if i == randomIndex {
				chosenClient = client
				break
			}
			i++
		}
		s.clientsLock.Unlock()

		if chosenClient != nil && chosenClient.BakingStream != nil {
			err := chosenClient.BakingStream.Send(&pb.BakingUpdate{ChosenAsBaker: true})
			if err != nil {
				log.Printf("Failed to send baking update to client %s: %v", chosenClient.UUID, err)
			} else {
				// Implement logic to wait for confirmation and update reputation
			}
		}
	}
}

func generateUUID() string {
	// Implement UUID generation
	return "sample-uuid"
}
