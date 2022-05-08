package dos

import (
	"fmt"
	"github.com/x19290/go.dos/testdata"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Demo() {
	for test := range testdata.LazyList() {
		DemoImpl(os.Stdout, test)
	}
}

func DemoImpl(w io.Writer, test string) {
	for _, feed := range strings.Split(test, "\n") {
		fmt.Fprintf(w, "<%v>\n", feed)
		fmt.Fprintln(w, DosWord(feed))
	}
}

var py = `
from subprocess import list2cmdline
from sys import argv, stdout
b = stdout.buffer
for y in argv[1:]:
	b.write(b"<%s>\n" % y.encode())
	b.write(b"%s\n" % list2cmdline([y]).encode())
`[1:]

func DemoCompare(w io.Writer, test string) {
	feeds := strings.Split(test, "\n")
	feeds = append([]string{"-c", py}, feeds...)
	x := exec.Command("python3", feeds...)
	x.Stdout = w
	err := x.Run()
	if err != nil {
		panic(err)
	}
}
