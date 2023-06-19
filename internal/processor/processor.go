package processor

import (
	"fmt"
	"io"
	"sync"

	"github.com/VladimirButakov/bxb_home_2/internal/downloader"
	"github.com/VladimirButakov/bxb_home_2/internal/hash"
	"github.com/VladimirButakov/bxb_home_2/internal/saver"
)

type ImageProcessor struct {
	Downloader     downloader.Downloader
	HashCalculator hash.Calculator
	ImageSaver     saver.ImageSaver
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{
		Downloader:     &downloader.ImageDownloader{},
		HashCalculator: &hash.MD5HashCalculator{},
		ImageSaver:     &saver.FileImageSaver{},
	}
}

func (p *ImageProcessor) ProcessImages(urls []string, workers int, savePath string) (map[string]string, error) {
	results := make(map[string]string)
	urlChan := make(chan string, len(urls))
	resultChan := make(chan struct {
		url  string
		hash string
	}, len(urls))
	var wg sync.WaitGroup

	closeContent := func(content io.Closer) {
		err := content.Close()
		if err != nil {
			fmt.Printf("Error closing image content: %s\n", err.Error())
		}
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlChan {
				content, err := p.Downloader.Download(url)
				if err != nil {
					fmt.Printf("Error downloading image from %s: %s\n", url, err.Error())
					return
				}
				closeContent(content)

				hash, err := p.HashCalculator.CalculateHash(content)
				if err != nil {
					fmt.Printf("Error processing image from %s: %s\n", url, err.Error())
					return
				}

				resultChan <- struct {
					url  string
					hash string
				}{url, hash}
			}
		}()
	}

	for _, url := range urls {
		urlChan <- url
	}

	close(urlChan)
	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		results[result.url] = result.hash

		content, err := p.Downloader.Download(result.url)
		if err != nil {
			fmt.Printf("Error downloading image from %s: %s\n", result.url, err.Error())
			continue
		}
		closeContent(content)

		err = p.ImageSaver.Save(result.hash, savePath, content)
		if err != nil {
			fmt.Printf("Error saving image from %s: %s\n", result.url, err.Error())
			continue
		}
	}

	return results, nil
}
