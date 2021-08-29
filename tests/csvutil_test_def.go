// Package tests
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package tests

import (
	"bytes"
	"encoding"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jszwec/csvutil"
	"reflect"
	"strconv"
	"unicode"
)

var Binary = []byte("binary-data")

var EncodedBinary = base64.StdEncoding.EncodeToString(Binary)

var BinaryLarge = bytes.Repeat([]byte("1"), 128*1024)

var EncodedBinaryLarge = base64.StdEncoding.EncodeToString(BinaryLarge)

type Float float64

type Enum uint8

const (
	EnumDefault = iota
	EnumFirst
	EnumSecond
)



func (e Enum) MarshalCSV() ([]byte, error) {
	switch e {
	case EnumFirst:
		return []byte("first"), nil
	case EnumSecond:
		return []byte("second"), nil
	default:
		return []byte("default"), nil
	}
}

func (e *Enum) UnmarshalCSV(data []byte) error {
	s := string(data)
	switch s {
	case "first":
		*e = EnumFirst
	case "second":
		*e = EnumSecond
	default:
		*e = EnumDefault
	}
	return nil
}

type ValueRecUnmarshaler struct {
	S *string
}

func (u ValueRecUnmarshaler) UnmarshalCSV(data []byte) error {
	*u.S = string(data)
	return nil
}

func (u ValueRecUnmarshaler) Scan(data []byte) error {
	*u.S = "scan: "
	*u.S += string(data)
	return nil
}

type ValueRecTextUnmarshaler struct {
	S *string
}

func (u ValueRecTextUnmarshaler) UnmarshalText(text []byte) error {
	*u.S = string(text)
	return nil
}

type IntStruct struct {
	Value int
}

func (i *IntStruct) Scan(state fmt.ScanState, verb rune) error {
	switch verb {
	case 'd', 'v':
	default:
		return errors.New("unsupported verb")
	}

	t, err := state.Token(false, unicode.IsDigit)
	if err != nil {
		return err
	}

	n, err := strconv.Atoi(string(t))
	if err != nil {
		return err
	}
	*i = IntStruct{Value: n}
	return nil
}

type EnumType struct {
	Enum Enum `csv:"enum"`
}

type Embedded1 struct {
	String string  `csv:"string"`
	Float  float64 `csv:"float"`
}

type Embedded2 struct {
	Float float64 `csv:"float"`
	Bool  bool    `csv:"bool"`
}

type Embedded3 map[string]string

func (e *Embedded3) UnmarshalCSV(s []byte) error {
	return json.Unmarshal(s, e)
}

func (e Embedded3) MarshalCSV() ([]byte, error) {
	return json.Marshal(e)
}

type Embedded4 interface{}

type Embedded5 struct {
	Embedded6
	Embedded7
	Embedded8
}

type Embedded6 struct {
	X int
}

type Embedded7 Embedded6

type Embedded8 struct {
	Embedded9
}

type Embedded9 struct {
	X int
	Y int
}

type Embedded10 struct {
	Embedded11
	Embedded12
	Embedded13
}

type Embedded11 struct {
	Embedded6
}

type Embedded12 struct {
	Embedded6
}

type Embedded13 struct {
	Embedded8
}

type Embedded17 struct {
	*Embedded18
}

type Embedded18 struct {
	X *float64
	Y *float64
}

type TypeA struct {
	Embedded1
	String string `csv:"string"`
	Embedded2
	Int int `csv:"int"`
}

type TypeB struct {
	Embedded3 `csv:"json"`
	String    string `csv:"string"`
}

type TypeC struct {
	*Embedded1
	String string `csv:"string"`
}

type TypeD struct {
	*Embedded3 `csv:"json"`
	String     string `csv:"string"`
}

type TypeE struct {
	String **string `csv:"string"`
	Int    *int     `csv:"int"`
}

