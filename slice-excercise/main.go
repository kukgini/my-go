package main

import (
    "log"
)

func main() {
    log.Println("hello")
    s := make_slice(5)
    s[0] = "a"
    //s[1] = "b"
    s[2] = "c"
    //s[3] = "d"
    s[4] = "e"
    log.Printf("new slice received: %s", s)
    pass_slice_as_args(s...)
    pass_slice_as_slice(s)
}

func make_slice(size int) []string {
    return make([]string, size)
}

func pass_slice_as_args(args...string) {
    log.Printf("args=%T", args)
    for i, x := range(args) {
        log.Printf("element %d=%s", i, x)
    }
}

func pass_slice_as_slice(args []string) {
    log.Printf("args=%T", args)
    for i, x := range(args) {
        log.Printf("element %d=%s", i, x)
    }
}
