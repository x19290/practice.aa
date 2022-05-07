package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	dos "github.com/x19290/practice.aa"
	"github.com/x19290/practice.aa/testdata"
	"strings"
	"testing"
)

func Test0(t *testing.T) {
	for test := range testdata.Stream() {
		expected := strings.Split(test.Expected, "\n")
		actual := actual(test.Feed)
		assert.Equal(t, expected, actual)
	}
}

func actual(feed string) (actual []string) {
	b := new(bytes.Buffer)
	actual = []string{}
	dos.DemoImpl(b, feed)
	out := strings.Split(b.String(), "\n")
	for i := 0; i < len(out) - 2; i += 2 {
		actual = append(actual, out[i + 1])
	}
	return
}