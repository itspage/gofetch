package product

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/itspage/gofetch/downloader"
	"github.com/itspage/gofetch/parser"
)

type Product struct {
	Title       string
	Size        string
	UnitPrice   float64
	Description string
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
		Title:       results["title"],
		Size:        fmt.Sprintf("%.2fkb", float64(content.Length)/1024),
		UnitPrice:   price,
		Description: results["description"],
	}
	return p, nil
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"title":       p.Title,
		"size":        p.Size,
		"unit_price":  fmt.Sprintf("%.2f", p.UnitPrice),
		"description": p.Description,
	})
}
