package main

import (
    "fmt"
    "./pkg/red"
    "./pkg/util"
    "./pkg/httpnet"
    "./pkg/htmlparse"
    "./pkg/article"
    "os"
    "github.com/PuerkitoBio/goquery"
)

var (
    noRedis bool
    get func(string) []byte
)

func init() {
    noRedis = os.Getenv("DO_NOT_USE_REDIS") == "true"
    if noRedis {
        fmt.Println("Not using redis")
        get = httpnet.Get
    } else {
        fmt.Println("Using redis")
        get = httpnet.GetWithCache
        red.Init()
    }
}

func main() {
    {
        if !noRedis {
            go red.Set()
            defer red.Client.Close()
        }
    }

    util.PFor(func(i util.IntType) {
        url := util.Page(i)
        htmlBody := get(url)

        article.Link(htmlBody, func(url string) {
            htmlBody := get(url)
            title := htmlparse.Text(htmlBody, "h1.post-title")

            fmt.Println(title)
        })
    })
}
