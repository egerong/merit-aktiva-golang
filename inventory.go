package merit

import "time"

type Location struct {
	CompanyID     int    `json:"CompanyId"`
	LocationID    string `json:"LocationId"`
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	InBPrefix     string `json:"InBPrefix"`
	InBNextNo     int    `json:"InBNextNo"`
	OutBPrefix    string `json:"OutBPrefix"`
	OutBNextNo    int    `json:"OutBNextNo"`
	Loc2LocPrefix string `json:"Loc2LocPrefix"`
	Loc2LocNextNo int    `json:"Loc2LocNextNo"`
	InvSetPrefix  string `json:"InvSetPrefix"`
	InvSetNextNo  int    `json:"InvSetNextNo"`
}

func (c *Client) GetListOfLocations() ([]Location, error) {
	locations := []Location{}
	err := c.post(epGetListOfLocations, map[string]interface{}{}, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

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
	ArticleGroups    []string  `json:"ArticleGroups,omitempty"`
	Location         string    `json:"Location,omitempty"`
	RepDate          time.Time `json:"RepDate,omitempty"`
	ShowZero         bool      `json:"ShowZero,omitempty"`
	WithReservations bool      `json:"WithReservations,omitempty"`
}

type getInventoryReportQueryFormated struct {
	ArticleGroups    []string  `json:"ArticleGroups,omitempty"`
	Location         string    `json:"Location,omitempty"`
	RepDate          QueryDate `json:"RepDate,omitempty"`
	ShowZero         bool      `json:"ShowZero,omitempty"`
	WithReservations bool      `json:"WithReservations,omitempty"`
}

func (c *Client) GetInventoryReport(query GetInventoryReportQuery) ([]Article, error) {
	queryFormated := getInventoryReportQueryFormated{
		ArticleGroups:    query.ArticleGroups,
		Location:         query.Location,
		RepDate:          QueryDate{query.RepDate},
		ShowZero:         query.ShowZero,
		WithReservations: query.WithReservations,
	}
	articles := []Article{}
	err := c.post(epGetInventoryReport, queryFormated, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
