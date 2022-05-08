package cmdline

import (
	"bytes"
	"github.com/google/shlex"
	"strings"
)

func FromShlex(shline string) (dosline string) {
	dosline, err := FromShlexEx(shline)
	if err != nil {
		panic(err)
	}
	return
}

func FromShlexEx(shline string) (dosline string, err error) {
	list, err := shlex.Split(shline)
	if err == nil {
		dosline = Make(list...)
	} else {
		dosline = ""
	}
	return
}

func Make(anystrs ...string) string {
	b := new(bytes.Buffer)
	for _, any := range anystrs {
		b.WriteString(SafeWord(any))
		b.WriteByte(' ')
	}
	if 1 <= len(anystrs) {
		b.Truncate(b.Len() - 1)
	}
	return b.String()
}

func SafeWord(any string) string {
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
	// nbs, repeatBs: bs=backslash
	nbs := 0
	repeatBs := func() {
		for i := 0; i < nbs; i++ { // nbs varies
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
	repeatBs() // remaining backslashes (if any)
	if quote {
		repeatBs() // backslashes before closing " must be doubled
		b.WriteByte('"')
	}
	return b.String()
}
