package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "rpc-tutorial/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAPIClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	a := &pb.Item{Title: "First", Body: "A first item"}
	b := &pb.Item{Title: "Second", Body: "A second item"}
	c := &pb.Item{Title: "Third", Body: "A third item"}

	// Add items
	for _, item := range []*pb.Item{a, b, c} {
		_, err := client.AddItem(ctx, item)
		if err != nil {
			log.Fatalf("AddItem error: %v", err)
		}
	}

	// Get DB
	dbResp, err := client.GetDB(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("GetDB error: %v", err)
	}
	fmt.Println("Database:", dbResp.Items)

	// Edit item
	_, err = client.EditItem(ctx, &pb.Item{Title: "Second", Body: "A new second item"})
	if err != nil {
		log.Fatalf("EditItem error: %v", err)
	}

	// Delete item
	_, err = client.DeleteItem(ctx, c)
	if err != nil {
		log.Fatalf("DeleteItem error: %v", err)
	}

	// Get DB again
	dbResp, err = client.GetDB(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("GetDB error: %v", err)
	}
	fmt.Println("Database after edits:", dbResp.Items)

	// Get by name
	itemResp, err := client.GetByName(ctx, &pb.Item{Title: "First"})
	if err != nil {
		log.Fatalf("GetByName error: %v", err)
	}
	fmt.Println("First item:", itemResp)
}
