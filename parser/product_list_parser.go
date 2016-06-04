package parser

import "io"

type ProductListParser struct{}

// Parse takes a Reader containing HTML source and streams the results
func (p *ProductListParser) Parse(source io.Reader) (results chan string) {
	url := &searchCriteria{
		Name:          "url",
		RequiredAttrs: map[string]string{"class": "productInfoWrapper"},
		RequiredTag:   "a",
	}

	criteria := []*searchCriteria{url}
	parsed, done := doParse(source, criteria)

	// Subscribe to both channels
	results = make(chan string)

	go func() {
		outer:
		for {
			select {
			case result := <-parsed:
				url := ""
				for _, attr := range result.Tag.Attr {
					if attr.Key == "href" {
						url = attr.Val
						break
					}
				}
				results <- url
			case <-done:
				break outer
			}
		}
		results <- ""
	}()

	return
}
