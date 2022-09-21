package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func (repo *Repository) FindDeletedClients() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	if repo.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&clients).Error != nil {
		return clients, fmt.Errorf("could not get clients from db: %w", repo.db.Error)
	}
	return clients, nil
}

func (repo *Repository) FindDeletedClientByID(clientId uuid.UUID) (model.Client, error) {
	client := model.Client{}
	if repo.db.Unscoped().Where("id = ?", clientId).Find(&client).Error != nil {
		return client, fmt.Errorf("Cannot find client by id: %w", repo.db.Error)
	}
	return client, nil
}

func (repo *Repository) FindClientByID(clientId uuid.UUID) (model.Client, error) {
	client := model.Client{}
	if repo.db.Where("id = ?", clientId).Find(&client).Error != nil {
		return client, fmt.Errorf("Cannot find client by id: %w", repo.db.Error)
	}
	return client, nil
}

func (repo *Repository) GetAllClients() ([]model.Client, error) {
	clients := make([]model.Client, 0)
	if repo.db.Find(&clients).Error != nil {
		return clients, fmt.Errorf("could not get clients from db: %w", repo.db.Error)
	}
	return clients, nil
}

func (repo *Repository) CreateClient(client *model.Client) (uuid.UUID, error) {
	if repo.db.Create(&client).Error != nil {
		return client.ID, fmt.Errorf("cannot create client: %w", repo.db.Error)
	}
	return client.ID, nil
}

func (repo *Repository) UpdateClient(client *model.Client) error {
	if repo.db.Save(&client).Error != nil {
		return fmt.Errorf("cannot update client: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) DeleteClient(client *model.Client) error {
	if repo.db.Delete(&client).Error != nil {
		return fmt.Errorf("cannot delete client: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) UnDeleteClient(client *model.Client) error {
	if repo.db.Model(&client).Unscoped().Update("deleted_at", nil).Error != nil {
		return fmt.Errorf("cannot undelete client: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) PermDeleteClient(client *model.Client) error {
	if repo.db.Unscoped().Delete(&client).Error != nil {
		return fmt.Errorf("cannot permanently delete client: %w", repo.db.Error)
	}
	return nil
}
