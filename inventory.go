package main

import "fmt"

func (c *Client) GetListOfLocations() ([]Item, error) {
	items := []Item{}
	err := c.post(epGetListOfLocations, map[string]interface{}{}, &items)
	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return items, nil
}
