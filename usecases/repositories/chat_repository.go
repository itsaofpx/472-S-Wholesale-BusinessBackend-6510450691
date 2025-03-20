package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type ChatRepository interface {
	CreateChat(c entities.Chat) (entities.Chat, error)
	GetAllChats() ([]entities.Chat, error)
	GetChatByUserID(id string) (entities.Chat, error)
}