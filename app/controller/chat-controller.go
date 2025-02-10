package controller

import (
	"context"

	"github.com/BlazeCode1/ProtoM/app/model"
	"github.com/BlazeCode1/ProtoM/app/service"
	pb "github.com/BlazeCode1/ProtoM/proto/gen/chat"
)

type ChatController struct {
	pb.UnimplementedChatServiceServer
	service *service.MessageService
}

func NewChatController(service *service.MessageService) *ChatController {
	return &ChatController{service: service}
}

func (c *ChatController) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg := &model.ChatMessage{
		SenderID:   req.Message.SenderId,
		ReceiverID: req.Message.ReceiverId,
		Content:    req.Message.Content,
	}
	msgID := c.service.SendMessage(msg)
	return &pb.SendMessageResponse{Success: true, MessageId: msgID}, nil
}

func (c *ChatController) StreamMessages(req *pb.StreamMessagesRequest, stream pb.ChatService_StreamMessagesServer) error {
	messages := c.service.StreamMessages(req.UserId)
	for _, msg := range messages {
		err := stream.Send(&pb.ChatMessage{
			SenderId:   msg.SenderID,
			ReceiverId: msg.ReceiverID,
			Content:    msg.Content,
			Timestamp:  msg.Timestamp,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
