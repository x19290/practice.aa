package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/x19290/go.dos/cmdline"
	"github.com/x19290/go.dos/cmdline/testdata"
	"testing"
)

func Test0(t *testing.T) {
	for test := range testdata.LazyList() {
		eb := new(bytes.Buffer)
		ab := new(bytes.Buffer)
		cmdline.DemoCompare(eb, test)
		cmdline.DemoImpl(ab, test)
		expected, actual := eb.String(), ab.String()
		assert.Equal(t, expected, actual)
	}
}
