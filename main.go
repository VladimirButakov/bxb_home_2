package main

import (
	"fmt"
	"github.com/VladimirButakov/bxb_home_2/internal/processor"
)

func main() {
	urls := []string{
		"https://hsto.org/getpro/habr/upload_files/79c/fd3/274/79cfd32741231fae774d649591801fe4.png",
		"https://hsto.org/r/w1560/files/ce4/897/ad0/ce4897ad07f34b8f8e1ed0270823e5de.jpg",
		"https://hsto.org/files/cc3/592/a6e/cc3592a6e53c436f9501e1c9d4d66f76.jpg",
	}

	imageProcessor := processor.NewImageProcessor()

	savePath := "./img"

	results, err := imageProcessor.ProcessImages(urls, 10, savePath)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	for url, hash := range results {
		fmt.Printf("Image URL: %s, Hash: %s\n", url, hash)
	}
}
