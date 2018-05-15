package article

import (
    "github.com/PuerkitoBio/goquery"
    "../htmlparse"
    "../util"
)

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
