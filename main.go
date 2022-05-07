package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ch := make(chan *Test)
	go testdata(ch)
	for test := range ch  {
		_main(os.Stdout, test.feed)
	}
}

func _main(w io.Writer, feed string) {
	for _, y := range strings.Split(feed, "\n") {
		fmt.Fprintln(w, Quote(y))
	}
}
