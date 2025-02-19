package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type ChatUsecase interface {
	CreateChat(c entities.Chat) (entities.Chat, error)
	GetAllChats() ([]entities.Chat, error)
	GetChatByUserID(id string) (entities.Chat, error)
}

type ChatService struct {
	repo repositories.ChatRepository
}

func InitiateChatService(repo repositories.ChatRepository) ChatUsecase {
	return &ChatService{
		repo: repo,
	}
}

func (cs *ChatService) CreateChat(c entities.Chat) (entities.Chat, error) {
	chat, err := cs.repo.CreateChat(c)
	if err != nil {
		return entities.Chat{}, err
	}
	return chat, nil
}

func (cs *ChatService) GetAllChats() ([]entities.Chat, error) {
	chats, err := cs.repo.GetAllChats()
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (cs *ChatService) GetChatByUserID(id string) (entities.Chat, error) {
	return cs.repo.GetChatByUserID(id)
}