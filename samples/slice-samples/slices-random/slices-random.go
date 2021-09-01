// Package slices_random
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	data := []point{
		{1, 2}, {3, 4},
		{5, 6}, {7, 8},
	}

	makeRandomData(&data)
}

func makeRandomData(dataPoints *[]point) {
	for i := 0; i < 10; i++ {
		if len(*dataPoints) > 0 {
			fmt.Println(generate(dataPoints))
		} else {
			fmt.Println("no more elements")
		}
	}
}

func generate(cities *[]point) []point {
	//create a new slice with the first item from the old slice
	solution := []point{(*cities)[0]}
	//remove the first item from the old slice
	*cities = append((*cities)[:0], (*cities)[1:]...)

	return solution
}
