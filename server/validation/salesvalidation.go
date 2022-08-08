package validation

import (
	"errors"
	"server/model"
	"server/request"
)

func SaleValidation(saleReq *request.SaleRequest) (sale model.Sale, err error) {
	if len(saleReq.Description) < 1 {
		return sale, errors.New("description is required")
	}
	if saleReq.Units < 1 {
		return sale, errors.New("unit(s) is required")
	}
	if saleReq.UnitCost < 0.01 {
		return sale, errors.New("invalid input for unit cost")
	}
	sale = model.Sale{
		ClientId: saleReq.ClientId,
		InvoiceId: saleReq.InvoiceId,
		Description: saleReq.Description,
		Units: int8(saleReq.Units),
		UnitCost: saleReq.UnitCost,
	}
	return sale, err
}