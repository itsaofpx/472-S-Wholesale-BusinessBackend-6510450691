package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type MessageUsecase interface {
	CreateMessage(m entities.Message) (entities.Message, error)
	CreateMessageByChatID(m entities.Message) (entities.Message, error)
}

type MessageService struct {
	repo repositories.MessageRepository
}

func InitiateMessageService(repo repositories.MessageRepository) MessageUsecase {
	return &MessageService{
		repo: repo,
	}
}

func (ms *MessageService) CreateMessage(m entities.Message) (entities.Message, error) {
	// Validate that User and Chat exist
	if err := ms.repo.ValidateReferences(m.UserID, m.ChatID); err != nil {
		return entities.Message{}, err
	}

	message, err := ms.repo.CreateMessage(m)
	if err != nil {
		return entities.Message{}, err
	}
	return message, nil
}

func (ms *MessageService) CreateMessageByChatID(m entities.Message) (entities.Message, error) {
	// Validate that User and Chat exist
	if err := ms.repo.ValidateReferences(m.UserID, m.ChatID); err != nil {
		return entities.Message{}, err
	}

	message, err := ms.repo.CreateMessage(m)
	if err != nil {
		return entities.Message{}, err
	}
	return message, nil
}
