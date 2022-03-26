package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "mahbub",
		Price: 10,
		SKU:   "abs-abc-def",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
