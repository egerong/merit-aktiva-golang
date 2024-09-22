package merit

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Item struct {
	ID                   string          `json:"ItemId"`
	Code                 string          `json:"Code"`
	Name                 string          `json:"Name"`
	UnitofMeasureName    string          `json:"UnitofMeasureName"`
	Type                 string          `json:"Type"`
	SalesPrice           decimal.Decimal `json:"SalesPrice"`
	InventoryQty         decimal.Decimal `json:"InventoryQty"`
	ReservedQty          decimal.Decimal `json:"ReservedQty"`
	VatTaxName           string          `json:"VatTaxName"`
	Usage                string          `json:"Usage"`
	SalesAccountCode     string          `json:"SalesAccountCode"`
	PurchaseAccountCode  string          `json:"PurchaseAccountCode"`
	InventoryAccountCode string          `json:"InventoryAccountCode"`
	ItemCostAccountCode  string          `json:"ItemCostAccountCode"`
	DiscountPct          decimal.Decimal `json:"DiscountPct"`
	LastPurchasePrice    decimal.Decimal `json:"LastPurchasePrice"`
	ItemUnitCost         decimal.Decimal `json:"ItemUnitCost"`
	InventoryCost        decimal.Decimal `json:"InventoryCost"`
	ItemGroupName        string          `json:"ItemGroupName"`
	DefLocName           string          `json:"DefLoc_Name"`
	EANCode              string          `json:"EANCode"`
}

type GetItemsQuery struct {
	ID           string `json:"Id,omitempty"`
	Code         string `json:"Code,omitempty"`
	Description  string `json:"Description,omitempty"`
	LocationCode string `json:"LocationCode,omitempty"`
}

func (c *Client) GetItems(query GetItemsQuery) ([]Item, error) {
	items := []Item{}
	err := c.post(epGetItems, query, &items)
	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return items, nil
}

type ItemGroup struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

func (c *Client) GetItemGroups() ([]ItemGroup, error) {
	itemGroups := []ItemGroup{}
	err := c.post(epGetItemGroups, map[string]interface{}{}, &itemGroups)
	if err != nil {
		return nil, err
	}

	fmt.Println(itemGroups)
	return itemGroups, nil
}
