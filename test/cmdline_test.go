package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	x "github.com/x19290/practice.aa/shline2dos"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", x.List2Cmdline([]string{}))
	assert.Equal(t, "α", x.List2Cmdline([]string{"α"}))
	assert.Equal(t, "α β", x.List2Cmdline([]string{"α", "β"}))
	assert.Equal(t, `"" " "`, x.List2Cmdline([]string{"", " "}))
	assert.Equal(t, `a'`, x.ShlineToDos(`"a"\'`))
}
