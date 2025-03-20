package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type ChatPostgresRepository struct {
	db *gorm.DB
}

func InitiateChatPostgresRepository(db *gorm.DB) repositories.ChatRepository {
	return &ChatPostgresRepository{
		db: db,
	}
}

func (cpr *ChatPostgresRepository) CreateChat(c entities.Chat) (entities.Chat, error) {
	if err := cpr.db.Create(&c).Error; err != nil {
			return entities.Chat{}, err
	}
	
	var chatWithData entities.Chat
	err := cpr.db.
			Preload("Messages", func(db *gorm.DB) *gorm.DB {
					return db.Order("messages.created_at ASC")
			}).
			Preload("Messages.User.TierList").
			First(&chatWithData, c.ID).Error
	
	if err != nil {
			return entities.Chat{}, err
	}
	
	return chatWithData, nil
}


func (cpr *ChatPostgresRepository) GetAllChats() ([]entities.Chat, error) {
	var chats []entities.Chat
	
	err := cpr.db.
			Preload("Messages", func(db *gorm.DB) *gorm.DB {
					return db.Order("messages.created_at ASC")
			}).
			Preload("Messages.User.TierList").
			Find(&chats).Error
	
	if err != nil {
			return nil, err
	}
	
	return chats, nil
}

func (cpr *ChatPostgresRepository) GetChatByUserID(id string) (entities.Chat, error) {
	var chat entities.Chat

	err := cpr.db.
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("messages.created_at ASC")
		}).
		Preload("Messages.User.TierList").
		First(&chat, "User_ID = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Chat{}, err
		}
		return entities.Chat{}, err
	}

	return chat, nil
}