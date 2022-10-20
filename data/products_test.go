package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "testname",
		Price: 123,
		SKU: "abs-abs-abs",
	}

	err := p.Validate()
	
	if err != nil {
		t.Fatal(err)
	}
}