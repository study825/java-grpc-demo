package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "service/ecommerce"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("start gRPC listener on port" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
