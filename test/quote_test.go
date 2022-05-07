package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/x19290/practice.aa/shline2dos"
	"github.com/x19290/practice.aa/testdata"
	"testing"
)

func Test0(t *testing.T) {
	ch := make(chan *testdata.Test)
	go testdata.Stream(ch)
	for test := range ch {
		expected := new(bytes.Buffer)
		actual := new(bytes.Buffer)
		expected.WriteString(test.Expected)
		expected.WriteByte('\n')
		shline2dos.DemoImpl(actual, test.Feed)
		assert.Equal(t, expected, actual)
	}
}