type TypeF struct {
	Int      int          `csv:"int" custom:"int"`
	Pint     *int         `csv:"pint" custom:"pint"`
	Int8     int8         `csv:"int8" custom:"int8"`
	Pint8    *int8        `csv:"pint8" custom:"pint8"`
	Int16    int16        `csv:"int16" custom:"int16"`
	Pint16   *int16       `csv:"pint16" custom:"pint16"`
	Int32    int32        `csv:"int32" custom:"int32"`
	Pint32   *int32       `csv:"pint32" custom:"pint32"`
	Int64    int64        `csv:"int64" custom:"int64"`
	Pint64   *int64       `csv:"pint64" custom:"pint64"`
	UInt     uint         `csv:"uint" custom:"uint"`
	Puint    *uint        `csv:"puint" custom:"puint"`
	Uint8    uint8        `csv:"uint8" custom:"uint8"`
	Puint8   *uint8       `csv:"puint8" custom:"puint8"`
	Uint16   uint16       `csv:"uint16" custom:"uint16"`
	Puint16  *uint16      `csv:"puint16" custom:"puint16"`
	Uint32   uint32       `csv:"uint32" custom:"uint32"`
	Puint32  *uint32      `csv:"puint32" custom:"puint32"`
	Uint64   uint64       `csv:"uint64" custom:"uint64"`
	Puint64  *uint64      `csv:"puint64" custom:"puint64"`
	Float32  float32      `csv:"float32" custom:"float32"`
	Pfloat32 *float32     `csv:"pfloat32" custom:"pfloat32"`
	Float64  float64      `csv:"float64" custom:"float64"`
	Pfloat64 *float64     `csv:"pfloat64" custom:"pfloat64"`
	String   string       `csv:"string" custom:"string"`
	PString  *string      `csv:"pstring" custom:"pstring"`
	Bool     bool         `csv:"bool" custom:"bool"`
	Pbool    *bool        `csv:"pbool" custom:"pbool"`
	V        interface{}  `csv:"interface" custom:"interface"`
	Pv       *interface{} `csv:"pinterface" custom:"pinterface"`
	Binary   []byte       `csv:"binary" custom:"binary"`
	PBinary  *[]byte      `csv:"pbinary" custom:"pbinary"`
}

type TypeG struct {
	String      string
	Int         int
	Float       float64 `csv:"-"`
	unexported1 int
	unexported2 int `csv:"unexported2"`
}

type TypeI struct {
	String string `csv:",omitempty"`
	Int    int    `csv:"int,omitempty"`
}

type TypeK struct {
	*TypeL
}

type TypeL struct {
	String string
	Int    int `csv:",omitempty"`
}

type Unmarshalers struct {
	CSVUnmarshaler      CSVUnmarshaler      `csv:"csv"`
	PCSVUnmarshaler     *CSVUnmarshaler     `csv:"pcsv"`
	TextUnmarshaler     TextUnmarshaler     `csv:"text"`
	PTextUnmarshaler    *TextUnmarshaler    `csv:"ptext"`
	CSVTextUnmarshaler  CSVTextUnmarshaler  `csv:"csv-text"`
	PCSVTextUnmarshaler *CSVTextUnmarshaler `csv:"pcsv-text"`
}

type EmbeddedUnmarshalers struct {
	CSVUnmarshaler     `csv:"csv"`
	TextUnmarshaler    `csv:"text"`
	CSVTextUnmarshaler `csv:"csv-text"`
}

type EmbeddedPtrUnmarshalers struct {
	*CSVUnmarshaler     `csv:"csv"`
	*TextUnmarshaler    `csv:"text"`
	*CSVTextUnmarshaler `csv:"csv-text"`
}

type CSVUnmarshaler struct {
	String string `csv:"string"`
}

func (t *CSVUnmarshaler) UnmarshalCSV(s []byte) error {
	t.String = "unmarshalCSV:" + string(s)
	return nil
}

type TextUnmarshaler struct {
	String string `csv:"string"`
}

func (t *TextUnmarshaler) UnmarshalText(text []byte) error {
	t.String = "unmarshalText:" + string(text)
	return nil
}

type CSVTextUnmarshaler struct {
	String string `csv:"string"`
}

func (t *CSVTextUnmarshaler) UnmarshalCSV(s []byte) error {
	t.String = "unmarshalCSV:" + string(s)
	return nil
}

func (t *CSVTextUnmarshaler) UnmarshalText(text []byte) error {
	t.String = "unmarshalText:" + string(text)
	return nil
}

type TypeWithInvalidField struct {
	String TypeI `csv:"string"`
}

type InvalidType struct {
	String struct{}
}

type TagPriority struct {
	Foo int
	Bar int `csv:"Foo"`
}

type embedded struct {
	Foo int `csv:"foo"`
	bar int `csv:"bar"`
}

type UnexportedEmbedded struct {
	embedded
}

type UnexportedEmbeddedPtr struct {
	*embedded
}

type A struct {
	B
	X int
}

type B struct {
	*A
	Y int
}

var Int = 10
var String = "string"
var PString = &String
var TypeISlice []TypeI

