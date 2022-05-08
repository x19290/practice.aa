package dos

import (
	"fmt"
	"github.com/x19290/practice.aa/testdata"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Demo() {
	// ch := make(chan *testdata.Test)
	// go testdata.Stream(ch)
	for test := range testdata.Stream() {
		DemoImpl(os.Stdout, test)
	}
}

func DemoImpl(w io.Writer, test string) {
	for _, feed := range strings.Split(test, "\n") {
		fmt.Fprintf(w, "<%v>\n", feed)
		fmt.Fprintln(w, Quote(feed))
	}
}

var py = `
from subprocess import list2cmdline
from sys import argv
print(list2cmdline([argv[1]]), end=r"")
`[1:]

func DemoCompare(w io.Writer, test string) {
	for _, feed := range strings.Split(test, "\n") {
		quoted, err := exec.Command("python3", "-c", py, feed).Output()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "<%v>\n", feed)
		fmt.Fprintf(w, "%s\n", quoted)
	}
}