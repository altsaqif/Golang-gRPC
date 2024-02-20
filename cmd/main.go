package main

import (
	"log"
	"net"

	"github.com/altsaqif/go-grpc/cmd/config"
	"github.com/altsaqif/go-grpc/cmd/services"
	productPb "github.com/altsaqif/go-grpc/github.com/altsaqif/go-grpc/pb/product"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listened %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server Started at %v", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to Serve %v", err.Error())
	}
}
