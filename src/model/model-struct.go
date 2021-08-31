// Package model
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package model

import "time"

type FlightSession struct {
	ID          int       `json:"id" csv:"id" pg:"id,pk,unique"`
	DroneID     int       `json:"drone_id" csv:"drone_id" pg:"drone_id"`
	Hash        string    `json:"hash" csv:"hash" sql:"hash" pg:",unique,notnull"`
	GPSStatus   float32   `json:"gps_status" csv:"gps_status" pg:"gps_status"`
	DroneStatus float32   `json:"drone_status" csv:"drone_status" pg:"drone_status"`
	LastUpdate  time.Time `json:"last_update" csv:"last_update" pg:"last_update"`
}

type FlightSessionReading struct {
	ID              int       `json:"id" csv:"id" pg:"id,pk,unique"`
	DroneID         int       `json:"drone_id" csv:"drone_id" pg:"drone_id"`
	FlightSessionID int       `json:"flight_session_id" pg:"flight_session_id" pg:"flight_session_id"`
	Latitude        float32   `json:"latitude" csv:"lat" pg:"latitude"`
	Longitude       float32   `json:"longitude" csv:"long" pg:"longitude"`
	Altitude        float32   `json:"altitude" csv:"alt" pg:"altitude"`
	Roll            float32   `json:"roll" csv:"roll" pg:"roll"`
	Pitch           float32   `json:"pitch" csv:"pitch" pg:"pitch"`
	Yaw             float32   `json:"yaw" csv:"yaw" pg:"yaw"`
	BatVoltage      float32   `json:"battery_voltage" csv:"battery_voltage" pg:"battery_voltage"`
	BatCurrent      float32   `json:"battery_current" csv:"battery_current" pg:"battery_current"`
	BatPercent      float32   `json:"battery_percentage" csv:"battery_percentage" pg:"battery_percentage"`
	BatTemperature  float32   `json:"battery_temperature" csv:"battery_temperature" pg:"battery_temperature"`
	Temperature     float32   `json:"temperature" csv:"temperature" pg:"temperature"`
	LastUpdate      time.Time `json:"last_update" csv:"last_update" pg:"last_update"`
}
