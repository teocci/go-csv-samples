// Package slice_internals
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package main

import (
	"fmt"
	"time"
)

type Team []Person
type Person struct {
	Name string
	Age  int
}

type Inventories []Inventory
type Inventory struct { //instead of: map[string]map[string]Pairs
	Warehouse string
	Item      string
	Batches   Lots
}

func (i *Inventories) CloneFrom(c Inventories) {
	inv := new(Inventories)
	for _, v := range c {
		batches := Lots{}
		for _, b := range v.Batches {
			batches = append(batches, Lot{
				Date:  b.Date,
				Key:   b.Key,
				Value: b.Value,
			})
		}

		*inv = append(*inv, Inventory{
			Warehouse: v.Warehouse,
			Item:      v.Item,
			Batches:   batches,
		})
	}

	(*i).ReplaceBy(inv)
}

func (i *Inventories) ReplaceBy(x *Inventories) {
	*i = *x
}

type Lots []Lot
type Lot struct {
	Date  time.Time
	Key   string
	Value float64
}

func main() {
	simpleClone()
	nestedClone()
}

func nestedClone() {
	ins := Inventory{
		Warehouse: "DMM",
		Item:      "Gloves",
		Batches: Lots{
			Lot{time.Now(), "Jan", 50},
			Lot{time.Now().Add(-time.Hour * 24), "Feb", 70},
		},
	}


	inv := Inventories{ins}
	fmt.Printf("Before clonning: %+v\n", inv)

	var inv2 Inventories
	inv2.CloneFrom(inv)
	inv2[0].Warehouse = "DMM 2"

	fmt.Printf("After clonning: %+v\n", inv)
	fmt.Println("------")
	fmt.Printf("Clone slice: %+v\n", inv2)
}

func simpleClone() {
	t := Team{
		Person{"Hasan", 34},
		Person{"Karam", 32},
	}
	fmt.Printf("Before clonning: %+v\n", t)
	fmt.Println("------")
	c := t.Clone()
	fmt.Printf("After clonning: %+v\n", t)
	fmt.Println("------")
	fmt.Printf("Clone slice: %+v\n", c)
}

func (t *Team) Clone() Team {
	var c = make(Team, len(*t))
	copy(c, *t)
	for i, _ := range c {
		c[i].Name = "New Name"
	}

	return c
}
