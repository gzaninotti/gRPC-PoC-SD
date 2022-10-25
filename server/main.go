package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "pucsp/sd/pocgrpc/poc-grpc"
	"strings"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedBroadcasterServer
}

// Broadcast -> Broadcasts a message
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastReq) (*pb.BroadcastRes, error) {
	log.Printf("Got: %v", req.GetMessage())

	return &pb.BroadcastRes{Message: strings.Title(req.GetMessage())}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBroadcasterServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
