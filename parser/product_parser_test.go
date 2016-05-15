package parser

import (
	"os"
	"testing"
)

func TestProductParser(t *testing.T) {
	// Load test file
	source, err := os.Open("../test_data/product_page.html")
	if err != nil {
		t.Error("Couldn't read test source data")
	}

	// Run it through the parser
	parser := new(ProductParser)
	results := parser.Parse(source)

	// Check product title
	expected := "Sainsbury's Golden Kiwi x4"
	if results["title"] != expected {
		t.Errorf("Title is %v, want %v", results["title"], expected)
	}

	// Check product description
	expected = "Gold Kiwi"
	if results["description"] != expected {
		t.Errorf("Description is %v, want %v", results["description"], expected)
	}

	// Check unit price
	expected = "£1.80"
	if results["unit_price"] != expected {
		t.Errorf("Unit Price is %v, want %v", results["unit_price"], expected)
	}
}
