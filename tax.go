package merit

import (
	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/shopspring/decimal"
)

type TaxObject struct {
	TaxID  string          `json:"TaxId"`  // Required. Use gettaxes endpoint to detect the guid needed
	Amount decimal.Decimal `json:"Amount"` // Required
}

// Field Name	Type	Comment
// Id	Guid
// Code	Str 16
// Name	Str 40
// NameEN	Str 40
// NameRU	Str 40
// TaxPct	Decimal 2.2

type Tax struct {
	ID     guid.GUID       `json:"Id"`
	Code   string          `json:"Code"`
	Name   string          `json:"Name"`
	NameEN string          `json:"NameEN"`
	NameRU string          `json:"NameRU"`
	TaxPct decimal.Decimal `json:"TaxPct"`
}

func (c *Client) GetTaxes() ([]Tax, error) {
	taxes := []Tax{}
	err := c.post(epGetListOfTaxes, struct{}{}, &taxes)
	if err != nil {
		return nil, err
	}
	return taxes, nil
}

func (c *Client) GetTaxByCode(code string) (*Tax, error) {
	taxes, err := c.GetTaxes()
	if err != nil {
		return nil, err
	}
	for _, tax := range taxes {
		if tax.Code == code {
			return &tax, nil
		}
	}
	return nil, nil
}
