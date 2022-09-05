package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func fastLineCounter(r io.Reader) (int64, error) {
	buf := make([]byte, 32*1024)
	var count int64
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += (int64)(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func main() {
	var lines int64
	flag.Int64Var(&lines, "lines", -1, "lines")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	if lines == -1 {
		lines, err = fastLineCounter(f)
		if err != nil {
			panic(err)
		}
	}

	var currentLine int64
	var scanner bufio.Scanner

	offset := rand.Int63n(lines)

	f.Seek(0, 0)
	scanner = *bufio.NewScanner(f)
	currentLine = 0
	for scanner.Scan() {
		currentLine++

		if currentLine < offset {
			continue
		}
		fmt.Println(scanner.Text())
	}

	f.Seek(0, 0)
	scanner = *bufio.NewScanner(f)
	currentLine = 0
	for scanner.Scan() {
		currentLine++
		if currentLine >= offset {
			break
		}
		fmt.Println(scanner.Text())
	}
}
