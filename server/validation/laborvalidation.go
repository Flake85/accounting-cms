package validation

import (
	"errors"
	"server/model"
	"server/request"
)

func LaborValidation(laborReq *request.LaborRequest) (labor model.Labor, err error) {
	if len(laborReq.Description) < 1 {
		return labor, errors.New("description is required")
	}
	if laborReq.HoursWorked < 0.01 {
		return labor, errors.New("invalid input for hours worked")
	}
	if laborReq.HourlyRate < 0.01 {
		return labor, errors.New("invalid input for hourly rate")
	}
	labor = model.Labor{
		Description: laborReq.Description,
		ClientID: laborReq.ClientId,
		HoursWorked: laborReq.HoursWorked,
		HourlyRate: laborReq.HourlyRate,
	}
	return labor, err
}
