package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/Oik17/gRPC-game/gen"
	"github.com/Oik17/gRPC-game/pkg"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Initialize the Pool (connection manager)
	pool := &pkg.Pool{}

	// Register the game service server
	proto.RegisterGameServiceServer(grpcServer, pool)

	fmt.Println("Game gRPC server is running on port 8080...")

	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
