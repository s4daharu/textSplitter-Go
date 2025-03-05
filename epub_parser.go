package main

import (
	epub "github.com/go-shiori/go-epub"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func extractChapters(epubPath string) ([]string, error) {
	book, err := epub.Open(epubPath)
	if err != nil {
		return nil, err
	}

	var chapters []string
	for _, item := range book.Items {
		if item.IsPage() || strings.HasSuffix(item.FileName, ".xhtml") {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Content))
			if err != nil {
				return nil, err
			}
			text := doc.Text()
			chapters = append(chapters, text)
		}
	}
	return chapters, nil
}