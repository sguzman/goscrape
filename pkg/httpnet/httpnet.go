package httpnet

import (
    "net/http"
    "io/ioutil"
    "../red"
    "../brotli"
)

func Get(url string) []byte {
    htmlResp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer htmlResp.Body.Close()
    body, err := ioutil.ReadAll(htmlResp.Body)
    if err != nil {
        panic(err)
    }

    return body
}

func GetWithCache(url string) string {
    cache := red.Cache
    if val, ok := cache[url]; ok {
        htmlBody := brotli.Decomp(val)
        return htmlBody
    }

    body := Get(url)
    red.InRedis <- red.KeyVal{url, brotli.Comp(body)}
    return string(body)
}