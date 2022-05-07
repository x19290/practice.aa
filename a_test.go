package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	ch := make(chan *Test)
	go testdata(ch)
	for test := range ch {
		expected := new(bytes.Buffer)
		actual := new(bytes.Buffer)
		expected.WriteString(test.expected)
		expected.WriteRune('\n')
		_main(actual, test)
		assert.Equal(t, expected, actual)
	}
}
