package main

import (
	"log"

	pb "github.com/chenhaonan-eth/dao/proto/server"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	context := context.Background()
	body := &pb.HelloRequest{
		Name: "Grpc",
	}

	r, err := c.SayHello(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
