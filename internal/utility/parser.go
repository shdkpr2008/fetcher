package utility

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func NumLinksImagesFromDocument(documentReader io.Reader) (int, int, error) {
	numLinks := 0
	images := 0

	document, err := goquery.NewDocumentFromReader(documentReader)
	if err != nil {
		return numLinks, images, nil
	}

	document.Find("a").Each(func(i int, selection *goquery.Selection) {
		if href, exists := selection.Attr("href"); exists && IsValidURL(href) {
			numLinks += 1
		}
	})

	document.Find("img").Each(func(i int, selection *goquery.Selection) {
		if src, exists := selection.Attr("src"); exists && IsValidURL(src) {
			images += 1
		}
	})

	return numLinks, images, nil
}
