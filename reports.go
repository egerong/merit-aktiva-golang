package merit

import "time"

/*
StartDate	Date
EndDate	Date
ReportType	Int	1-By Invoices, 2-By Vendors, 3-By-Articles, 4-By Fixed assets
VendChoice	Int	1-Vendor and Reporting entry, 2-Vendor, 3-Reporting entry
VendGrpFilter	Str
VendFilter	Str
ItemGrFilter	Str
ItemFilter	Str	Array of items
DepartFilter	Str	Array of departments
FixAssetFilter	Str	Array of Fixed assets
ByEntryNo	Bool	true/false
*/

type GetPurchaseReportQuery struct {
	StartDate      time.Time
	EndDate        time.Time
	VendChoice     int
	VendGrpFilter  string
	VendFilter     string
	ItemGrFilter   string
	ItemFilter     []string
	DepartFilter   []string
	FixAssetFilter []string
	ByEntryNo      bool
}

type reportType int

const (
	_ reportType = iota
	reportTypeByInvoices
	reportTypeByVendors
	reportTypeByArticles
	reportTypeByFixedAssets
)

type getPurchaseReportQueryFormated struct {
	StartDate      QueryDate  `json:"StartDate,omitempty"`
	EndDate        QueryDate  `json:"EndDate,omitempty"`
	ReportType     reportType `json:"ReportType,omitempty"`    // 1-By Invoices, 2-By Vendors, 3-By-Articles, 4-By Fixed assets
	VendChoice     int        `json:"VendChoice,omitempty"`    // 1-Vendor and Reporting entry, 2-Vendor, 3-Reporting entry
	VendGrpFilter  string     `json:"VendGrpFilter,omitempty"` // 1-Vendor and Reporting entry, 2-Vendor, 3-Reporting entry
	VendFilter     string     `json:"VendFilter,omitempty"`
	ItemGrFilter   string     `json:"ItemGrFilter,omitempty"`
	ItemFilter     []string   `json:"ItemFilter,omitempty"`
	DepartFilter   []string   `json:"DepartFilter,omitempty"`
	FixAssetFilter []string   `json:"FixAssetFilter,omitempty"`
	ByEntryNo      bool       `json:"ByEntryNo,omitempty"`
}

type PurchaseReportByInvoice struct {
	DocId          string  `json:"docId"`
	InvoiceNo      string  `json:"invoiceNo"`
	CurrencyCode   string  `json:"currencyCode"`
	CurrencyRate   float64 `json:"currencyRate"`
	VendorId       string  `json:"vendorId"`
	VendorName     string  `json:"vendorName"`
	RegNo          string  `json:"regNo"`
	VatRegNo       string  `json:"vatRegNo"`
	InvoiceDate    string  `json:"invoiceDate"`
	Amount         float64 `json:"amount"`
	VatAmount      float64 `json:"vatAmount"`
	RoundingAmount float64 `json:"roundingAmount"`
	TotalAmount    float64 `json:"totalAmount"`
	ExpenseClaim   bool    `json:"expenseClaim"`
	BatchId        string  `json:"batchId"`
	BatchCode      string  `json:"batchCode"`
	BatchNo        int     `json:"batchNo"`
	Ctry           string  `json:"ctry"`
}

type GetPurchaseReportByInvoiceQuery struct {
	StartDate      time.Time
	EndDate        time.Time
	VendChoice     int
	VendGrpFilter  string
	VendFilter     string
	ItemGrFilter   string
	ItemFilter     []string
	DepartFilter   []string
	FixAssetFilter []string
	ByEntryNo      bool
}

func (c *Client) GetPurchaseReportByInvoice(query GetPurchaseReportQuery) ([]PurchaseReportByInvoice, error) {
	queryFormated := getPurchaseReportQueryFormated{
		StartDate:      QueryDate{query.StartDate},
		EndDate:        QueryDate{query.EndDate},
		ReportType:     reportTypeByInvoices,
		VendChoice:     query.VendChoice,
		VendGrpFilter:  query.VendGrpFilter,
		VendFilter:     query.VendFilter,
		ItemGrFilter:   query.ItemGrFilter,
		ItemFilter:     query.ItemFilter,
		DepartFilter:   query.DepartFilter,
		FixAssetFilter: query.FixAssetFilter,
		ByEntryNo:      query.ByEntryNo,
	}
	reports := []PurchaseReportByInvoice{}
	err := c.post(epGetPurchaseReport, queryFormated, &reports)
	if err != nil {
		return nil, err
	}
	return reports, nil

}

type PurchaseReportByVendor struct {
	CustomerId      string  `json:"customerId"`
	CurrencyCode    string  `json:"currencyCode"`
	VendorName      string  `json:"vendorName"`
	RegNo           string  `json:"regNo"`
	VatRegNo        string  `json:"vatRegNo"`
	Amount          float64 `json:"amount"`
	CAmount         float64 `json:"cAmount"`
	VatAmount       float64 `json:"vatAmount"`
	RoundingAmount  float64 `json:"roundingAmount"`
	TotalAmount     float64 `json:"totalAmount"`
	CVatAmount      float64 `json:"cVatAmount"`
	CRoundingAmount float64 `json:"cRoundingAmount"`
	CTotalAmount    float64 `json:"cTotalAmount"`
	LnCnt           float64 `json:"lnCnt"`
	DiscAmt         float64 `json:"discAmt"`
	CDiscAmt        float64 `json:"cDiscAmt"`
}

type PurchaseReportByArticle struct {
	ItemId       string  `json:"itemId"`
	ItemCode     string  `json:"itemCode"`
	ItemName     string  `json:"itemName"`
	ItemGrId     string  `json:"itemGrId"`
	ItemGrName   string  `json:"itemGrName"`
	CurrencyCode string  `json:"currencyCode"`
	Unit         string  `json:"unit"`
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	UOMId1       string  `json:"uomId1"`
	UOMId2       string  `json:"uomId2"`
}

func (c *Client) GetPurchaseReportByArticle(query GetPurchaseReportQuery) ([]PurchaseReportByArticle, error) {
	queryFormated := getPurchaseReportQueryFormated{
		StartDate:      QueryDate{query.StartDate},
		EndDate:        QueryDate{query.EndDate},
		ReportType:     reportTypeByArticles,
		VendChoice:     query.VendChoice,
		VendGrpFilter:  query.VendGrpFilter,
		VendFilter:     query.VendFilter,
		ItemGrFilter:   query.ItemGrFilter,
		ItemFilter:     query.ItemFilter,
		DepartFilter:   query.DepartFilter,
		FixAssetFilter: query.FixAssetFilter,
		ByEntryNo:      query.ByEntryNo,
	}
	reports := []PurchaseReportByArticle{}
	err := c.post(epGetPurchaseReport, queryFormated, &reports)
	if err != nil {
		return nil, err
	}
	return reports, nil

}

type PurchaseReportByFixedAsset struct {
	DocId        string  `json:"docId"`
	FaId         string  `json:"faId"`
	InventaryNo  string  `json:"inventaryNo"`
	Name         string  `json:"name"`
	CurrencyCode string  `json:"currencyCode"`
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	RemAmount    float64 `json:"remAmount"`
	DocNo        string  `json:"docNo"`
	DocDate      string  `json:"docDate"`
}
