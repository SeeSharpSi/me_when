package main

import (
	"fmt"
    "flag"
)

func main() {
    jumble_flag := flag.Bool("j", false, "jubmle the output")
    flag.Parse()
    args := flag.Args()
    lol(args, *jumble_flag)
}

func lol(me_when []string, jmbl bool) {
    for _, v := range me_when {
        if jmbl {
            v = jumble(v)
        }
        fmt.Printf("me when %s\n", v)
    }
}

func jumble(input string) (output string) {
    bytes := []byte(input)
    backwards_bytes := make([]byte, 0)
    for _, v := range bytes {
        backwards_bytes = append([]byte{v}, backwards_bytes...)
    }
    output = string(backwards_bytes)
    return
}
