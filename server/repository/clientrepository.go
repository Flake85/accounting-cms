package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func FindClientByID(clientId uuid.UUID) (model.Client, error) {
	client := model.Client{}
	if DB.Where("id = ?", clientId).Find(&client).Error != nil {
		return client, fmt.Errorf("Cannot find client by id: %w", DB.Error)
	}
	return client, nil
}

func GetAllClients() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	if DB.Find(&clients).Error != nil {
		return clients, fmt.Errorf("could not get clients from db: %w", DB.Error)
	}
	return clients, nil
}

func CreateClient(client *model.Client) (uuid.UUID, error) {
	if DB.Create(client).Error != nil {
		return client.ID, fmt.Errorf("cannot create client: %w", DB.Error)
	}
	return client.ID, nil
}

func UpdateClient(client *model.Client) error {
	if DB.Save(client).Error != nil {
		return fmt.Errorf("cannot update client: %w", DB.Error)
	}
	return nil
}

func DeleteClient(client *model.Client) error {
	if DB.Delete(client).Error != nil {
		return fmt.Errorf("cannot delete client: %w", DB.Error)
	}
	return nil
}