package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	dos "github.com/x19290/go.dos"
	"github.com/x19290/go.dos/testdata"
	"testing"
)

func Test0(t *testing.T) {
	for test := range testdata.LazyList() {
		eb := new(bytes.Buffer)
		ab := new(bytes.Buffer)
		dos.DemoCompare(eb, test)
		dos.DemoImpl(ab, test)
		expected, actual := eb.String(), ab.String()
		assert.Equal(t, expected, actual)
	}
}
