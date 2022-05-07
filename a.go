package main

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
)

//go:embed *.txt
var emb embed.FS

func testdata(ch chan string) {
	data, err := emb.ReadFile("testdata.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(data), "\n----\n")
	for i := 0; i < len(split)-3; i += 3 {
		ch <- split[i+1]
		ch <- split[i+2]
	}
	close(ch)
}

func main() {
	_main(os.Stdout)
}

func _main(w io.Writer) {
	ch := make(chan string)
	go testdata(ch)
	for {
		feed, ok := <-ch
		if !ok {
			break
		}
		_, ok = <-ch
		if !ok {
			panic("?")
		}
		for _, y := range strings.Split(feed, "\n") {
			fmt.Fprintln(w, quotedog(y))
		}
	}
}

func quotedog(word string) string {
	j := make(chan rune)
	var needquote bool
	switch {
	case len(word) == 0:
		needquote = true
	case 0 <= strings.Index(word, " "), 0 <= strings.Index(word, "\t"):
		needquote = true
	default:
		needquote = false
	}
	go func() {
		nbs := 0
		if needquote {
			j <- '"'
		}
		for _, y := range word {
			if y == '\'' {
				nbs++
			} else {
				if y == '"' {
					nbs = nbs*2 + 1
				}
				for i := 0; i < nbs; i++ {
					j <- '\\'
				}
				nbs = 0
				j <- y
			}
		}
		for i := 0; i < nbs; i++ {
			j <- '\\'
		}
		if needquote {
			for i := 0; i < nbs; i++ {
				j <- '\\'
			}
			j <- '"'
		}
		close(j)
	}()
	b := new(bytes.Buffer)
	for {
		c, ok := <-j
		if !ok {
			break
		}
		b.WriteRune(c)
	}
	return b.String()
}
