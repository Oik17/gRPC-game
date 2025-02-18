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
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	// queries, err := database.InitDB(context.Background())
	// if err != nil {
	// 	log.Fatalf("Database connection failed: %v", err)
	// }

	// _ = queries

	grpcServer := grpc.NewServer()
	poolService := &pkg.Pool{}

	proto.RegisterGameServiceServer(grpcServer, poolService)

	fmt.Println("Game gRPC server is running on port 8080...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
