package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/x19290/go.dos/cmdline"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", cmdline.Cmdline())
	assert.Equal(t, "α", cmdline.Cmdline("α"))
	assert.Equal(t, "α β", cmdline.Cmdline("α", "β"))
	assert.Equal(t, `"" " "`, cmdline.Cmdline("", " "))
	assert.Equal(t, `a'`, cmdline.ShlexToCmdline(`"a"\'`))
}
