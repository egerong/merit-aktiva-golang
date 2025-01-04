package merit

import (
	// "merit"

	"github.com/shopspring/decimal"
)

type GetCustomersQuery struct {
	ID       string `json:"Id,omitempty"`       // If filled, next fields will be ignored
	RegNo    string `json:"RegNo,omitempty"`    // Exact match.
	VatRegNo string `json:"VatRegNo,omitempty"` // Exact match.
	Name     string `json:"Name,omitempty"`     // Broad match.
}

func (c *Client) GetCustomers(query GetCustomersQuery) ([]CustomerObject, error) {
	customers := []CustomerObject{}
	err := c.post(epGetListOfCustomers, query, &customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

type CustomerObject struct {
	Id              string          `json:"CustomerId,omitempty"` // If filled and customer is found in the database then following fields are not important. If not found, the customer is added using the following fields.
	Name            string          `json:"Name,omitempty"`       // Required when customer is added
	RegNo           string          `json:"RegNo,omitempty"`
	NotTDCustomer   bool            `json:"NotTDCustomer,omitempty"` // Required when customer is added. EE True for physical persons and foreign companies. PL True for physical persons. Allowed “true” or “false” (lowercase).
	VatRegNo        string          `json:"VatRegNo,omitempty"`
	CurrencyCode    string          `json:"CurrencyCode,omitempty"`
	PaymentDeadLine int             `json:"PaymentDeadLine,omitempty"` // If missing then taken from default settings.
	OverDueCharge   decimal.Decimal `json:"OverDueCharge,omitempty"`   // If missing then taken from default settings.
	Address         string          `json:"Address,omitempty"`
	City            string          `json:"City,omitempty"`
	County          string          `json:"County,omitempty"`
	PostalCode      string          `json:"PostalCode,omitempty"`
	CountryCode     string          `json:"CountryCode,omitempty"` // Required when adding
	PhoneNo         string          `json:"PhoneNo,omitempty"`
	PhoneNo2        string          `json:"PhoneNo2,omitempty"`
	HomePage        string          `json:"HomePage,omitempty"`
	Email           string          `json:"Email,omitempty"`
	SalesInvLang    string          `json:"SalesInvLang,omitempty"` // Invoice language for this specific customer.(ET,EN,RU,FI,PL,SV)
	GLNCode         string          `json:"GLNCode,omitempty"`
	PartyCode       string          `json:"PartyCode,omitempty"`
	RefNoBase       string          `json:"RefNoBase,omitempty"`
	EInvPaymId      string          `json:"EInvPaymId,omitempty"`
	EInvOperator    int             `json:"EInvOperator,omitempty"` // 1 - Not exist, 2 - E-invoices to the bank through Omniva, 3 - Bank ( full extent E-invoice), 4- Bank (limited extent E-invoice)
	BankAccount     string          `json:"BankAccount,omitempty"`
}
