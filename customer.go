package merit

import "github.com/shopspring/decimal"

/*
Id	Guid	If filled and customer is found in the database then following fields are not important. If not found, the customer is added using the following fields.
Name	Str 150	Required when customer is added
RegNo	Str 30
NotTDCustomer	Bool	Required when customer is added. EE True for physical persons and foreign companies. PL True for physical persons. Allowed “true” or “false” (lowercase).
VatRegNo	Str 30
CurrencyCode	Str 30
PaymentDeadLine	Int	If missing then taken from default settings.
OverDueCharge	Decimal 5.2	If missing then taken from default settings.
Address	Str 100
City	Str 30
County	Str 100
PostalCode	Str 15
CountryCode	Str 2	Required when adding
PhoneNo	Str 50
PhoneNo2	Str 50
HomePage	Str 80
Email	Str 80
SalesInvLang	Str 8	Invoice language for this specific customer.(ET,EN,RU,FI,PL,SV)
GLNCode	Str 14
PartyCode	Str 20
RefNoBase	Str 36
EInvPaymId	Str 20
EInvOperator	Int	1 - Not exist, 2 - E-invoices to the bank through Omniva, 3 - Bank ( full extent E-invoice), 4- Bank (limited extent E-invoice)
BankAccount	Str 50
*/

type Customer struct {
	Id              string          `json:"Id,omitempty"`   // If filled and customer is found in the database then following fields are not important. If not found, the customer is added using the following fields.
	Name            string          `json:"Name,omitempty"` // Required when customer is added
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
