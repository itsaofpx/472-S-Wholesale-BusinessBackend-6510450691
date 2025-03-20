package database

import (
	"fmt"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type MessagePostgresRepository struct {
	db *gorm.DB
}

func InitiateMessagePostgresRepository(db *gorm.DB) repositories.MessageRepository{
	return &MessagePostgresRepository{
		db: db,
	}
}

func (mpr *MessagePostgresRepository) ValidateReferences(userID, chatID int) error {
	// Check if user exists
	var user entities.User
	if err := mpr.db.First(&user, userID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
					return fmt.Errorf("user with ID %d not found", userID)
			}
			return err
	}

	// Check if chat exists
	var chat entities.Chat
	if err := mpr.db.First(&chat, chatID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
					return fmt.Errorf("chat with ID %d not found", chatID)
			}
			return err
	}

	return nil
}
func (mpr *MessagePostgresRepository) CreateMessage(m entities.Message) (entities.Message, error) {
	if err := mpr.db.Create(&m).Error; err != nil {
			return entities.Message{}, err
	}
	
	var messageWithData entities.Message
	err := mpr.db.
			Preload("User.TierList").
			First(&messageWithData, m.ID).Error
	
	if err != nil {
			return entities.Message{}, err
	}
	
	return messageWithData, nil
}

func (mpr *MessagePostgresRepository) CreateMessageByChatID(m entities.Message) (entities.Message, error) {
	if err := mpr.db.Create(&m).Error; err != nil {
			return entities.Message{}, err
	}
	
	var messageWithData entities.Message
	err := mpr.db.
			Preload("User.TierList").
			First(&messageWithData, m.ID).Error
	
	if err != nil {
			return entities.Message{}, err
	}
	
	return messageWithData, nil
}