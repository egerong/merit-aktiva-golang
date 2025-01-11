package merit

type DimensionsObject struct {
	ID         string `json:"Id"`
	DimID      int    `json:"DimId"`
	DimValueId string `json:"DimValueId"`
	DimCode    string `json:"DimCode"`
}

type Dimension struct {
	DimID   int    `json:"DimId"`
	DimName string `json:"DimName"`
	ID      string `json:"Id"`
	Code    string `json:"Code"`
	Name    string `json:"Name"`
	EndDate string `json:"EndDate"`
}

func (c *Client) GetDimensions() ([]Dimension, error) {
	var query struct{}
	dimensions := []Dimension{}
	err := c.post(epGetListOfDimensions, query, &dimensions)
	if err != nil {
		return nil, err
	}
	return dimensions, nil
}
