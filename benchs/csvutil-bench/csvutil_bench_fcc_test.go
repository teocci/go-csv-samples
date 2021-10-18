// Package csvutil_bench
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package csvutil_bench

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	"github.com/yunabe/easycsv"
	csvgo "trimmer.io/go-csv"
)

func BenchmarkFCCUnmarshal(b *testing.B) {
	type GEOData struct {
		FCCTime float32 `json:"fcc_time" csv:"FCCTime" name:"FCCTime"`
		Lat     float32 `json:"lat" csv:"Lat" name:"Latitude"`
		Long    float32 `json:"long" csv:"Long" name:"Longitude"`
		Alt     float32 `json:"alt" csv:"Alt" name:"Altitude"`
		Roll    float32 `json:"roll" csv:"Roll" name:"Roll"`
		Pitch   float32 `json:"pitch" csv:"Pitch" name:"Pitch"`
		Yaw     float32 `json:"yaw" csv:"Yaw" name:"Yaw"`
	}

	type FCC struct {
		FCCTime        float32 `json:"fcc_time" csv:"FCCTime" name:"FCCTime"`
		GPSTime        float32 `json:"gps_time" csv:"GPSTime" name:"GPSTime"`
		Temperature    float32 `json:"temperature" csv:"Temp" name:"Temperature"`
		BatVoltage     float32 `json:"bat_voltage" csv:"Bat" name:"BatVoltage"`
		BatCurrent     float32 `json:"bat_current" csv:"BatCurr" name:"BatCurrent"`
		BatPercent     float32 `json:"bat_percent" csv:"BatPercent" name:"BatPercent"`
		BatTemperature float32 `json:"bat_temperature" csv:"BatTemp" name:"BatTemperature"`
	}

	fixture := []struct {
		desc    string
		records string
		isFCC   bool
	}{
		{
			desc:    "GEOData",
			records: "../../tmp/GEOdata.csv",
			isFCC:   false,
		},
		{
			desc:    "FCC",
			records: "../../tmp/FCC.csv",
			isFCC:   true,
		},
	}

	tests := []struct {
		desc string
		fn   func(b *testing.B, data []byte, isFF bool)
	}{
		{
			desc: "csvutil.Unmarshal",
			fn: func(b *testing.B, data []byte, isFF bool) {
				if isFF {
					var rec []FCC
					if err := csvutil.Unmarshal(data, &rec); err != nil {
						b.Error(err)
					}
				} else {
					var rec []GEOData
					if err := csvutil.Unmarshal(data, &rec); err != nil {
						b.Error(err)
					}
				}
			},
		},
		{
			desc: "csvgo.Unmarshal",
			fn: func(b *testing.B, data []byte, isFF bool) {
				if isFF {
					var rec []FCC
					if err := csvgo.Unmarshal(data, &rec); err != nil {
						b.Error(err)
					}
				} else {
					var rec []GEOData
					if err := csvgo.Unmarshal(data, &rec); err != nil {
						b.Error(err)
					}
				}
			},
		},
		{
			desc: "gocsv.Unmarshal",
			fn: func(b *testing.B, data []byte, isFF bool) {
				if isFF {
					var rec []FCC
					if err := gocsv.UnmarshalBytes(data, &rec); err != nil {
						b.Error(err)
					}
				} else {
					var rec []GEOData
					if err := gocsv.UnmarshalBytes(data, &rec); err != nil {
						b.Error(err)
					}
				}
			},
		},
		{
			desc: "easycsv.ReadAll",
			fn: func(b *testing.B, data []byte, isFF bool) {
				r := easycsv.NewReader(bytes.NewReader(data))
				if isFF {
					var rec []FCC
					if err := r.ReadAll(&rec); err != nil {
						b.Error(err)
					}
				} else {
					var rec []GEOData
					if err := r.ReadAll(&rec); err != nil {
						b.Error(err)
					}
				}
			},
		},
	}

	for _, t := range tests {
		b.Run(t.desc, func(b *testing.B) {
			for _, f := range fixture {
				b.Run(f.desc, func(b *testing.B) {
					//dir, err := os.Getwd()
					//if err != nil {
					//	b.Fatal(err)
					//}
					//b.Log(dir)
					data := genData(f.records)
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						t.fn(b, data, f.isFCC)
					}
				})
			}
		})
	}
}

func genData(f string) []byte {
	var file, err = os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(file); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
