package product

import (
	"encoding/json"
	"fmt"

	"github.com/itspage/gofetch/downloader"
	"github.com/itspage/gofetch/parser"
)

type ProductList struct {
	Results []*Product
	Total   float64
}

// NewProductListFromDownloader takes a URL and will search this URL for any links to a product.
// For each product, the link is followed and the product details are downloaded.
// The total price of all products is summed and return with the list of Product
func NewProductListFromDownloader(dl downloader.Downloader, url string) (*ProductList, error) {
	// Fetch the URL
	content, err := dl.Download(url)
	if err != nil {
		return nil, err
	}

	parser := new(parser.ProductListParser)
	urls := parser.Parse(content.Data)

	// Parse the product list
	pList := new(ProductList)

	products := make(chan *Product)
	errs := make(chan error)

	// Concurrently fetch each product
	for _, url := range urls {
		go func(url string) {
			p, err := NewProductFromDownloader(dl, url)
			if err != nil {
				errs <- err
			}
			products <- p
		}(url)
	}

	// Collect product results
	for i := 0; i < len(urls); i++ {
		select {
		case p := <-products:
			if p != nil {
				pList.Results = append(pList.Results, p)
				pList.Total += p.UnitPrice
			}
		case err := <-errs:
			return nil, err
		}
	}

	return pList, nil
}

func (p ProductList) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"results": p.Results,
		"total":   fmt.Sprintf("%.2f", p.Total),
	})
}
