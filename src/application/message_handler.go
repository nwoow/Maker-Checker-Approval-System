package application

import (
	"maker_checker_sqlite/src/domain/entities"
	"maker_checker_sqlite/src/domain/repositories"
	"github.com/google/uuid"
)

// MessageHandler handles application logic for messages
type MessageHandler struct {
	MessageRepo repositories.MessageRepository
}

// SendMessage creates a new message and saves it
func (h *MessageHandler) SendMessage(content, recipient string, maxApprovers int) (string, error) {
	message := &entities.Message{
		ID:           uuid.New().String(),
		Content:      content,
		Recipient:    recipient,
		Status:       "Pending",
		MaxApprovers: maxApprovers,
	}

	if err := h.MessageRepo.Save(message); err != nil {
		return "", err
	}
	return message.ID, nil
}

// ApproveMessage approves a pending message
func (h *MessageHandler) ApproveMessage(messageID, approver string) error {
	message, err := h.MessageRepo.FindByID(messageID)
	if err != nil {
		return err
	}

	if err := message.Approve(approver); err != nil {
		return err
	}

	return h.MessageRepo.Save(message)
}
