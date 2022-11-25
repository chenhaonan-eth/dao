package server

import (
	"context"
	"fmt"

	pb "github.com/chenhaonan-eth/dao/proto/server"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("SayHello")
	return &pb.HelloReply{
		Message: "test",
	}, nil
}
