package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hakuna86/grpc-sample/protos/chatsample"
	"github.com/hakuna86/grpc-sample/protos/common"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func chatSample(client chatsample.RpcServiceClient) {
	notes := []*common.ChatMessage{
		{Code: 1, Message: "message 001"},
		{Code: 2, Message: "message 002"},
		{Code: 3, Message: "message 003"},
		{Code: 4, Message: "message 004"},
		{Code: 5, Message: "message 005"},
		{Code: 6, Message: "message 006"},
		{Code: 7, Message: "message 007"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("%v.chatSample(_) = _, %v", client, err)
	}
	defer stream.CloseSend()

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at code %d)", in.Message, in.Code)
		}
	}()

	//loop:
	//	for {
	//		log.Println("Send message attempts")
	//		for _, note := range notes {
	//			if err := stream.Send(note); err != nil {
	//				log.Fatalf("Failed to send a note: %v", err)
	//			}
	//		}
	//
	//		time.Sleep(1*time.Second)
	//
	//		select {
	//		case <-waitc:
	//			log.Println("waitc message drived")
	//			break loop
	//		}
	//	}

	for {
		log.Println("Send message attempts")
		for _, note := range notes {
			if err := stream.Send(note); err != nil {
				log.Fatalf("Failed to send a note: %v", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func puchSample(client chatsample.RpcServiceClient) {
	ctx := context.Background()
	stream, err := client.Push(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("%v.puchSample(_) = _, %v", client, err)
	}
	defer stream.CloseSend()

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message make block at code %t)", in.Code == chatsample.Process_MAKE_BLOCK)
		}
	}()
	<-waitc
}

func main() {
	waitc := make(chan struct{})
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := chatsample.NewRpcServiceClient(conn)
	go chatSample(c)
	go puchSample(c)

	<-waitc

}
