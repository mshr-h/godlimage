package godlimage

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/mshr-h/godl"
)

func getPage(url string) []string {
	results := []string{}
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if exists {
			results = append(results, url)
		}
	})
	return results
}

func getImageUrls(url string) []string {
	results := getPage(url)
	filterd_urls := []string{}
	for _, result := range results {
		if isImageUrl(result) {
			filterd_urls = append(filterd_urls, result)
		}
	}
	return results
}

func isImageUrl(url string) bool {
	r := regexp.MustCompile(`.*\.(gif|png|jpg|jpeg|bmp)`)
	return r.MatchString(url)
}

func DownloadImages(url string) {
	results := getImageUrls(url)

	parallels := 0
	done := make(chan bool)

	for _, result := range results {
		parallels++
		go godl.DownloadFromURLParallel(result, "", done)
	}

	for i := 0; i < parallels; i++ {
		<-done
		fmt.Printf(" (%d/%d)\n", i+1, parallels)
	}
}
