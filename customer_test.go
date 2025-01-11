package merit_test

import (
	"encoding/json"
	"testing"

	"github.com/egerong/merit-aktiva-go"
)

func TestGetCustomers(t *testing.T) {
	query := merit.GetCustomersQuery{}
	customers, err := testClient.GetCustomers(query)
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
}
