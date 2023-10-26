package main

import (
	"context"
	"log"
	"net"

	pb "github.com/charafzellou/grpc-golang-template/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type RequestServer struct {
	pb.UnimplementedRequestServiceServer
}

func (s *RequestServer) CreateTodo(ctx context.Context, in *pb.InputRequest) (*pb.OutputRequest, error) {
	log.Printf("Received request: %v", in.GetName())
	todo := &pb.OutputRequest{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterRequestServiceServer(s, &RequestServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve server: %v", err)
	}
}
