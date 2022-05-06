package main

import (
    "bytes"
    "github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
    ch := make(chan string)
    eb := new(bytes.Buffer)
    ab := new(bytes.Buffer)
    go testdata(ch)
    for {
        _, ok := <-ch
        if !ok {
            break
        }
        exp, ok := <-ch
        if !ok {
            panic("?")
        }
        eb.WriteString(exp)
        eb.WriteRune('\n')
    }
    _main(ab)
    assert.Equal(t, eb.String(), ab.String())
}