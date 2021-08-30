// Package data
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package data

import (
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
	DroneID         int     `json:"id" csv:"id" name:"id"`
	FlightSessionID int     `json:"flight_session_id" csv:"flight_session_id" name:"flight_session_id"`
	Lat             float32 `json:"lat" csv:"lat" name:"lat"`
	Long            float32 `json:"long" csv:"lat" name:"lat"`
	Alt             float32 `json:"alt" csv:"lat" name:"lat"`
	Roll            float32 `json:"roll" csv:"lat" name:"lat"`
	Pitch           float32 `json:"pitch" csv:"lat" name:"lat"`
	Yaw             float32 `json:"yaw" csv:"lat" name:"lat"`
	BatVoltage      float32 `json:"battery_voltage" csv:"battery_voltage" name:"battery_voltage"`
	BatCurrent      float32 `json:"battery_current" csv:"battery_current" name:"battery_current"`
	BatPercent      float32 `json:"battery_percentage" csv:"battery_percentage" name:"battery_percentage"`
	BatTemperature  float32 `json:"battery_temperature" csv:"battery_temperature" name:"battery_temperature"`
	Temperature     float32 `json:"temperature" csv:"temperature" name:"temperature"`
	GPSTime         float32 `json:"modify_date" csv:"modify_date" name:"modify_date"`
}

func ParseGEOData(data []string) *GEOData {
	fccTime, _ := strconv.ParseFloat(data[0], 64)
	lat, _ := strconv.ParseFloat(data[1], 64)
	long, _ := strconv.ParseFloat(data[2], 64)
	alt, _ := strconv.ParseFloat(data[3], 64)
	roll, _ := strconv.ParseFloat(data[4], 64)
	pitch, _ := strconv.ParseFloat(data[5], 64)
	yaw, _ := strconv.ParseFloat(data[6], 64)

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

func ParseFCC(data []string) *FCC {
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

func ParseRTT(data []string) *RTT {
	droneID, _ := strconv.Atoi(data[1])
	sessionID, _ := strconv.Atoi(data[2])
	lat, _ := strconv.ParseFloat(data[3], 64)
	long, _ := strconv.ParseFloat(data[4], 64)
	alt, _ := strconv.ParseFloat(data[5], 64)
	roll, _ := strconv.ParseFloat(data[6], 64)
	pitch, _ := strconv.ParseFloat(data[7], 64)
	yaw, _ := strconv.ParseFloat(data[8], 64)
	temp, _ := strconv.ParseFloat(data[9], 64)
	batVol, _ := strconv.ParseFloat(data[10], 64)
	batCurr, _ := strconv.ParseFloat(data[11], 64)
	batPct, _ := strconv.ParseFloat(data[12], 64)
	batTemp, _ := strconv.ParseFloat(data[13], 64)
	gpsTime, _ := strconv.ParseFloat(data[14], 64)

	return &RTT{
		DroneID:         droneID,
		FlightSessionID: sessionID,
		Lat:             float32(lat),
		Long:            float32(long),
		Alt:             float32(alt),
		Roll:            float32(roll),
		Pitch:           float32(pitch),
		Yaw:             float32(yaw),
		BatVoltage:      float32(batVol),
		BatCurrent:      float32(batCurr),
		BatPercent:      float32(batPct),
		BatTemperature:  float32(batTemp),
		Temperature:     float32(temp),
		GPSTime:         float32(gpsTime),
	}
}
