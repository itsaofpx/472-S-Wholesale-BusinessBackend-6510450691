package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type MessageRepository interface {
	CreateMessage(m entities.Message) (entities.Message, error)
	ValidateReferences(userID, chatID int) error 
	CreateMessageByChatID(m entities.Message) (entities.Message, error)
}