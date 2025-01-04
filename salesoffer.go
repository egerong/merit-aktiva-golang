package merit

import (
	"time"

	"github.com/shopspring/decimal"
)

type SalesOffer struct {
	SIHId           string          `json:"SIHId"`
	DepartmentName  string          `json:"DepartmentName"`
	Dimension1Code  string          `json:"Dimension1Code"`
	Dimension2Code  string          `json:"Dimension2Code"`
	Dimension3Code  string          `json:"Dimension3Code"`
	Dimension4Code  string          `json:"Dimension4Code"`
	Dimension5Code  string          `json:"Dimension5Code"`
	Dimension6Code  string          `json:"Dimension6Code"`
	Dimension7Code  string          `json:"Dimension7Code"`
	BatchInfo       string          `json:"BatchInfo"`
	OfferNo         string          `json:"OfferNo"`
	DocumentDate    string          `json:"DocumentDate"`
	TransactionDate string          `json:"TransactionDate"`
	CustomerId      string          `json:"CustomerId"`
	CustomerName    string          `json:"CustomerName"`
	HComment        string          `json:"HComment"`
	FComment        string          `json:"FComment"`
	DueDate         string          `json:"DueDate"`
	CurrencyCode    string          `json:"CurrencyCode"`
	CurrencyRate    decimal.Decimal `json:"CurrencyRate"`
	TaxAmount       decimal.Decimal `json:"TaxAmount"`
	RoundingAmount  decimal.Decimal `json:"RoundingAmount"`
	TotalAmount     decimal.Decimal `json:"TotalAmount"`
	ProfitAmount    decimal.Decimal `json:"ProfitAmount"`
	TotalSum        decimal.Decimal `json:"TotalSum"`
	UserName        string          `json:"UserName"`
	ReferenceNo     string          `json:"ReferenceNo"`
	PriceInclVat    bool            `json:"PriceInclVat"`
	VatRegNo        string          `json:"VatRegNo"`
	PaidAmount      decimal.Decimal `json:"PaidAmount"`
	EInvSent        bool            `json:"EInvSent"`
	EmailSent       string          `json:"EmailSent"`
	DocType         int             `json:"DocType"`
	DocStatus       int             `json:"DocStatus"`
	DeliveryDate    string          `json:"DeliveryDate"`
	Paid            bool            `json:"Paid"`
	ChangedDate     string          `json:"ChangedDate"`
}

type GetSalesOffersQuery struct {
	PeriodStart time.Time `json:"PeriodStart,omitempty"`
	PeriodEnd   time.Time `json:"PeriodEnd,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
}

type getSalesOffersQueryFormated struct {
	PeriodStart queryDate `json:"PeriodStart,omitempty"`
	PeriodEnd   queryDate `json:"PeriodEnd,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
}

func (query GetSalesOffersQuery) format() getSalesOffersQueryFormated {
	return getSalesOffersQueryFormated{
		PeriodStart: queryDate{query.PeriodStart, "20060102"},
		PeriodEnd:   queryDate{query.PeriodEnd, "20060102"},
		DateType:    query.DateType,
		UnPaid:      query.UnPaid,
	}
}

func (c *Client) GetListOfSalesOffers(query GetSalesOffersQuery) ([]SalesOffer, error) {
	queryFormated := query.format()
	salesOffers := []SalesOffer{}
	err := c.post(epGetListOfSalesOffers, queryFormated, &salesOffers)
	if err != nil {
		return nil, err
	}
	return salesOffers, nil
}

type CreateSalesOfferQuery struct {
	Customer       CustomerObject // CustomerObject
	DocDate        time.Time
	ExpireDate     time.Time // if DocType 2 or 3, ExpireDate=DueDate
	DeliveryDate   time.Time
	DocType        int    // 1=quote, 2=sales order, 3=prepayment invoice
	DocStatus      int    // 1=created, 2=sent, 3=approved, 4=rejected, 5=comment received, 6=invoice created, 7=canceled
	OfferNo        string // Recuired
	RefNo          string // Please validate this number yourself.
	CurrencyCode   string
	DepartmentCode string // If used then must be found in the company database.
	ProjectCode    string // If used then must be found in the company database.
	OfferRow       []OfferRow
	TaxAmount      []OfferVAT      // Required
	RoundingAmount decimal.Decimal // Use it for getting PDF invoice to round number. Does not affect TotalAmount.
	TotalAmount    decimal.Decimal // Amount without VAT
	Payment        OfferPayment
	ContactInfo    string   // 4K
	Hcomment       string   // If not specified, API will get it from client record, if it is written there.
	Fcomment       string   // If not specified, API will get it from client record, if it is written there.
	ProcCodes      []string // Poland only. Values: SW, EE, TP, TT_WNT, TT_D, MR_T, MR_UZ, I_42, I_63, B_SPV, B_SPV_DOSTAWA, BMRW_PROWIZJA, MPP
	PolDocType     int      // Poland only. Values: 1-RO, 2-WEW, 3-FP, 4-OJPK
}

type OfferRow struct {
	Item           OfferItem       `json:"Item"` // Sometimes the volume of transactions in the sales software is very high and there is no need to duplicate all the data in accounting. In those cases, you could consider using the same item code for the items with the same VAT rate.
	Quantity       decimal.Decimal `json:"Quantity"`
	Price          decimal.Decimal `json:"Price"`
	DiscountPct    decimal.Decimal `json:"DiscountPct"`
	DiscountAmount decimal.Decimal `json:"DiscountAmount"` // Amount * Price * (DiscountPCt / 100). This is not rounded. Will be substracted from row amount before row roundings.
	TaxId          string          `json:"TaxId"`
	LocationCode   string          `json:"LocationCode"`   // Used for stock items and multiple stocks. If used then must be found in the company database.
	DepartmentCode string          `json:"DepartmentCode"` // If used then must be found in the company database.
	ItemCostAmount decimal.Decimal `json:"ItemCostAmount"` // Required for credit invoices when crediting stock items.
	GLAccountCode  string          `json:"GLAccountCode"`  // If used, must be found in the company database.
	ProjectCode    string          `json:"ProjectCode"`    // If used, must be found in the company database.
	CostCenterCode string          `json:"CostCenterCode"` // If used, must be found in the company database.
}

type OfferItem struct {
	Code            string `json:"Code"`            // Required
	Description     string `json:"Description"`     // Required
	Type            int    `json:"Type"`            // 1 = stock item, 2 = service, 3 = item. Required.
	UOMName         string `json:"UOMName"`         //
	DefLocationCode string `json:"DefLocationCode"` //
	EANCode         string `json:"EANCode"`         //
}

type OfferPayment struct {
	PaymentMethod string          `json:"PaymentMethod"` // Name of the payment method. Must be found in the company database.
	PaidAmount    decimal.Decimal `json:"PaidAmount"`    // Amount with VAT (not more) or less if partial payment
	PaymDate      string          `json:"PaymDate"`      // YYYYmmddHHii
}

type OfferVAT struct {
	TaxId  string          `json:"TaxId"`  // Required. Use gettaxes endpoint to detect the guid needed
	Amount decimal.Decimal `json:"Amount"` // Required
}
