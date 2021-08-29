// Package csv_pool
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-29
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/teocci/go-csv-samples/src/data"
)

func main() {
	initProcess()
}

// with Worker pools
func initProcess() {
	// open the first file
	base := loadData(data.GEOPath)
	defer closeFile()(base)

	csvReader := csv.NewReader(base)
	geos := make([]*data.GEOData, 0)

	numWps := runtime.NumCPU()
	jobs := make(chan []string, numWps)
	res := make(chan *data.GEOData)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *data.GEOData) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}

				results <- data.ParseGEODataStruct(job)
			}
		}
	}

	// init workers
	for w:=0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			rStr, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		geos = append(geos, r)
	}

	for i, rec := range geos {
		if i < 50 {
			fmt.Printf("%#v\n", rec)
		}
	}

	fmt.Println("Count Concurrent ", len(geos))
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