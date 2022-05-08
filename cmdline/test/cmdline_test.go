package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/x19290/go.dos/cmdline"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", cmdline.Make())
	assert.Equal(t, "α", cmdline.Make("α"))
	assert.Equal(t, "α β", cmdline.Make("α", "β"))
	assert.Equal(t, `"" " "`, cmdline.Make("", " "))
	assert.Equal(t, `a'`, cmdline.FromShlex(`"a"\'`))
}
