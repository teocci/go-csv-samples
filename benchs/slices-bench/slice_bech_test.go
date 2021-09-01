// Package slice_internals
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package slices_bench

import (
	"testing"
)

var gslice = make([]string, 1000)

func global(s string) {
	for i := 0; i < 100; i++ { // Cycle to access slice may times
		_ = s
		_ = gslice // Access global-slice
	}
}

func param(s string, ss []string) {
	for i := 0; i < 100; i++ { // Cycle to access slice may times
		_ = s
		_ = ss // Access parameter-slice
	}
}

func paramPointer(s string, ss *[]string) {
	for i := 0; i < 100; i++ { // Cycle to access slice may times
		_ = s
		_ = ss // Access parameter-slice
	}
}

func BenchmarkPerformance(b *testing.B){
	fixture := []struct {
		desc    string
		records int
	}{
		{
			desc:    "1 record",
			records: 1,
		},
		{
			desc:    "10 records",
			records: 10,
		},
		{
			desc:    "100 records",
			records: 100,
		},
		{
			desc:    "1000 records",
			records: 1000,
		},
		{
			desc:    "10000 records",
			records: 10000,
		},
		{
			desc:    "100000 records",
			records: 100000,
		},
	}

	tests := []struct {
		desc string
		fn   func(b *testing.B, n int)
	}{
		{
			desc: "ParameterPointer",
			fn: func(b *testing.B, n int) {
				for j := 0; j < n; j++ {
					paramPointer("hi", &gslice)
				}
			},
		},
		{
			desc: "Parameter",
			fn: func(b *testing.B, n int) {
				for j := 0; j < n; j++ {
					param("hi", gslice)
				}
			},
		},
		{
			desc: "Global",
			fn: func(b *testing.B, n int) {
				for j := 0; j < n; j++ {
					global("hi")
				}
			},
		},
	}

	for _, t := range tests {
		b.Run(t.desc, func(b *testing.B) {
			for _, f := range fixture {
				b.Run(f.desc, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						t.fn(b, f.records)
					}
				})
			}
		})
	}
}


func BenchmarkParameterPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		paramPointer("hi", &gslice)
	}
}


func BenchmarkParameter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		param("hi", gslice)
	}
}

func BenchmarkGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		global("hi")
	}
}