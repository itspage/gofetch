package parser

import "io"

type ProductListParser struct{}

// Parse takes a Reader containing HTML source and returns a slice of Product URLs
func (p *ProductListParser) Parse(source io.Reader) (results []string) {
	url := &searchCriteria{
		Name:          "url",
		RequiredAttrs: map[string]string{"class": "productInfoWrapper"},
		RequiredTag:   "a",
	}

	criteria := []*searchCriteria{url}
	resultsChannel, done := doParse(source, criteria)

	// Subscribe to both channels
	results = []string{}
	outer:
	for {
		select {
		case result := <-resultsChannel:
			url := ""
			for _, attr := range result.Tag.Attr {
				if attr.Key == "href" {
					url = attr.Val
					break
				}
			}
			results = append(results, url)
		case <-done:
			break outer
		}
	}

	return results
}
