// Package gpstm_test
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-31
package gpstm_test

import (
	"fmt"
	"github.com/teocci/go-csv-samples/src/gpstm"
	"time"
)

// Display the GPS time of the given time, in microseconds.
func ExampleGpsTime() {
	fmt.Println(int64(gpstm.GpsTime(time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC)).Gps() / time.Microsecond))
	// Output: 948731799000000
}
