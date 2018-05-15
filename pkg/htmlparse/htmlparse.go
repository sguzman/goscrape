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