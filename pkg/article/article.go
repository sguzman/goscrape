package article

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/markphelps/optional"
    "../htmlparse"
    "../util"
    "bytes"
)

type BookType struct {
    title, img, desc string
    pub, pubDate, isbn10, isbn13, pages, format, size optional.String
}

type docType = *goquery.Document

func Link(htmlBody []byte, f func(string)) {
    htmlparse.FlatMap(htmlBody, "h2.post-title > a[href]", func(node *goquery.Selection) {
        href, err := node.Attr("href")
        if !err {
            panic(err)
        }

        path := util.StripHost(href)
        url := util.Path(path)
        f(url)
    })
}

func title(doc docType) string {
    return doc.Find("h1.post-title").Text()
}

func img(doc docType) string {
    src, err := doc.Find("div.book-cover > img[src]").Attr("src")
    if !err {
        panic(err)
    }

    return src
}

func desc(doc docType) string {
    return doc.Find("div.entry-inner").Text()
}

func Book(htmlBody []byte) BookType {
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))
    if err != nil {
        panic(err)
    }

    title := title(doc)
    img := img(doc)
    desc := desc(doc)
    pub := optional.String{}
    pubDate := optional.String{}
    isbn10 := optional.String{}
    isbn13 := optional.String{}
    pages := optional.String{}
    format := optional.String{}
    size := optional.String{}

    return BookType{
    title,
    img,
    desc,
    pub,
    pubDate,
    isbn10,
    isbn13,
    pages,
    format,
    size,
    }
}
