// Package csvutil_merge
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package main

import (
	"fmt"
	"github.com/teocci/go-csv-samples/src/csvmgr"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/teocci/go-csv-samples/src/data"
)

func main() {
	var geos []data.GEOData
	// open the first file
	geoBuff := csvmgr.LoadDataBuff(data.GEOPath)
	if err := gocsv.UnmarshalBytes(geoBuff, &geos); err != nil {
		log.Fatal(err)
	}

	var fccs []data.FCC
	// open the first file
	fccBuff := csvmgr.LoadDataBuff(data.FCCPath)
	if err := gocsv.UnmarshalBytes(fccBuff, &fccs); err != nil {
		log.Fatal(err)
	}

	var rtts []data.FCC
	// create a file writer
	rttFN := data.RTTPrefix + "_RTTdata"
	fmt.Println("rttFN:", rttFN)
	rttPath := filepath.Join(data.DestPath, rttFN+".csv")
	_ = rtts

	w := csvmgr.CreateFile(rttPath)
	defer csvmgr.CloseFile()(w)

	for _, geo := range geos {
		var last int
		var rtt data.RTT
		for i := last; i < len(fccs); i++ {
			if geo.FCCTime == fccs[i].FCCTime {
				fcc := fccs[i]
				last = i
				rtt = data.RTT{
					DroneID: 1,
					FlightSessionID: 1,
					DroneLat: geo.Lat,
					DroneLong: geo.Long,
					DroneAlt: geo.Alt,
					DroneRoll: geo.Roll,
					DronePitch: geo.Pitch,
					DroneYaw: geo.Yaw,
					BatVoltage: fcc.BatVoltage,
					BatCurrent: fcc.BatCurrent,
					BatPercent: fcc.BatPercent,
					BatTemperature: fcc.BatTemperature,
					Temperature: fcc.Temperature,
					GPSTime: fcc.GPSTime,
				}

				_ = rtt

				sec, dec := math.Modf(float64(fcc.FCCTime))
				t := time.Unix(int64(sec), int64(dec*(1e3)))

				fmt.Printf("%+v\n", t)
			}
		}
	}


}
