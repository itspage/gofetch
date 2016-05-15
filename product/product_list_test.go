package product

import (
	"os"
	"testing"

	"github.com/itspage/gofetch/downloader"
)

type TestDownloader struct{}

func (t *TestDownloader) Download(url string) (*downloader.Content, error) {
	if url == "LIST_URL" {
		r, _ := os.Open("../test_data/product_list_page.html")
		return &downloader.Content{r, 999}, nil
	} else {
		r, _ := os.Open("../test_data/product_page.html")
		return &downloader.Content{r, 123}, nil
	}
}

func TestNewProductListFromDownloader(t *testing.T) {
	td := &TestDownloader{}
	pl, err := NewProductListFromDownloader(td, "LIST_URL")
	if err != nil {
		t.Error(err)
	}

	expectedLen := 7
	if len(pl.Results) != expectedLen {
		t.Errorf("Product results len is %v, want %v", len(pl.Results), expectedLen)
	}

	expected := "12.60"
	if pl.TotalStr != expected {
		t.Errorf("Total is %v, want %v", pl.Total, expected)
	}
}
