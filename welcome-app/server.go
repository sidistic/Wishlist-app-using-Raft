package main

import (
	"fmt"
	"log"
	"net"
	"welcome-app/login"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen to login: %v", err)
	}

	l := login.Server{}

	lServer := grpc.NewServer()

	login.RegisterAuthServiceServer(lServer, &l)

	if err := lServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve for login: %s", err)
	}

}
