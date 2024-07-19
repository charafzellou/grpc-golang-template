package main

import (
	"context"
	"crypto"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/charafzellou/grpc-golang-template/proto"
	"google.golang.org/grpc"
)

const (
	portNumber = ":50051"
)

type Block struct {
	Number            int32
	PreviousBlockHash string
	BlockHash         string
	Data              string
	Transactions      []*pb.Transaction
}

type server struct {
	pb.UnimplementedBlockchainServer
	mu           sync.Mutex
	clients      map[string]int32
	currentBaker string
	blocks       []*Block
	transactions []*pb.Transaction
}

func (s *server) Register(ctx context.Context, in *pb.Empty) (*pb.RegisterResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	uuid := generateUUID()
	reputation := int32(100)
	s.clients[uuid] = reputation
	log.Printf("Client %s registered with reputation %d", uuid, reputation)
	return &pb.RegisterResponse{Uuid: uuid, Reputation: reputation}, nil
}

func (s *server) Subscribe(ctx context.Context, in *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.clients[in.Uuid]; !exists {
		log.Printf("Client %s not registered: ", in.Uuid)
		return &pb.SubscribeResponse{Message: "Client not registered"}, nil
	}
	log.Printf("Client %s subscribed: ", in.Uuid)
	return &pb.SubscribeResponse{Message: "Subscribed successfully!"}, nil
}

func (s *server) GetLastBlock(ctx context.Context, in *pb.Empty) (*pb.BlockInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.blocks) == 0 {
		return &pb.BlockInfo{}, nil
	}
	lastBlock := s.blocks[len(s.blocks)-1]
	log.Printf("Last block info: %v", lastBlock)
	return &pb.BlockInfo{
		BlockHash:         lastBlock.BlockHash,
		PreviousBlockHash: lastBlock.PreviousBlockHash,
		BlockNumber:       lastBlock.Number,
		Data:              lastBlock.Data,
	}, nil
}

func (s *server) AddTransaction(ctx context.Context, in *pb.Transaction) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions = append(s.transactions, in)
	log.Printf("Transaction added: %v", in)
	return &pb.Empty{}, nil
}

func (s *server) BakeBlock(ctx context.Context, in *pb.BakeRequest) (*pb.BakeResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.currentBaker != "" {
		return &pb.BakeResponse{Uuid: s.currentBaker, Message: "Baking in progress..."}, nil
	}
	if _, exists := s.clients[in.Uuid]; !exists {
		return &pb.BakeResponse{Message: "Client not registered"}, nil
	}
	s.currentBaker = in.Uuid
	log.Printf("Client %s is baking a block...", in.Uuid)
	return &pb.BakeResponse{Uuid: s.currentBaker, Message: "You have been chosen to bake a block !"}, nil
}

func (s *server) ConfirmBake(ctx context.Context, in *pb.ConfirmRequest) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if in.Uuid == s.currentBaker {
		s.clients[in.Uuid]++
		s.currentBaker = ""
		s.mineBlock()
		log.Printf("Client %s confirmed block bake", in.Uuid)
	} else {
		s.clients[in.Uuid]--
		log.Printf("Client %s is not the chosen baker", in.Uuid)
	}
	return &pb.Empty{}, nil
}

func (s *server) mineBlock() {
	lastBlock := s.blocks[len(s.blocks)-1]
	newBlock := &Block{
		Number:            lastBlock.Number + 1,
		PreviousBlockHash: lastBlock.BlockHash,
		BlockHash:         string(crypto.SHA256.New().Sum([]byte(generateUUID()))),
		Data:              string(len(s.transactions)),
		Transactions:      s.transactions,
	}
	s.blocks = append(s.blocks, newBlock)
	s.transactions = nil
	log.Printf("Block mined: %v", newBlock)
}

func generateUUID() string {
	return fmt.Sprintf("%d", rand.Int63())
}

func main() {
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on port 50051")
	s := grpc.NewServer()
	initialBlock := &Block{
		Number:            0,
		PreviousBlockHash: "",
		BlockHash:         generateUUID(),
		Data:              "Genesis block",
		Transactions:      []*pb.Transaction{},
	}
	srv := &server{
		clients: make(map[string]int32),
		blocks:  []*Block{initialBlock},
	}
	pb.RegisterBlockchainServer(s, srv)
	log.Println("Server started successfully")
	go func() {
		for {
			time.Sleep(30 * time.Second)
			srv.mu.Lock()
			if srv.currentBaker != "" {
				log.Println("Backer has been selected: ", srv.currentBaker)
				srv.mineBlock()
			}
			srv.mu.Unlock()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
