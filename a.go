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

type Test struct {
	feed string
	expected string
}

func testdata(ch chan *Test) {
	data, err := emb.ReadFile("testdata.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(data), "\n----\n")
	for i := 0; i < len(split)-3; i += 3 {
		ch <- &Test{split[i+1], split[i+2]}
	}
	close(ch)
}

func main() {
	ch := make(chan *Test)
	go testdata(ch)
	for test := range ch  {
		_main(os.Stdout, test)
	}
}

func _main(w io.Writer, test *Test) {
	for _, y := range strings.Split(test.feed, "\n") {
		fmt.Fprintln(w, quotedog(y))
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
