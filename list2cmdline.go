package dos

import (
	"bytes"
	"strings"

	"github.com/google/shlex"
)

func ShlexToCmdline(shline string) (dosline string) {
	dosline, err := ShlexToCmdlineEx(shline)
	if err != nil {
		panic(err)
	}
	return
}

func ShlexToCmdlineEx(shline string) (dosline string, err error) {
	list, err := shlex.Split(shline)
	if err == nil {
		dosline = Cmdline(list...)
	} else {
		dosline = ""
	}
	return
}

func Cmdline(anystrs ...string) string {
	b := new(bytes.Buffer)
	for _, any := range anystrs {
		b.WriteString(DosWord(any))
		b.WriteByte(' ')
	}
	if 1 <= len(anystrs) {
		b.Truncate(b.Len() - 1)
	}
	return b.String()
}

func DosWord(any string) string {
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
	repeatBs() // \\+ (?="? $ (if any)
	if quote {
		repeatBs() // \\+ before closing " must be doubled
		b.WriteByte('"')
	}
	return b.String()
}
