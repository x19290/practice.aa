package main

import (
	"bytes"
	"strings"
)

func Quote(word string) string {
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
