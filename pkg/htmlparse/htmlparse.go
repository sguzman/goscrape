package htmlparse

import (
    "github.com/PuerkitoBio/goquery"
    "bytes"
)

func FlatMap(body []byte, css string, f func(selection *goquery.Selection)) {
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
    if err != nil {
        panic(err)
    }

    doc.Find(css).Each(func(i int, selection *goquery.Selection) {
        f(selection)
    })
}

func node(body []byte, css string) *goquery.Selection {
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
    if err != nil {
        panic(err)
    }

    return doc.Find(css)
}

func Text(body []byte, css string) string {
    return node(body, css).Text()
}