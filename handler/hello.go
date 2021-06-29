package handler

import (
	"context"
	"log"

	pb "github.com/quyenphamkhac/skoppi/proto"
)

type HelloSrv struct {
	pb.UnimplementedGreeterServer
}

func (e *HelloSrv) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}
