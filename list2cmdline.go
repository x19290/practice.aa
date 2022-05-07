package dos

import (
	"bytes"
	"github.com/google/shlex"
	"strings"
)

func ShlineToDos(shline string) (dosline string) {
	dosline, err := ShlineToDosLine(shline)
	if err != nil {
		panic(err)
	}
	return
}

func ShlineToDosLine(shline string) (dosline string, err error) {
	list, err := shlex.Split(shline)
	if err == nil {
		dosline = List2Cmdline(list)
	} else {
		dosline = ""
	}
	return
}

func List2Cmdline(anystrs []string) string {
	b := new(bytes.Buffer)
	for _, word := range anystrs {
		b.WriteString(Quote(word))
		b.WriteByte(' ')
	}
	if 1 <= len(anystrs) {
		b.Truncate(b.Len() - 1)
	}
	return b.String()
}

func Quote(any string) (dosword string) {
	// This algorithm is stolen from python3...list2cmdline().
	quote := func() bool {
		switch {
		case len(any) == 0:
			return true
		case 0 <= strings.Index(any, " "), 0 <= strings.Index(any, "\t"):
			return true
		default:
			return false
		}
	}()
	b := new(bytes.Buffer) // return b.String() at last.
	// n'bs', repeat'Bs': bs=backslash
	nbs := 0
	repeatBs := func() {
		for i := 0; i < nbs; i++ {
			b.WriteByte('\\')
		}
	}
	if quote {
		b.WriteByte('"')
	}
	for _, y := range any {
		if y == '\\' {
			nbs++
		} else {
			if y == '"' {
				nbs = nbs*2 + 1
			}
			repeatBs()
			b.WriteRune(y)
			nbs = 0
		}
	}
	repeatBs()
	if quote {
		repeatBs()
		b.WriteByte('"')
	}
	dosword = b.String()
	return
}
