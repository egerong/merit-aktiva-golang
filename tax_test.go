package merit_test

import (
	"encoding/json"
	"testing"
)

func TestGetTaxes(t *testing.T) {
	taxes, err := testClient.GetTaxes()
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(taxes, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
}

func TestGetTaxByCode(t *testing.T) {
	code := "22%"
	tax, err := testClient.GetTaxByCode(code)
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(tax, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
}
