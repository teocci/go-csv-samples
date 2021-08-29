// +build go1.9

package tests

import (
	"encoding/csv"
	"io"
)

func newCSVReader(r io.Reader) *csv.Reader {
	rr := csv.NewReader(r)
	rr.ReuseRecord = true
	return rr
}

var testUnmarshalInvalidSecondLineErr = &csv.ParseError{
	Line:   2,
	Column: 1,
	Err:    csv.ErrQuote,
}

var ptrUnexportedEmbeddedDecodeErr error