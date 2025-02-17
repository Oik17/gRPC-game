package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Oik17/gRPC-game/config"
	database "github.com/Oik17/gRPC-game/db/init"
	db "github.com/Oik17/gRPC-game/db/sqlc"
	proto "github.com/Oik17/gRPC-game/gen"
	"github.com/Oik17/gRPC-game/pkg"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	dsn := config.Config("DSN")
	pool, err := database.InitDB(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer pool.Close()

	sqlDB := stdlib.OpenDBFromPool(pool)

	queries := db.New(sqlDB)
	_ = queries

	grpcServer := grpc.NewServer()
	poolService := &pkg.Pool{}

	proto.RegisterGameServiceServer(grpcServer, poolService)

	fmt.Println("Game gRPC server is running on port 8080...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
