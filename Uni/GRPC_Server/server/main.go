package main

import (
	"log"
	"net"

	pb "../api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// "github.com/tutorialedge/go-grpc-tutorial/api"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) ReverseString(ctx context.Context, in *api.ReverseRequest) (*api.ReverseReply, error) {
	n := 0
	rune := make([]rune, len(in.Data))
	for _, r := range in.Data {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	return &api.ReverseReply{Reversed: string(rune)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterReverseServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}