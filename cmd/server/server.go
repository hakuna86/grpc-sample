package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hakuna86/grpc-sample/protos/chatsample"
	"github.com/hakuna86/grpc-sample/protos/common"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

//go:generate protoc -I $PWD/protos --go_out=plugins=grpc:$PWD/protos $PWD/protos/chatsample/*.proto
//go:generate protoc -I $PWD/protos --go_out=plugins=grpc:$PWD/protos $PWD/protos/common/*..proto

type server struct {
	mu sync.Mutex
}

// realtime chatting
func (s *server) Chat(stream chatsample.RpcService_ChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.mu.Lock()
		s.mu.Unlock()

		log.Println("server received", in.Code, in.Message)

		received := new(common.ChatMessage)
		received.Code = in.Code + 1
		received.Message = "message pong"

		if err := stream.Send(received); err != nil {
			return err
		}
	}
}

func (s *server) Push(empty *empty.Empty, stream chatsample.RpcService_PushServer) error {

	msg := chatsample.Code{
		Code: chatsample.Process_MAKE_BLOCK,
	}
	for {
		log.Println("Send message")
		if err := stream.Send(&msg); err != nil {
			log.Println("server puch rpc errror", err)
			return err
		}

		time.Sleep(1 * time.Second)
	}
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chatsample.RegisterRpcServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
