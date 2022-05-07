package dos

import (
	"fmt"
	"github.com/x19290/practice.aa/testdata"
	"io"
	"os"
	"strings"
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
		fmt.Fprintf(w, "<%v>\n", y)
		fmt.Fprintln(w, Quote(y))
	}
}
