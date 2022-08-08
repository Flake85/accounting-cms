package validation

import (
	"errors"
	"server/model"
	"server/request"
)

func ClientValidation(clientReq *request.ClientRequest) (client model.Client, err error) {
	if len(clientReq.Name) < 1 {
		return client, errors.New("name is required")
	}
	if len(clientReq.Email) < 1 {
		return client, errors.New("email is required")
	}
	if len(clientReq.Address) < 1 {
		return client, errors.New("address is required")
	}
	client = model.Client{
		Name: clientReq.Name,
		Email: clientReq.Email,
		Address: clientReq.Address,
	}
	return client, err
}