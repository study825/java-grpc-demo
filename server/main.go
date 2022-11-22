package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/pb"
)

const (
	port = ":7080"
)

type server struct { //服务的结构类型
	*pb.UnimplementedHelloServiceServer
}

func (s *server) hello(ctx context.Context, in *pb.HelloRequest)(*pb.HelloResponse,error) {
	return &pb.HelloResponse{Greeting: "1"}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	log.Printf("start gRPC listener on port" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
