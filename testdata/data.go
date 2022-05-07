package testdata

import (
	"embed"
	"strings"
)

//go:embed *.txt
var emb embed.FS

type Test struct {
	Feed     string
	Expected string
}

func Stream() (ch chan *Test) {
	ch = make(chan *Test)
	go stream(ch)
	return
}

func stream(ch chan *Test) {
	data, err := emb.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(data), "\n----\n")
	for i := 0; i < len(split)-3; i += 3 {
		ch <- &Test{split[i+1], split[i+2]}
	}
	close(ch)
}
