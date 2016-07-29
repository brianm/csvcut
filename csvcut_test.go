package main

import (
	"testing"
	"strings"
	"bytes"
)

func TestAllFieldsCommaDelim(t *testing.T) {
	in := strings.NewReader("\"a\t\",b,c\n1,2,3\n")
	out := bytes.Buffer{}
	opts := options {
		delimiter: rune(','),
	}
	process(opts, in, &out)
	output := string(out.Bytes())
	if output != "\"a\t\"\tb\tc\n1\t2\t3\n" {
		t.Fatalf("expected tab delimited output, got %s", output)
	}
}

func TestOneFieldCommaDelim(t *testing.T) {
	in := strings.NewReader("\"a\t\",b,c\n1,2,3\n")
	out := bytes.Buffer{}
	opts := options {
		delimiter: rune(','),
		fields: []int {1},
	}
	process(opts, in, &out)
	output := string(out.Bytes())
	if output != "b\n2\n" {
		t.Fatalf("expected tab delimited output, got %s", output)
	}
}

func TestTwoFieldCommaDelim(t *testing.T) {
	in := strings.NewReader("\"a\t\",b,c\n1,2,3\n")
	out := bytes.Buffer{}
	opts := options {
		delimiter: rune(','),
		fields: []int {2,1},
	}
	process(opts, in, &out)
	output := string(out.Bytes())
	if output != "c\tb\n3\t2\n" {
		t.Fatalf("expected two column, tab delimited output, got %s", output)
	}
}
