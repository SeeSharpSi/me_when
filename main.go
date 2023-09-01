package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args[1:]
    lol(args)
}

func lol(me_when []string) {
    for _, v := range me_when {
        fmt.Printf("me when %s\n", v)
    }
}
