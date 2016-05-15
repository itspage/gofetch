package parser

import (
	"io"

	"golang.org/x/net/html"
)

type searchCriteria struct {
	Name          string
	RequiredAttrs map[string]string
	RequiredTag   string
}

type searchResult struct {
	Name string
	Tag  html.Token
	Text string
}

func doParse(source io.Reader, criteria []*searchCriteria) (results chan *searchResult, done chan struct{}) {
	results = make(chan *searchResult)
	done = make(chan struct{})

	go func() {
		// Iterate through the html and put results back onto the channel
		z := html.NewTokenizer(source)

		// Keep track of the current tag
		var currentTag html.Token
		currentAttrs := [][]html.Attribute{}

		for {
			tokenType := z.Next()

			switch {
			case tokenType == html.ErrorToken:
				done <- struct{}{}
				return
			case tokenType == html.StartTagToken:
				currentTag = z.Token()
				currentAttrs = append(currentAttrs, currentTag.Attr)
			case tokenType == html.EndTagToken:
				currentTag = z.Token()
				currentAttrs = currentAttrs[:len(currentAttrs) - 1]
			case tokenType == html.TextToken:
				// Only process text if we're at the start of a tag
				if currentTag.Type != html.StartTagToken {
					continue
				}
				for _, c := range criteria {
					if matchesSearchCriteria(currentTag, currentAttrs, c) {
						searchResult := &searchResult{
							Name: c.Name,
							Tag:  currentTag,
							Text: string(z.Text()),
						}
						results <- searchResult
					}
				}
				currentTag = z.Token()
			}
		}

	}()
	return
}

func matchesSearchCriteria(currentTag html.Token, currentAttrs [][]html.Attribute, criteria *searchCriteria) bool {
	// Check current tag is correct
	if currentTag.Data != criteria.RequiredTag {
		return false
	}
	// Check all required attrs are present
	for k, v := range criteria.RequiredAttrs {
		if !attrsContains(currentAttrs, k, v) {
			return false
		}
	}
	return true
}

func attrsContains(allAttributes [][]html.Attribute, key, value string) bool {
	for _, attributes := range allAttributes {
		for _, a := range attributes {
			if a.Key == key && a.Val == value {
				return true
			}
		}
	}
	return false
}
