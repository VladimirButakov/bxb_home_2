package downloader

import (
	"fmt"
	"io"
	"net/http"
)

type Downloader interface {
	Download(url string) (io.ReadCloser, error)
}

type ImageDownloader struct{}

func (d *ImageDownloader) Download(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading image from %s: %w", url, err)
	}

	return resp.Body, nil
}
