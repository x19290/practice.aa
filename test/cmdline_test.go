package test

import (
	"github.com/stretchr/testify/assert"
	dos "github.com/x19290/go.dos"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", dos.List2Cmdline([]string{}))
	assert.Equal(t, "α", dos.List2Cmdline([]string{"α"}))
	assert.Equal(t, "α β", dos.List2Cmdline([]string{"α", "β"}))
	assert.Equal(t, `"" " "`, dos.List2Cmdline([]string{"", " "}))
	assert.Equal(t, `a'`, dos.ShlexToCmdline(`"a"\'`))
}
