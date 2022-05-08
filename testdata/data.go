package testdata

import (
	"embed"
	"strings"
)

//go:embed *.txt
var emb embed.FS

func Stream() (ch chan string) {
	ch = make(chan string)
	go func() {
		bits, err := emb.ReadFile("data.txt")
		if err != nil {
			panic(err)
		}
		data := strings.TrimRight(string(bits), "\n")
		split := strings.Split(data, "\n----\n")
		for i := 0; i <= len(split)-2; i += 2 {
			ch <- split[i+1]
		}
		close(ch)
	}()
	return
}
