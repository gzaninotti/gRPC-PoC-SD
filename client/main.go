package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "pucsp/sd/pocgrpc/poc-grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Hello world!"
)

var (
	addr = flag.String("addr", "localhost:50051", "address to connect")
	msg  = flag.String("msg", defaultName, "User message")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewBroadcasterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Broadcast(ctx, &pb.BroadcastReq{Message: *msg})
	if err != nil {
		log.Fatalf("could not broadcast this message: %v", err)
	}
	log.Printf("Msg: %s", resp.GetMessage())

}
