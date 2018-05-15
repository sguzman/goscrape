package util

import (
    "strings"
    "strconv"
)

const (
    Limit = IntType(1290)
    Cores = IntType(8)
)

type IntType uint64

func Page(page IntType) string {
    const base = "http://23.95.221.108/page/"
    return strings.Join([]string{base, strconv.FormatUint(uint64(page), 10)}, "")
}

func StripHost(url string) string {
    return strings.TrimPrefix(url, "https://it-eb.com")
}

func Path(path string) string {
    const base = "http://23.95.221.108"
    return strings.Join([]string{base, path}, "")
}

func PFor(f func(IntType)) {
    comm := make(chan IntType)
    newLimit := Limit + 1

    for i := IntType(0); i < Cores; i += 1 {
        go func() {
            for {
                getIdx := <- comm
                if getIdx == 0 {
                    return
                }

                f(getIdx)
            }
        }()
    }

    for i := IntType(1); i < newLimit; i += 1 {
        comm <- i
    }

    for i := IntType(0); i < Cores; i += 1 {
        comm <- 0
    }
}
