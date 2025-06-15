package main

import (
	"context"
	"log"
	"net"

	pb "rpc-tutorial/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAPIServer
	database []*pb.Item
}

func (s *server) GetDB(ctx context.Context, _ *pb.Empty) (*pb.ItemList, error) {
	return &pb.ItemList{Items: s.database}, nil
}

func (s *server) GetByName(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	for _, it := range s.database {
		if it.Title == item.Title {
			return it, nil
		}
	}
	return &pb.Item{}, nil // or return error if not found
}

func (s *server) AddItem(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	s.database = append(s.database, item)
	return item, nil
}

func (s *server) EditItem(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	for i, it := range s.database {
		if it.Title == item.Title {
			s.database[i] = item
			return item, nil
		}
	}
	return &pb.Item{}, nil
}

func (s *server) DeleteItem(ctx context.Context, item *pb.Item) (*pb.Item, error) {
	for i, it := range s.database {
		if it.Title == item.Title && it.Body == item.Body {
			s.database = append(s.database[:i], s.database[i+1:]...)
			return item, nil
		}
	}
	return &pb.Item{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAPIServer(grpcServer, &server{})

	log.Println("gRPC server listening on :4040")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
