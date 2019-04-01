package main

import (
	"context"
	pb "github.com/zzell/pingpong/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"time"
)

const (
	addr = "127.0.0.1:8040"
)

var ctx, _ = context.WithCancel(context.Background())

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := pb.NewPingpongClient(conn)
	stream, err := client.Play(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// streaming read from db
    pipe := dbRead()
	go func() {
		defer func() {
			err := stream.CloseSend()
			if err != nil {
				log.Printf("failed to close stream: %s", err)
			}
		}()

		for chunk := range pipe {
			ping := pb.Ping{
				Msg:  "ping",
				Data: chunk,
			}

			err := stream.Send(&ping)
			if err != nil {
				log.Fatalf("failed to send: %s", err)
			}

			log.Println("sent ping...")
		}
	}()

    done := make(chan bool)
	go func() {
		for {
			pong, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					done <- true
					return
				} else {
					log.Printf("failed to receive: %s", err)
				}
			}

			// server received chunk, everything ok
			log.Printf("received %s", pong.Msg)
		}
	}()

    <-done
}

func dbRead() chan []byte {
	var ch = make(chan []byte)

    // dbRead from db
	go func() {
		defer close(ch)
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Second)

			chunk := make([]byte, 10)
			rand.Read(chunk)

			ch <- chunk
		}
	}()

	return ch
}