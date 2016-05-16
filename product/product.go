package product

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/itspage/gofetch/downloader"
	"github.com/itspage/gofetch/parser"
)

type Product struct {
	Title        string  `json:"title"`
	Size         string  `json:"size"`
	UnitPrice    float64 `json:"-"`
	UnitPriceStr string  `json:"unit_price"`
	Description  string  `json:"description"`
}

// NewProductFromDownloader takes a downloader and builds a Product struct
func NewProductFromDownloader(dl downloader.Downloader, url string) (*Product, error) {
	// Fetch the URL
	content, err := dl.Download(url)
	if err != nil {
		return nil, err
	}

	// Parse the product
	parser := new(parser.ProductParser)
	results := parser.Parse(content.Data)

	// Convert price to float
	unit_price := strings.Replace(results["unit_price"], "£", "", -1)
	price, err := strconv.ParseFloat(unit_price, 64)
	if err != nil {
		return nil, err
	}
	// Return new Product
	p := &Product{
		Title:        results["title"],
		Size:         fmt.Sprintf("%.2fkb", float64(content.Length) / 1024),
		UnitPrice:    price,
		UnitPriceStr: fmt.Sprintf("%.2f", price),
		Description:  results["description"],
	}
	return p, nil
}
