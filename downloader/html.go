package downloader

import "net/http"

type HTMLDownloader struct{}

func (d *HTMLDownloader) Download(url string) (*Content, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return &Content{resp.Body, resp.ContentLength}, nil
}
