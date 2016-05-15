package parser

import (
	"io"
	"strings"
)

type ProductParser struct{}


// Parse takes a Reader containing HTML source and returns a map of product details
func (p *ProductParser) Parse(source io.Reader) (results map[string]string) {
	title := &searchCriteria{
		Name:        "title",
		RequiredTag: "h1",
	}
	description := &searchCriteria{
		Name:          "description",
		RequiredAttrs: map[string]string{"class": "productText"},
		RequiredTag:   "p",
	}
	price := &searchCriteria{
		Name:          "unit_price",
		RequiredAttrs: map[string]string{"class": "pricePerUnit"},
		RequiredTag:   "p",
	}

	criteria := []*searchCriteria{title, description, price}
	resultsChannel, done := doParse(source, criteria)

	// Subscribe to both channels
	results = map[string]string{}
	outer:
	for {
		select {
		case result := <-resultsChannel:
			results[result.Name] = strings.TrimSpace(result.Text)
			if result.Name == "description" {
				// Stop as soon as we find the first description
				break outer
			}
		case <-done:
			break outer
		}
	}

	return results
}
