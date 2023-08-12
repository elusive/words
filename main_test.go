package main

import (
    "bytes"
    "testing"
)

const Value string = "word1 word2 word3\n word4 word5\n word6 word7\n word8\n"

func TestCountWords(t *testing.T) {
    b := bytes.NewBufferString(Value)
    exp := 8 

    res := count(b, false, false)

    if (res != exp) {
        t.Errorf("Expected %d, got %d instead.\n", exp, res)
    }
}

func TestCountLines(t *testing.T) {
    b := bytes.NewBufferString(Value)
    exp := 4

    res := count(b, true, false)

    if (res != exp) {
        t.Errorf("Expected %d, got %d instead.\n", exp, res)
    }
}

func TestCountBytes(t *testing.T) {
    b := bytes.NewBufferString(Value)
    exp := 51

    res := count(b, false, true)

    if (res != exp) {
        t.Errorf("Expected %d, got %d instead.\n", exp, res)
    }
}
