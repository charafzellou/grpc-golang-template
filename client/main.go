package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/charafzellou/grpc-golang-template/proto"
)

const (
	ADDRESS = "localhost:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {
	conn, err := grpc.NewClient(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}

	c := pb.NewRequestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	todos := []TodoTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
		{Name: "Make YouTube Video", Description: "Start Go for beginners series", Done: false},
		{Name: "Go to the gym", Description: "Leg day", Done: false},
		{Name: "Buy groceries", Description: "Buy tomatoes, onions, mangos", Done: false},
		{Name: "Meet with mentor", Description: "Discuss blockers in my project", Done: false},
	}

	for _, todo := range todos {
		res, err := c.RequestMethod(ctx, &pb.InputRequest{Name: todo.Name, Description: todo.Description, Done: todo.Done})
		if err != nil {
			log.Fatalf("Could not create user: \n%v", err)
		}

		log.Printf(`
           ID : %s
           Name : %s
           Description : %s
           Done : %v,
       `, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}

	defer conn.Close()
	defer cancel()
}
