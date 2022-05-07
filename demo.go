package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Demo() {
	ch := make(chan *Test)
	go TestData(ch)
	for test := range ch  {
		DemoImpl(os.Stdout, test.feed)
	}
}

func DemoImpl(w io.Writer, feed string) {
	for _, y := range strings.Split(feed, "\n") {
		fmt.Fprintln(w, Quote(y))
	}
}
