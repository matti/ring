package main

import (
	"flag"
	"fmt"

	"github.com/matti/ring"
)

func main() {
	flag.Parse()

	filePath := flag.Arg(0)
	lines, err := ring.ReadLines(filePath)
	if err != nil {
		panic(err)
	}

	for {
		line, ok := <-lines
		if !ok {
			break
		}
		fmt.Println(line)
	}
}
