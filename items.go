package merit

import (
	"fmt"
)

type Item struct {
	ID                   string  `json:"ItemId"`
	Code                 string  `json:"Code"`
	Name                 string  `json:"Name"`
	UnitofMeasureName    string  `json:"UnitofMeasureName"`
	Type                 string  `json:"Type"`
	SalesPrice           float64 `json:"SalesPrice"`
	InventoryQty         float64 `json:"InventoryQty"`
	ReservedQty          float64 `json:"ReservedQty"`
	VatTaxName           string  `json:"VatTaxName"`
	Usage                string  `json:"Usage"`
	SalesAccountCode     string  `json:"SalesAccountCode"`
	PurchaseAccountCode  string  `json:"PurchaseAccountCode"`
	InventoryAccountCode string  `json:"InventoryAccountCode"`
	ItemCostAccountCode  string  `json:"ItemCostAccountCode"`
	DiscountPct          float64 `json:"DiscountPct"`
	LastPurchasePrice    float64 `json:"LastPurchasePrice"`
	ItemUnitCost         float64 `json:"ItemUnitCost"`
	InventoryCost        float64 `json:"InventoryCost"`
	ItemGroupName        string  `json:"ItemGroupName"`
	DefLocName           string  `json:"DefLoc_Name"`
	EANCode              string  `json:"EANCode"`
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
