package merit_test

import (
	"encoding/json"
	"testing"
)

func TestGetListOfDimensions(t *testing.T) {
	dimensions, err := testClient.GetDimensions()
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(dimensions, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
}
