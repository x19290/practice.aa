package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
)

func Test0(t *testing.T) {
	ch := make(chan *Test)
	go testdata(ch)
	for test := range ch {
		expected := new(bytes.Buffer)
		actual := new(bytes.Buffer)
		expected.WriteString(test.expected)
		expected.WriteByte('\n')
		_main(actual, test.feed)
		os.WriteFile("e", expected.Bytes(), os.ModePerm)
		os.WriteFile("a", actual.Bytes(), os.ModePerm)
		assert.Equal(t, expected, actual)
	}
}
