// Package live_output
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-02
package main

import (
	"fmt"
	"time"

	"github.com/teocci/go-csv-samples/src/iolive"
)

func main() {
	writer := iolive.New()
	// start listening for updates and render
	writer.Start()

	for i := 0; i <= 100; i++ {
		fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Fprintln(writer, "Finished: Downloaded 100GB")
	writer.Stop() // flush and stop rendering
}
