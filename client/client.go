package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/BlazeCode1/ProtoM/proto/gen/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Could not connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new ChatService client
	client := pb.NewChatServiceClient(conn)

	// Send a message
	msgReq := &pb.SendMessageRequest{
		Message: &pb.ChatMessage{
			SenderId:   "user1",
			ReceiverId: "user2",
			Content:    "Hello, how are you?",
		},
	}
	msgResp, err := client.SendMessage(context.Background(), msgReq)
	if err != nil {
		log.Fatalf("❌ Error sending message: %v", err)
	}
	fmt.Printf("Message sent! ID: %s\n", msgResp.MessageId)

	// Stream messages
	streamReq := &pb.StreamMessagesRequest{UserId: "user2"}
	stream, err := client.StreamMessages(context.Background(), streamReq)
	if err != nil {
		log.Fatalf("❌ Error streaming messages: %v", err)
	}

	fmt.Println("Streaming messages for user2:")
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("❌ Error receiving message: %v", err)
		}
		fmt.Printf("Received: %s from %d: %s\n", msg.SenderId, msg.Timestamp, msg.Content)
		time.Sleep(2 * time.Second)
	}
}
