package ring

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"os"
	"time"
)

func fastLineCounter(r io.Reader) (uint64, error) {
	buf := make([]byte, 32*1024)
	var count uint64
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += (uint64)(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ReadLines(path string) (chan string, error) {
	var lines chan string = make(chan string, 1024)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	go func() {
		rand.Seed(time.Now().UnixNano())

		lineCount, err := fastLineCounter(f)
		if err != nil {
			panic(err)
		}
		f.Seek(0, 0)

		offset := rand.Int63n((int64)(lineCount))
		var scanner *bufio.Scanner
		scanner = bufio.NewScanner(f)

		for i := int64(0); scanner.Scan(); i++ {
			// rewind file forward up to offset
			if i < offset {
				continue
			}
			//lines <- strconv.Itoa(int(i)+1) + ": " + scanner.Text()
			lines <- scanner.Text()
		}

		f.Seek(0, 0)
		scanner = bufio.NewScanner(f)
		for i := int64(0); scanner.Scan(); i++ {
			if i >= offset {
				break
			}
			//lines <- strconv.Itoa(int(i)+1) + ": " + scanner.Text()
			lines <- scanner.Text()
		}

		close(lines)
	}()

	return lines, nil
}
