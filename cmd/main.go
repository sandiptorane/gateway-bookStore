package main

import (
	books "gateway-bookstore/bookproto"
	"gateway-bookstore/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	addr := ":8081"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	s := server.NewServer()
	s.InitializeServer()
	books.RegisterBookServiceServer(grpcServer, s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	err = server.Run(addr)
	log.Fatalln(err)
}