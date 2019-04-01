package main

import (
	"io"
	"log"
	"net"

	pb "github.com/zzell/pingpong/proto"
	"google.golang.org/grpc"
)

const port = ":8040"

type server struct{}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen failed: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingpongServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s server) Play(srv pb.Pingpong_PlayServer) error {
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				log.Println(err)
				continue
			}
		}

		log.Printf("received %s", req.Msg)
		log.Printf("data: %v", req.Data)

		pong := pb.Pong{
			Msg: "Pong",
		}

		err = srv.Send(&pong)
		if err != nil {
			log.Printf("failed to send: %s", err)
		}
	}
}
