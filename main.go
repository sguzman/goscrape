package main

import (
    "fmt"
    "./pkg/red"
    "./pkg/util"
    "./pkg/httpnet"
    "./pkg/htmlparse"
    "os"
    "golang.org/x/net/html"
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

        htmlparse.FlatMap(htmlBody, "h2.post-title > a[href]", func(node *html.Node) {
            href := node.Attr[0].Val
            fmt.Println(href)
        })
    })
}
