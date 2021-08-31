// Package timemgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package timemgr

import (
	"math"
	"time"
)

func UnixTime(gpsTime float32) time.Time {
	sec, dec := math.Modf(float64(gpsTime))
	return time.Unix(int64(sec), int64(dec*1e9))
}

func GenBaseDate(day int) time.Time {
	t := time.Now()
	return time.Date(2021, 8, day, 13, 0, 0, 0, t.Location())
}
