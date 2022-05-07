package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/x19290/practice.aa/shline2dos"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", shline2dos.List2Cmdline([]string{}))
	assert.Equal(t, "α", shline2dos.List2Cmdline([]string{"α"}))
	assert.Equal(t, "α β", shline2dos.List2Cmdline([]string{"α", "β"}))
	assert.Equal(t, `"" " "`, shline2dos.List2Cmdline([]string{"", " "}))
	assert.Equal(t, `a'`, shline2dos.ShlineToDos(`"a"\'`))
}
