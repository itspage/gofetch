package downloader

import "io"

type Downloader interface {
	Download(url string) (*Content, error)
}

type Content struct {
	Data   io.Reader
	Length int64
}
