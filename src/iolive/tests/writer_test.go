// Package iolive_test
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-02
package iolive_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/teocci/go-csv-samples/src/iolive"
)

func TestWriter(t *testing.T) {
	w := iolive.New()
	b := &bytes.Buffer{}
	w.Out = b
	w.Start()
	for i := 0; i < 2; i++ {
		_, _ = fmt.Fprintln(w, "foo")
	}
	w.Stop()
	_, _ = fmt.Fprintln(b, "bar")

	want := "foo\nfoo\nbar\n"
	if b.String() != want {
		t.Fatalf("want %q, got %q", want, b.String())
	}
}

func TestStartCalledTwice(t *testing.T) {
	w := iolive.New()
	b := &bytes.Buffer{}
	w.Out = b

	w.Start()
	w.Stop()
	w.Start()
	w.Stop()
}
