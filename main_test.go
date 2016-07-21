package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	test := "golang.org/x/net/trace"
	test2 := "github.com/odinliu/got"

	assert.Equal(
		t, "github.com/golang/net/trace",
		replace(test), "they should be equal",
	)
	assert.Equal(t, replace(test2), test2, "they should be equal")
}

func TestMakeCmd(t *testing.T) {
	test := []string{
		"got", "-u", "-v", "golang.org/x/net/trace",
	}
	target := []string{
		"get", "-u", "-v", "github.com/golang/net/trace",
	}

	assert.Equal(t, target, makeCmd(test), "they should be equal")
}
