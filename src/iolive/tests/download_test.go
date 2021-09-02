// Package iolive_test
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-02
package iolive_test

import (
	"fmt"
	"time"

	"github.com/teocci/go-csv-samples/src/iolive"
)

func ExampleWriter() {
	writer := iolive.New()

	// start listening to updates and render
	writer.Start()

	for i := 0; i <= 100; i++ {
		_, _ = fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}

	_, _ = fmt.Fprintln(writer, "Finished: Downloaded 100GB")
	writer.Stop() // flush and stop rendering
	// Output:
}
