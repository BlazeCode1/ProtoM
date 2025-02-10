package main

import (
	"fmt"
	"log"
	"net"
	"github.com/BlazeCode1/ProtoM/app/controller"
	"github.com/BlazeCode1/ProtoM/app/service"
	pb "github.com/BlazeCode1/ProtoM/proto/gen/chat"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	// Initialize services
	messageService := service.NewMessageService()
	chatController := controller.NewChatController(messageService)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, chatController)

	fmt.Println("🚀 gRPC Server running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
