package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	dos "github.com/x19290/practice.aa"
	"github.com/x19290/practice.aa/testdata"
	"testing"
)

func Test0(t *testing.T) {
	for test := range testdata.Stream() {
		expected := new(bytes.Buffer)
		actual := new(bytes.Buffer)
		expected.WriteString(test.Expected)
		expected.WriteByte('\n')
		dos.DemoImpl(actual, test.Feed)
		assert.Equal(t, expected, actual)
	}
}