func pint(n int) *int                       { return &n }
func pint8(n int8) *int8                    { return &n }
func pint16(n int16) *int16                 { return &n }
func pint32(n int32) *int32                 { return &n }
func pint64(n int64) *int64                 { return &n }
func puint(n uint) *uint                    { return &n }
func puint8(n uint8) *uint8                 { return &n }
func puint16(n uint16) *uint16              { return &n }
func puint32(n uint32) *uint32              { return &n }
func puint64(n uint64) *uint64              { return &n }
func pfloat32(f float32) *float32           { return &f }
func pfloat64(f float64) *float64           { return &f }
func pstring(s string) *string              { return &s }
func pbool(b bool) *bool                    { return &b }
func pinterface(v interface{}) *interface{} { return &v }

func ppint(n int) **int       { p := pint(n); return &p }
func pppint(n int) ***int     { p := ppint(n); return &p }
func ppTypeI(v TypeI) **TypeI { p := &v; return &p }



var Error = errors.New("error")

var nilIface interface{}

var nilPtr *TypeF

var nilIfacePtr interface{} = nilPtr

type embeddedMap map[string]string

type Embedded14 Embedded3

func (e *Embedded14) MarshalCSV() ([]byte, error) {
	return json.Marshal(e)
}

type Embedded15 Embedded3

func (e *Embedded15) MarshalText() ([]byte, error) {
	return json.Marshal(Embedded3(*e))
}

type CSVMarshaler struct {
	Err error
}

func (m CSVMarshaler) MarshalCSV() ([]byte, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return []byte("csvmarshaler"), nil
}

type PtrRecCSVMarshaler int

func (m *PtrRecCSVMarshaler) MarshalCSV() ([]byte, error) {
	return []byte("ptrreccsvmarshaler"), nil
}

func (m *PtrRecCSVMarshaler) CSV() ([]byte, error) {
	return []byte("ptrreccsvmarshaler.CSV"), nil
}

type PtrRecTextMarshaler int

func (m *PtrRecTextMarshaler) MarshalText() ([]byte, error) {
	return []byte("ptrrectextmarshaler"), nil
}

type TextMarshaler struct {
	Err error
}

func (m TextMarshaler) MarshalText() ([]byte, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return []byte("textmarshaler"), nil
}

type CSVTextMarshaler struct {
	CSVMarshaler
	TextMarshaler
}

type Inline struct {
	J1      TypeJ  `csv:",inline"`
	J2      TypeJ  `csv:"prefix-,inline"`
	String  string `csv:"top-string"`
	String2 string `csv:"STR"`
}

type Inline2 struct {
	S string
	A Inline3 `csv:"A,inline"`
	B Inline3 `csv:",inline"`
}

type Inline3 struct {
	Inline4 `csv:",inline"`
}

type Inline4 struct {
	A string
}

type Inline5 struct {
	A Inline2 `csv:"A,inline"`
	B Inline2 `csv:",inline"`
}

type Inline6 struct {
	A Inline7 `csv:",inline"`
}

type Inline7 struct {
	A *Inline6 `csv:",inline"`
	X int
}

type Inline8 struct {
	F  *Inline4 `csv:"A,inline"`
	AA int
}

type TypeH struct {
	Int     int         `csv:"int,omitempty"`
	Int8    int8        `csv:"int8,omitempty"`
	Int16   int16       `csv:"int16,omitempty"`
	Int32   int32       `csv:"int32,omitempty"`
	Int64   int64       `csv:"int64,omitempty"`
	UInt    uint        `csv:"uint,omitempty"`
	Uint8   uint8       `csv:"uint8,omitempty"`
	Uint16  uint16      `csv:"uint16,omitempty"`
	Uint32  uint32      `csv:"uint32,omitempty"`
	Uint64  uint64      `csv:"uint64,omitempty"`
	Float32 float32     `csv:"float32,omitempty"`
	Float64 float64     `csv:"float64,omitempty"`
	String  string      `csv:"string,omitempty"`
	Bool    bool        `csv:"bool,omitempty"`
	V       interface{} `csv:"interface,omitempty"`
}

type TypeM struct {
	*TextMarshaler `csv:"text"`
}

type TypeJ struct {
	String string `csv:"STR" json:"string"`
	Int    string `csv:"int" json:"-"`
	Embedded16
	Float string `csv:"float"`
}

type Embedded16 struct {
	Bool  bool  `json:"bool"`
	Uint  uint  `csv:"-"`
	Uint8 uint8 `json:"-"`
}

var (
	textUnmarshaler = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
	csvUnmarshaler  = reflect.TypeOf((*csvutil.Unmarshaler)(nil)).Elem()
)