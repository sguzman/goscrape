package main

import (
    "fmt"
    "./pkg/red"
    "./pkg/util"
    "./pkg/httpnet"
    "github.com/andybalholm/cascadia"
    "golang.org/x/net/html"
    "strings"
)

func init() {
    red.Init()
}

func main() {
    {
        go red.Set()
        defer red.Client.Close()
    }

    util.PFor(func(i util.IntType) {
        url := util.Page(i)

        htmlBody := httpnet.GetWithCache(url)
        art, err := cascadia.Compile("h2.post-title > a[href]")
        if err != nil {
            panic(err)
        }

        tree, err := html.Parse(strings.NewReader(htmlBody))
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
