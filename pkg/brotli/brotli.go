package brotli

import (
    "gopkg.in/kothar/brotli-go.v0/dec"
    "gopkg.in/kothar/brotli-go.v0/enc"
)

func Decomp(str string) []byte {
    input := []byte(str)

    decompressed, err := dec.DecompressBuffer(input, make([]byte, 0))
    if err != nil {
        panic(err)
    }

    return decompressed
}

func Comp(input []byte) []byte {
    compressed, err := enc.CompressBuffer(nil, input, make([]byte, 0))
    if err != nil {
        panic(err)
    }

    return compressed
}