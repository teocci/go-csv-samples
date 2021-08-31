// Package csv_parse
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-18
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/teocci/go-csv-samples/src/data"
)

var (
	// list of channels to communicate with workers
	// Those will be accessed synchronously no mutex required
	workers = make(map[string]chan []string)

	// wg is to make sure all workers done before exiting main
	wg = sync.WaitGroup{}

	// mu used only for sequential printing, not relevant for program logic
	mu = sync.Mutex{}
)

func main() {
	initProcess()
}

func initProcess() {
	// wait for all workers to finish up before exit
	defer waitTilEnd()()

	// open the first file
	base := loadData(data.GEODatPath)
	defer closeFile()(base)

	r := csv.NewReader(base)
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				savePartitions()
				return
			}
			log.Fatal(err) // sorry for the panic
		}

		process(rec)
	}
}

func process(rec []string) {
	part := rec[0]

	if c, ok := workers[part]; ok {
		// send rec to worker
		c <- rec
	} else {
		// if no worker for the partition

		// make a chan
		nc := make(chan []string)
		workers[part] = nc

		// start worker with this chan
		go worker(nc)

		// send rec to worker via chan
		nc <- rec
	}
}

func worker(c chan []string) {
	// wg.Done signals to main worker completion
	wg.Add(1)
	defer wg.Done()

	var part [][]string
	for {
		// wait for a rec or close(chan)
		rec, ok := <-c
		if ok {
			// save the rec
			// instead of accumulation in memory
			// this can be saved to file directly
			part = append(part, rec)
		} else {
			// channel closed on EOF

			// dump partition
			// locks ensures sequential printing
			// not a required for independent files
			mu.Lock()
				fmt.Printf("%+v\n", part)
			mu.Unlock()

			return
		}
	}
}

// simply signals to workers to stop
func savePartitions() {
	for _, c := range workers {
		// signal to all workers to exit
		close(c)
	}
}

func waitTilEnd() func() {
	return func() {
		wg.Wait()
		fmt.Println("File processed.")
	}
}

func closeFile() func(f *os.File) {
	return func(f *os.File) {
		fmt.Println("Defer: closing file.")
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadData(f string) *os.File {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func createFile(f string) *os.File {
	w, err := os.Create(f)
	if err != nil {
		log.Fatal(err)
	}

	return w
}
