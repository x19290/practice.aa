package main

import (
	"embed"
	"strings"
)

//go:embed *.txt
var emb embed.FS

type Test struct {
	feed string
	expected string
}

func testdata(ch chan *Test) {
	data, err := emb.ReadFile("testdata.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(data), "\n----\n")
	for i := 0; i < len(split)-3; i += 3 {
		ch <- &Test{split[i+1], split[i+2]}
	}
	close(ch)
}
