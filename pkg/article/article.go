package article

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/markphelps/optional"
    "../htmlparse"
    "../util"
    "bytes"
    "strings"
    "fmt"
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

func trimLower(str string) string {
    return strings.TrimSuffix(strings.ToLower(str), ":")
}

func buildMap(parent *goquery.Selection) map[string]optional.String {
    details := make(map[string]optional.String)
    {
        strs := []string{"publisher", "publication date", "isbn-10", "isbn-13", "pages", "format", "size"}
        for i := range strs {
            str := strs[i]
            details[str] = optional.String{}
        }
    }

    keys := parent.Find("span")
    keyStr := make([]string, len(keys.Nodes))

    keys.Each(func(i int, selection *goquery.Selection) {
        keyStr[i] = selection.Text()
    })

    vals := parent.Find("li")
    valStr := make([]string, len(keys.Nodes))
    vals.Each(func(i int, selection *goquery.Selection) {
        valStr[i] = selection.Text()
    })

    for i := range keyStr {
        key := keyStr[i]
        val := valStr[i]
        details[trimLower(key)] = optional.NewString(strings.TrimPrefix(val, key))
    }

    return details
}

func detailMap(doc docType) map[string]optional.String {
    parent := doc.Find("div.book-details > ul")
    return buildMap(parent)
}

func (book BookType) Str() string {
    return fmt.Sprintf("BookType{%s, %s, %s, %s, %s, %s, %s, %s, %s, %s}",
        book.title, book.title, book.desc,
        book.pub.OrElse("nil"),
        book.pubDate.OrElse("nil"),
        book.isbn10.OrElse("nil"),
        book.isbn13.OrElse("nil"),
        book.pages.OrElse("nil"),
        book.format.OrElse("nil"),
        book.size.OrElse("nil"))
}

func Book(htmlBody []byte) BookType {
    doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))
    if err != nil {
        panic(err)
    }

    details := detailMap(doc)

    title := title(doc)
    img := img(doc)
    desc := desc(doc)
    pub := details["publisher"]
    pubDate := details["publication date"]
    isbn10 := details["isbn-10"]
    isbn13 := details["isbn-13"]
    pages := details["pages"]
    format := details["format"]
    size := details["size"]

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
