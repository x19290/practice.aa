package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	dos "github.com/x19290/go.dos"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", dos.Cmdline())
	assert.Equal(t, "α", dos.Cmdline("α"))
	assert.Equal(t, "α β", dos.Cmdline("α", "β"))
	assert.Equal(t, `"" " "`, dos.Cmdline("", " "))
	assert.Equal(t, `a'`, dos.ShlexToCmdline(`"a"\'`))
}
