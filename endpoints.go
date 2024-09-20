package merit

import "log"

const apiPath = "api/"

type apiEndpoint int

const (
	_ apiEndpoint = iota

	// Sales invoices
	// epGetListOfSalesInvoices apiEndpoint = iota
	// epGetSalesInvoiceDetails
	// epDeleteInvoice
	// epCreateSalesInvoice

	// Sales offers
	epGetListOfSalesOffers
	// epCreateSalesOffer
	// epSetOfferStatus
	// epCreateInvoiceFromOffer
	// epGetSalesOfferDetails
	// epUpdateSalesOffer

	// Get price
	// epGetPrice

	// Recurring Invoices
	// epCreateRecurringInvoice
	// epSendIndicationValues
	// epGetReccurringInvoicesClientAddressList
	// epGetReccurringInvoicesList
	// epGetReccurringInvoiceDetails

	// Purchase invoices
	// epGetListOfPurchaseInvoices
	// epGetPurchaseInvoiceDetails
	// epDeletePurchaseInvoice
	// epCreatePurchaseInvoice

	// Inventory movements
	epGetListOfLocations
	// epCreateInventoryMovement

	// Payments
	// epGetListOfPayments
	// epGetListOfPaymentTypes
	// epCreatePaymentOfSalesInvoice
	// epCreatePaymentOfPurchaseInvoice
	// epCreatePaymentOfSalesOffer
	// epDeletePayment

	// General ledger transactions
	// epCreateGeneralLedgerTransaction
	// epGetListOfGeneralLedgerTransactions
	// epGetGeneralLedgerTransactionDetails
	// epGetGeneralLedgerTransactionFullDetails

	// Tax list
	// epGetTaxList

	// Send tax
	// epSendTax

	// Customers
	// epGetListOfCustomers
	// epCreateCustomer
	// epUpdateCustomer

	// Vendors
	// epGetListOfVendors
	// epCreateVendor
	// epUpdateVendor

	// Accounts list
	// epGetAccountsList

	// Project List
	// epGetProjectList

	// Cost centers list
	// epGetCostCentersList

	// Dimensions
	// epGetDimensionsList
	// epCreateDimensions
	// epAddDimensionValues

	// Departments List
	// epGetDepartmentsList

	// Unit of measure list
	// epGetUnitOfMeasureList

	// Banks list
	// epGetBanksList

	// Financial years
	// epGetFinancialYearsList

	// Items
	epGetItems
	epGetItemGroups
	epAddItems
	epAddItemGroups
	epUpdateItem

	// Reports
	// epCustomerDebtsReport
	// epCustomerPaymentReport
	// epStatementOfProfitOrLoss
	// epStatementOfFinancialPosition
	epInventoryReport
	// epSalesReport
	// epPurchaseReport
)

var endpointMap = map[apiEndpoint]string{
	// Sales invoices
	// Sales offers
	epGetListOfSalesOffers: "v2/getoffers",
	// Get price
	// Recurring Invoices
	// Purchase invoices
	// Inventory movements
	epGetListOfLocations: "v2/getlocations",
	// Payments
	// General ledger transactions
	// Tax list
	// Send tax
	// Customers
	// Vendors
	// Accounts list
	// Project List
	// Cost centers list
	// Dimensions
	// Departments List
	// Unit of measure list
	// Banks list
	// Financial years
	// Items
	epGetItems:      "v1/getitems",
	epGetItemGroups: "v2/getitemgroups",
	// Reports
	epInventoryReport: "v2/getinventoryreport",
}

func (e apiEndpoint) String() string {
	epStr, ok := endpointMap[e]
	if !ok {
		log.Fatalf("Unknown endpoint: %d", e)
	}

	return apiPath + epStr
}
