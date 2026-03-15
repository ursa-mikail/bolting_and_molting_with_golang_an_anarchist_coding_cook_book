package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "grpc-demo/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})

    fmt.Println("Server listening on :50051")
    log.Fatal(s.Serve(lis))
}