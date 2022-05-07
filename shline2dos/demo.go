package shline2dos

import (
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/x19290/practice.aa/testdata"
)

func Demo() {
	// ch := make(chan *testdata.Test)
	// go testdata.Stream(ch)
	for test := range testdata.Stream() {
		DemoImpl(os.Stdout, test.Feed)
	}
}

func DemoImpl(w io.Writer, feed string) {
	for _, y := range strings.Split(feed, "\n") {
		fmt.Fprintln(w, Quote(y))
	}
}
