package main

import (
    "context"
    "fmt"
    "log"

    pb "grpc-demo/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    c := pb.NewGreeterClient(conn)

    resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp.Message) // Hello, World!
}