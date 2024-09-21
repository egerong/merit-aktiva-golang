package merit

type apiEndpoint string

const (

	// Sales invoices
	// epGetListOfSalesInvoices apiEndpoint = iota
	// epGetSalesInvoiceDetails
	// epDeleteInvoice
	// epCreateSalesInvoice

	// Sales offers
	epGetListOfSalesOffers apiEndpoint = "v2/getoffers"
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
	epGetListOfLocations          apiEndpoint = "v2/getlocations"
	epGetListOfInventoryMovements apiEndpoint = "v2/getinvmovements"
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
	epGetItems      apiEndpoint = "v1/getitems"
	epGetItemGroups apiEndpoint = "v2/getitemgroups"
	// epAddItems
	// epAddItemGroups
	// epUpdateItem

	// Reports
	// epCustomerDebtsReport
	// epCustomerPaymentReport
	// epStatementOfProfitOrLoss
	// epStatementOfFinancialPosition
	epGetInventoryReport apiEndpoint = "v2/getinventoryreport"
	// epSalesReport
	// epPurchaseReport
)
