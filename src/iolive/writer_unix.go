// Package iolive
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-02
//go:build !windows
// +build !windows

package iolive

import (
	"fmt"
	"strings"
)

// clear the line and move the cursor up
var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)

func (w *Writer) clearLines() {
	_, _ = fmt.Fprint(w.Out, strings.Repeat(clear, w.lineCount))
}
