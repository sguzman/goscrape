package main

import (
    "fmt"
    "./pkg/red"
    "./pkg/util"
    "./pkg/httpnet"
    "github.com/andybalholm/cascadia"
    "golang.org/x/net/html"
    "os"
    "bytes"
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
        if noRedis {
            go red.Set()
            defer red.Client.Close()
        }
    }

    util.PFor(func(i util.IntType) {
        url := util.Page(i)

        htmlBody := get(url)
        art, err := cascadia.Compile("h2.post-title > a[href]")
        if err != nil {
            panic(err)
        }

        tree, err := html.Parse(bytes.NewReader(htmlBody))
        if err != nil {
            panic(err)
        }

        nodes := art.MatchAll(tree)
        for i := range nodes {
            node := nodes[i]
            href := node.Attr[0].Val
            fmt.Println(href)
        }
    })
}
