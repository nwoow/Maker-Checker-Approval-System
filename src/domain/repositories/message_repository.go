package repositories

import "maker_checker_sqlite/src/domain/entities"

// MessageRepository defines the repository interface for messages
type MessageRepository interface {
	Save(message *entities.Message) error
	FindByID(id string) (*entities.Message, error)
}
