// Package csvmgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package csvmgr

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func CloseFile() func(f *os.File) {
	return func(f *os.File) {
		fmt.Println("Defer: closing file.")
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func LoadDataBuff(fn string) []byte {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal("LoadDataBuff:", err)
	}
	defer CloseFile()(f)

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(f); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func OpenFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func CreateFile(fn string) *os.File {
	w, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}

	return w
}