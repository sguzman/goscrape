package main

import (
    "fmt"
    "./brotli"
)

func main() {
    fmt.Println(brotli.Decomp(brotli.Comp([]byte("LOL"))))
}
