// Package reflection-get-type
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Ab struct{}

type Potato struct{}

type address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type employee struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Salary  int     `json:"salary"`
	Address address `json:"address"`
}
var (
	pab   = new(Ab)
	ppab  = &pab
	pppab = &ppab

	foos = []interface{}{
		"string",
		10,
		1.2,
		Ab{},
		pab,
		ppab,
		pppab,
	}
)

func main() {
	customReflection()

	usingFMTPackage()

	inJSONForm()
}

func customReflection() {
	for _, r := range foos {
		fmt.Println(getType(r))
	}
	fmt.Println("-----")

	for _, r := range foos {
		fmt.Println(getTypeNoMainFail(r))
	}
	fmt.Println("-----")

	for _, r := range foos {
		fmt.Println(getTypeNoMain(r))
	}
	fmt.Println("-----")
}

func usingFMTPackage() {
	fmt.Printf("I have a %T\n", Potato{})
	fmt.Println("-----")

	emp := employee{Name: "Sam", Age: 31, Salary: 2000}
	fmt.Printf("Emp: %v\n", emp)
	fmt.Printf("Emp: %+v\n", emp)
	fmt.Printf("Emp: %#v\n", emp)
	fmt.Println(emp)
	fmt.Println("-----")
}

func inJSONForm() {
	address := address{City: "some_city", Country: "some_country"}
	emp := employee{Name: "Sam", Age: 31, Salary: 2000, Address: address}
	//Converting to json
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
	fmt.Println("-----")
}

func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func getTypeNoMainFail(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func getTypeNoMain(v interface{}) (res string) {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		res += "*"
	}

	return res + t.Name()
}
