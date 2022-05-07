package shline2dos

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Demo() {
	ch := make(chan *Test)
	go TestData(ch)
	for test := range ch {
		DemoImpl(os.Stdout, test.Feed)
	}
}

func DemoImpl(w io.Writer, feed string) {
	for _, y := range strings.Split(feed, "\n") {
		fmt.Fprintln(w, Quote(y))
	}
}
