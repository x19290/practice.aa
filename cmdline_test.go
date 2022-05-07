package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert.Equal(t, "", List2Cmdline([]string{}))
	assert.Equal(t, "α", List2Cmdline([]string{"α"}))
	assert.Equal(t, "α β", List2Cmdline([]string{"α", "β"}))
}
