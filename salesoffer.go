package merit

import (
	"time"
)

type SalesOffer struct {
	SIHId           string  `json:"SIHId"`
	DepartmentName  string  `json:"DepartmentName"`
	Dimension1Code  string  `json:"Dimension1Code"`
	Dimension2Code  string  `json:"Dimension2Code"`
	Dimension3Code  string  `json:"Dimension3Code"`
	Dimension4Code  string  `json:"Dimension4Code"`
	Dimension5Code  string  `json:"Dimension5Code"`
	Dimension6Code  string  `json:"Dimension6Code"`
	Dimension7Code  string  `json:"Dimension7Code"`
	BatchInfo       string  `json:"BatchInfo"`
	OfferNo         string  `json:"OfferNo"`
	DocumentDate    string  `json:"DocumentDate"`
	TransactionDate string  `json:"TransactionDate"`
	CustomerId      string  `json:"CustomerId"`
	CustomerName    string  `json:"CustomerName"`
	HComment        string  `json:"HComment"`
	FComment        string  `json:"FComment"`
	DueDate         string  `json:"DueDate"`
	CurrencyCode    string  `json:"CurrencyCode"`
	CurrencyRate    float64 `json:"CurrencyRate"`
	TaxAmount       float64 `json:"TaxAmount"`
	RoundingAmount  float64 `json:"RoundingAmount"`
	TotalAmount     float64 `json:"TotalAmount"`
	ProfitAmount    float64 `json:"ProfitAmount"`
	TotalSum        float64 `json:"TotalSum"`
	UserName        string  `json:"UserName"`
	ReferenceNo     string  `json:"ReferenceNo"`
	PriceInclVat    bool    `json:"PriceInclVat"`
	VatRegNo        string  `json:"VatRegNo"`
	PaidAmount      float64 `json:"PaidAmount"`
	EInvSent        bool    `json:"EInvSent"`
	EmailSent       string  `json:"EmailSent"`
	DocType         int     `json:"DocType"`
	DocStatus       int     `json:"DocStatus"`
	DeliveryDate    string  `json:"DeliveryDate"`
	Paid            bool    `json:"Paid"`
	ChangedDate     string  `json:"ChangedDate"`
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

func (c *Client) GetListOfSalesOffers(query GetSalesOffersQuery) ([]SalesOffer, error) {
	formatQuery := getSalesOffersQueryFormated{
		PeriodStart: queryDate{query.PeriodStart},
		PeriodEnd:   queryDate{query.PeriodEnd},
		DateType:    query.DateType,
		UnPaid:      query.UnPaid,
	}
	salesOffers := []SalesOffer{}
	err := c.post(epGetListOfSalesOffers, formatQuery, &salesOffers)
	if err != nil {
		return nil, err
	}
	return salesOffers, nil
}
