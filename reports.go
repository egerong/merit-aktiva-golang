package merit

type Article struct {
	ItemCode         string  `json:"ItemCode"`
	EANCode          string  `json:"EANCode"`
	ItemName         string  `json:"ItemName"`
	LocName          string  `json:"LocName"`
	Quantity         float64 `json:"Quantity"`
	ReservedQuantity float64 `json:"ReservedQuantity"`
	UnitCode         string  `json:"UnitCode"`
	Amount           float64 `json:"Amount"`
	Price            float64 `json:"Price"`
}

type GetInventoryReportQuery struct {
	ArticleGroups    []string `json:"ArticleGroups,omitempty"`
	Location         string   `json:"Location,omitempty"`
	RepDate          string   `json:"RepDate,omitempty"`
	ShowZero         bool     `json:"ShowZero,omitempty"`
	WithReservations bool     `json:"WithReservations,omitempty"`
}

func (c *Client) GetInventoryReport(query GetInventoryReportQuery) ([]Article, error) {
	inventoryReport := []Article{}
	err := c.post(epInventoryReport, query, &inventoryReport)
	if err != nil {
		return nil, err
	}
	return inventoryReport, nil
}
