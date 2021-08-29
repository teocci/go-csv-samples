// Package data
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package data

import (
	"fmt"
	"strconv"
)

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

type RTT struct {
	DroneID         int     `json:"drone_id" csv:"drone_id" name:"drone_id"`
	FlightSessionID int     `json:"flight_session_id" csv:"flight_session_id" name:"flight_session_id"`
	DroneLat        float32 `json:"drone_lat" csv:"drone_lat" name:"drone_lat"`
	DroneLong       float32 `json:"drone_long" csv:"drone_lat" name:"drone_lat"`
	DroneAlt        float32 `json:"drone_alt" csv:"drone_lat" name:"drone_lat"`
	DroneRoll       float32 `json:"drone_roll" csv:"drone_lat" name:"drone_lat"`
	DronePitch      float32 `json:"drone_pitch" csv:"drone_lat" name:"drone_lat"`
	DroneYaw        float32 `json:"drone_yaw" csv:"drone_lat" name:"drone_lat"`
	BatVoltage      float32 `json:"battery_voltage" csv:"battery_voltage" name:"battery_voltage"`
	BatCurrent      float32 `json:"battery_current" csv:"battery_current" name:"battery_current"`
	BatPercent      float32 `json:"battery_percentage" csv:"battery_percentage" name:"battery_percentage"`
	BatTemperature  float32 `json:"battery_temperature" csv:"battery_temperature" name:"battery_temperature"`
	Temperature     float32 `json:"temperature" csv:"temperature" name:"temperature"`
	GPSTime         float32 `json:"modify_date" csv:"modify_date" name:"modify_date"`
}

func ParseGEODataStruct(data []string) *GEOData {
	fccTime, _ := strconv.ParseFloat(data[0], 64)
	lat, _ := strconv.ParseFloat(data[1], 64)
	long, _ := strconv.ParseFloat(data[2], 64)
	alt, _ := strconv.ParseFloat(data[3], 64)
	roll, _ := strconv.ParseFloat(data[4], 64)
	pitch, _ := strconv.ParseFloat(data[5], 64)
	yaw, _ := strconv.ParseFloat(data[6], 64)

	fmt.Printf("%+v\n", lat)

	return &GEOData{
		FCCTime: float32(fccTime),
		Lat:     float32(lat),
		Long:    float32(long),
		Alt:     float32(alt),
		Roll:    float32(roll),
		Pitch:   float32(pitch),
		Yaw:     float32(yaw),
	}
}

func ParseFCCStruct(data []string) *FCC {
	fccTime, _ := strconv.ParseFloat(data[0], 64)
	gpsTime, _ := strconv.ParseFloat(data[1], 64)
	temp, _ := strconv.ParseFloat(data[2], 64)
	batVol, _ := strconv.ParseFloat(data[3], 64)
	batCurr, _ := strconv.ParseFloat(data[4], 64)
	batPct, _ := strconv.ParseFloat(data[5], 64)
	batTemp, _ := strconv.ParseFloat(data[6], 64)

	return &FCC{
		FCCTime:        float32(fccTime),
		GPSTime:        float32(gpsTime),
		Temperature:    float32(temp),
		BatVoltage:     float32(batVol),
		BatCurrent:     float32(batCurr),
		BatPercent:     float32(batPct),
		BatTemperature: float32(batTemp),
	}
}
