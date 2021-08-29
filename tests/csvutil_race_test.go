// Package main
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package tests


import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/jszwec/csvutil"
)


func TestCacheDataRaces(t *testing.T) {
	const routines = 16
	const rows = 1000000

	v := TypeF{
		Int:      1,
		Pint:     pint(2),
		Int8:     3,
		Pint8:    pint8(4),
		Int16:    5,
		Pint16:   pint16(6),
		Int32:    7,
		Pint32:   pint32(8),
		Int64:    9,
		Pint64:   pint64(10),
		UInt:     11,
		Puint:    puint(12),
		Uint8:    13,
		Puint8:   puint8(14),
		Uint16:   15,
		Puint16:  puint16(16),
		Uint32:   17,
		Puint32:  puint32(18),
		Uint64:   19,
		Puint64:  puint64(20),
		Float32:  21,
		Pfloat32: pfloat32(22),
		Float64:  23,
		Pfloat64: pfloat64(24),
		String:   "25",
		PString:  pstring("26"),
		Bool:     true,
		Pbool:    pbool(true),
		V:        ppint(100),
		Pv:       pinterface(ppint(200)),
		Binary:   Binary,
		PBinary:  &Binary,
	}

	t.Run("encoding", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < routines; i++ {
			tag := "csv"
			if i%2 == 0 {
				tag = "custom"
			}

			wg.Add(1)
			go func() {
				defer wg.Done()

				var buf bytes.Buffer
				w := csv.NewWriter(&buf)
				enc := csvutil.NewEncoder(w)
				enc.Tag = tag
				for i := 0; i < rows; i++ {
					if err := enc.Encode(v); err != nil {
						panic(err)
					}
					fmt.Println("Encoder:", i)
				}
				w.Flush()
			}()
		}
		wg.Wait()
	})

	t.Run("decoding", func(t *testing.T) {
		vs := make([]*TypeF, 0, rows)
		for i := 0; i < rows; i++ {
			vs = append(vs, &v)
		}

		data, err := csvutil.Marshal(vs)
		if err != nil {
			t.Fatal(err)
		}

		var wg sync.WaitGroup
		for i := 0; i < routines; i++ {
			tag := "csv"
			if i%2 == 0 {
				tag = "custom"
			}

			wg.Add(1)
			go func() {
				defer wg.Done()

				dec, err := csvutil.NewDecoder(csv.NewReader(bytes.NewReader(data)))
				if err != nil {
					t.Fatal(err)
				}
				dec.Tag = tag
				var i int
				for {
					var val TypeF
					if err := dec.Decode(&val); err == io.EOF {
						break
					} else if err != nil {
						panic(err)
					}
					fmt.Println("Decode:", i)
					i++
				}
			}()
		}
		wg.Wait()
	})
}