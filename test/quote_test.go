package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	x "github.com/x19290/practice.aa/shline2dos"
)

func Test0(t *testing.T) {
	ch := make(chan *x.Test)
	go x.TestData(ch)
	for test := range ch {
		expected := new(bytes.Buffer)
		actual := new(bytes.Buffer)
		expected.WriteString(test.Expected)
		expected.WriteByte('\n')
		x.DemoImpl(actual, test.Feed)
		assert.Equal(t, expected, actual)
	}
}
