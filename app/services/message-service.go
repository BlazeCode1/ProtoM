package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/BlazeCode1/ProtoM/tree/main/ProtoM-GRPC/app/model"
)

// MessageService handles message operations
type MessageService struct {
	messages []model.ChatMessage
	mu       sync.Mutex
}

// NewMessageService creates a new instance of MessageService
func NewMessageService() *MessageService {
	return &MessageService{}
}

// SendMessage processes an incoming message
func (s *MessageService) SendMessage(msg *model.ChatMessage) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg.ID = fmt.Sprintf("%d", len(s.messages)+1)
	msg.Timestamp = time.Now().Unix()

	s.messages = append(s.messages, *msg)
	fmt.Printf("📩 New message from %s: %s\n", msg.SenderID, msg.Content)

	return msg.ID
}

// StreamMessages streams messages for a user
func (s *MessageService) StreamMessages(userID string) []*model.ChatMessage {
	var userMessages []*model.ChatMessage
	for _, msg := range s.messages {
		if msg.ReceiverID == userID {
			userMessages = append(userMessages, &msg)
		}
	}
	return userMessages
}
