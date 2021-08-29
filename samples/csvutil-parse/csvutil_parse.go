// Package csv_parse
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-24
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/jszwec/csvutil"
	"github.com/teocci/go-csv-samples/src/data"
)

const (
	geoPath   = "./tmp/GEOdata.csv"
	fccPath   = "./tmp/FCC.csv"
	destPath = "./tmp"
	rttPrefix = "01"
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
	// wait for all workers to finish up before exit
	defer waitTilEnd()()

	// open the first file
	base := loadData(geoPath)
	defer closeFile()(base)

	// open second file
	fcc := loadData(fccPath)
	defer closeFile()(fcc)

	// create a file writer
	rttFN := rttPrefix + "_RTTdata"
	fmt.Println("rttFN:", rttFN)
	rttPath := filepath.Join(destPath, rttFN+".csv")

	w := createFile(rttPath)
	defer closeFile()(w)

	// wrap the file readers with CSV readers
	bReader := csv.NewReader(base)
	bReader.Comma = ','
	bReader.Comment = '#'
	//fr := csv.NewReader(fcc)

	geoSlice := make([]*data.GEOData, 0)
	//fccSlice := make([]*FCC, 0)

	// wrap the out file writer with a CSV writer
	//cw := csv.NewWriter(w)
	//sessionDataSlice := make([]*FSessionData, 0)

	dec, err := csvutil.NewDecoder(bReader)
	dec.Tag = "csv"
	if err != nil {
		log.Fatal(err)
	}

	bHeader := dec.Header()
	fmt.Println(bHeader)

	numWps := 16
	jobs := make(chan *data.GEOData, numWps)
	res := make(chan *data.GEOData)

	worker := func(jobs <-chan *data.GEOData, results chan<- *data.GEOData) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}

				results <- job
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			geoData := new(data.GEOData)
			if err := dec.Decode(&geoData); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%#v\n", geoData)
			jobs <- geoData

			//rec, err := bReader.Read()
			//if err == io.EOF {
			//	break
			//}
			//if err != nil {
			//	fmt.Println("ERROR: ", err.Error())
			//	break
			//}
			//jobs <- rec
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		geoSlice = append(geoSlice, r)
		//fmt.Printf("%+v\n", r)
	}

	//for i, fccRec := range geoSlice {
	//	if i < 50 {
	//		fmt.Printf("%#v\n", fccRec)
	//	}
	//}

	fmt.Println("Count Concurrent ", len(geoSlice))

	//replaced by dec.Decode(&geoData)
	//for {
	//	rec, err := bReader.Read()
	//	if err != nil {
	//		if err == io.EOF {
	//			savePartitions()
	//			return
	//		}
	//		log.Fatal(err) // sorry for the panic
	//	}
	//	processCSV(rec, true)
	//}
}

func processCSV(rec []string, first bool) {
	l := len(rec)
	part := rec[l-1]

	if c, ok := workers[part]; ok {
		// send rec to workerClosure
		c <- rec
	} else {
		// if no workerClosure for the partition

		// make a chan
		nc := make(chan []string)
		workers[part] = nc

		// start workerClosure with this chan
		go workerClosure(nc, first)

		// send rec to workerClosure via chan
		nc <- rec
	}
}

func workerClosure(c chan []string, first bool) {
	// wg.Done signals to main workerClosure completion
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
			for _, p := range part {
				if first {
					fmt.Printf("%+v\n", p)
				}
			}
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
