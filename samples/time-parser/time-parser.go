// Package time_parser
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-31
package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

var (
	fccTimes = []float64{
		190.108, 190.124, 190.144, 190.164, 190.184, 190.204, 190.224, 190.244, 190.264, 190.284,
		190.308, 190.324, 190.344, 190.364, 190.384, 190.404, 190.424, 190.444, 190.464, 190.484,
		190.508, 190.524, 190.544, 190.564, 190.584, 190.604, 190.624, 190.644, 190.664, 190.684,
		190.708, 190.724, 190.744, 190.764, 190.784, 190.804, 190.824, 190.844, 190.864, 190.884,
		190.908, 190.924, 190.944, 190.964, 190.984, 191.004, 191.024, 191.044, 191.064, 191.084,
		191.108, 191.124, 191.144, 191.164, 191.184, 191.204, 191.224, 191.244, 191.264, 191.284,
		191.308, 191.324, 191.344, 191.364, 191.384, 191.404, 191.424, 191.444, 191.464, 191.484,
		191.508, 191.524, 191.544, 191.564, 191.584, 191.604, 191.624, 191.644, 191.664, 191.684,
		191.708, 191.724, 191.744, 191.764, 191.784, 191.804, 191.824, 191.844, 191.864, 191.884,
	}

	gpsTimes = []float64{
		189061.406,
		189061.594,
		189061.797,
		189062.0,
		189062.203,
		189062.406,
		189062.594,
		189062.797,
		189063.016,
	}

	gcsTime = []int64{
		1630391340169,
		1630391341182,
		1630391342177,
		1630391343175,
		1630391344174,
		1630391345170,
		1630391346169,
		1630391347181,
		1630391348168,
		1630391349181,
		1630391350178,
		1630391351167,
		1630391352172,
		1630391353177,
		1630391354170,
		1630391355168,
		1630391356179,
		1630391357178,
		1630391358174,
		1630391359172,
		1630391360170,
		1630391361180,
		1630391362174,
		1630391363167,
		1630391364177,
		1630391365167,
		1630391366171,
		1630391367176,
		1630391368172,
		1630391369181,
		1630391370180,
		1630391371168,
		1630391372178,
		1630391373170,
	}
)

func main() {
	d := 61 * time.Microsecond
	fmt.Println(d) // Output: 61µs

	ms := float64(d) / float64(time.Millisecond)
	fmt.Println("ms:", ms) // Output: ms: 100

	fcc1 := 496.604
	fmt.Println(reflect.TypeOf(fcc1))
	sec, dec := math.Modf(fcc1)
	fcct1 := time.Unix(int64(sec), int64(dec*1e9))
	fmt.Println("fcct1 :", fcct1)

	fcc2 := 499.108
	sec, dec = math.Modf(fcc2)
	fcct2 := time.Unix(int64(sec), int64(dec*1e6))
	fmt.Println("fcct2 :", fcct2)

	fmt.Printf("The call took %v to run.\n", fcct2.Sub(fcct1))
	fmt.Println("-----")

	var fccBase time.Time
	var count int
	for i, v := range fccTimes {
		sec, dec = math.Modf(v)
		curr := time.Unix(int64(sec), int64(dec*1e9))
		if count == 0 {
			fccBase = curr
			fmt.Println("fccBase was set :", fccBase)
		}
		count++
		if count == 10 {
			count = 0
		}

		fmt.Printf("%d | The call took %v to run.\n", i, curr.Sub(fccBase))
	}
	fmt.Println("-----")

	var gpsBase time.Time
	for i, v := range gpsTimes {
		sec, dec = math.Modf(v)
		curr := time.Unix(int64(sec), int64(dec*1e9))
		if i == 0 {
			gpsBase = curr
			fmt.Println("gpsBase was set :", gpsBase)
		}

		fmt.Printf("%d | The call took %v to run.\n", i, curr.Sub(gpsBase))
	}
	fmt.Println("-----")

	gpsTimesToString(gcsTime)

	sinceEpoch()

	unixNano()

	gpsEpoch()

	genBaseDate()
}

func genBaseDate() {
	t := time.Now()
	t = time.Date(2021, 8, 1, 13, 0, 0, 0, t.Location())
	fmt.Println(t.Format("2006-01-02, 15:04:05"))
}

func gpsTimesToString(gcsTimes []int64) {
	var gcsBase time.Time
	for i, v := range gcsTimes {
		curr := time.Unix(v%1e6, v*1e6)
		fmt.Println(curr.Format("2006.01.02, 15:04:05"))
		if i == 0 {
			gcsBase = curr
			fmt.Println("gcsBase was set :", gcsBase)
		}

		fmt.Printf("%d | The call took %v to run.\n", i, curr.Sub(gcsBase))
	}
	fmt.Println("-----")
}

func sinceEpoch() {
	fmt.Println("sinceEpoch")
	tms := time.Now().Sub(time.Unix(0, 0)).Milliseconds()
	fmt.Printf("Unix time: %d ms\n", tms)
	fmt.Println("-----")
}

func unixNano() {
	fmt.Println("unixNano")
	un := time.Now().UnixNano()
	fmt.Println(un) // prints: 1630394660994421800
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(timestamp) // prints: 1630394660994421800
	fmt.Println("-----")
}

func gpsEpoch() {
	t := time.Date(1980, time.January, 6, 0, 0, 0, 0, time.UTC)
	fmt.Println(t.Unix())
}

func friday13(t time.Time) time.Time {
	const day = 24 * time.Hour
	// get daylight saving time out of the way
	t = time.Date(t.Year(), 8, t.Day(), 12, 0, 0, 0, t.Location())
	// compute next Friday
	t = t.Add(6 * day)
	t = t.Add(-time.Duration(t.Add(-5*day).Weekday()) * day)
	// check all subsequent Fridays
	for ; t.Day() != 13; t = t.Add(7 * day) {
	}
	return t
}
