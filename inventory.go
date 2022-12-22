package merit

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
