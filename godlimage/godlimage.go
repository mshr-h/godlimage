package main

import (
	"log"
	"os"

	"github.com/mshr-h/godlimage"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("./get_images url")
	}

	url := os.Args[1]

	godlimage.DownloadImages(url)
}
