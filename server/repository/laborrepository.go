package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func FindLaborByID(laborId uuid.UUID) (model.Labor, error) {
	labor := model.Labor{}
	if DB.Where("id = ?", laborId).Find(&labor).Error != nil {
		return labor, fmt.Errorf("Cannot find labor by id: %w", DB.Error)
	}
	return labor, nil
}

func GetAllLabors() ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	if DB.Find(&labors).Error != nil {
		return labors, fmt.Errorf("could not get labors from db: %w", DB.Error)
	}
	return labors, nil
}

func CreateLabor(labor *model.Labor) (uuid.UUID, error) {
	if DB.Create(labor).Error != nil {
		return labor.ID, fmt.Errorf("cannot create labor: %w", DB.Error)
	}
	return labor.ID, nil
}

func UpdateLabor(labor *model.Labor) error {
	if DB.Save(labor).Error != nil {
		return fmt.Errorf("cannot update labor: %w", DB.Error)
	}
	return nil
}

func DeleteLabor(labor *model.Labor) error {
	if DB.Delete(labor).Error != nil {
		return fmt.Errorf("cannot delete labor: %w", DB.Error)
	}
	return nil
}