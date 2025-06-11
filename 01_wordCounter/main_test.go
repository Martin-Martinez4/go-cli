package main

import (
	"bytes"
	"testing"
)

func TestWordCounter(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := "Words: 4\n"

	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %s, but got %s words counted.\n", exp, res)
	}
}

func TestLinecount(t *testing.T) {
	b := bytes.NewBufferString("line1\nline2\nline3\nend")

	exp := "Words: 4\nLines: 4\n"

	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %s, but got %s lines counted.\n", exp, res)
	}

}
