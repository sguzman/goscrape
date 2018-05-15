package htmlparse

import (
    "golang.org/x/net/html"
    "github.com/andybalholm/cascadia"
    "bytes"
)

func FlatMap(body []byte, css string, f func(*html.Node)) {
    art, err := cascadia.Compile(css)
    if err != nil {
        panic(err)
    }

    tree, err := html.Parse(bytes.NewReader(body))
    if err != nil {
        panic(err)
    }

    nodes := art.MatchAll(tree)
    for i := range nodes {
        f(nodes[i])
    }
}