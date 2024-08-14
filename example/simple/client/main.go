package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/tokopedia/gripmock/protogen/example/simple"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Set up a connection to the server.
	conn, err := grpc.DialContext(ctx, "localhost:4770", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGripmockClient(conn)

	// Contact the server and print out its response.
	name := "tokopedia"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}
	log.Printf("Greeting: %s (return code %d)", r.Message, r.ReturnCode)

	name = "world"
	r, err = c.SayHello(context.Background(), &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}
	log.Printf("Greeting: %s (return code %d)", r.Message, r.ReturnCode)
	name = "error"
	r, err = c.SayHello(context.Background(), &pb.Request{Name: name})
	if err == nil {
		log.Fatalf("Expected error, but return %d", r.ReturnCode)
	}
	log.Printf("Greeting error: %s", err)
	name = "error_code"
	r, err = c.SayHello(context.Background(), &pb.Request{Name: name})
	if err == nil {
		log.Fatalf("Expected error, but return %d", r.ReturnCode)
	}
	log.Printf("Greeting error: %s", err)
}
