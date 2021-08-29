// Package sync_pool_bench
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package sync_pool_bench

import (
	"sync"
	"testing"
)
type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func BenchmarkPerson(b *testing.B) {
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
			desc: "without-poll",
			fn: func(b *testing.B, n int) {
				var p *Person
				for j := 0; j < n; j++ {
					p = new(Person)
					p.Age = 23
				}
			},
		},
		{
			desc: "with-poll",
			fn: func(b *testing.B, n int) {
				var p *Person
				for j := 0; j < n; j++ {
					p = personPool.Get().(*Person)
					p.Age = 23
					personPool.Put(p)
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

func BenchmarkMenAlloc(b *testing.B) {
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
			desc: "without-poll",
			fn: func(b *testing.B, n int) {
				for j := 0; j < n; j++ {
					i := 0
					i = i
				}
			},
		},
		{
			desc: "with-poll",
			fn: func(b *testing.B, n int) {
				var p sync.Pool
				for j := 0; j < n; j++ {
					p.Put(1)
					p.Get()
				}
			},
		},
	}

	for _, t := range tests {
		b.Run(t.desc, func(b *testing.B) {
			for _, f := range fixture {
				b.RunParallel(func(pb *testing.PB) {
					b.ReportAllocs()
					b.ResetTimer()
					for pb.Next() {
						t.fn(b, f.records)
					}
				})
			}
		})
	}
}

func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkAllocation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := 0
			i = i
		}
	})
}