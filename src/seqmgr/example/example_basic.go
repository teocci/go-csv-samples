// Package bootseq
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package main

import (
	"context"
	"fmt"
	"github.com/teocci/go-csv-samples/src/seqmgr"
	"strings"
)

var words []string

func main() {
	// Let's use a boot sequence to construct a sentence!
	// For the shutdown sequence, we'll "deconstruct" it by removing each word.
	seq := seqmgr.New("Basic Example")
	seq.Register("welcome", add("Welcome"), rem)
	seq.Register("to", add("to"), rem).After("welcome")
	seq.Register("my", add("my"), rem).After("to")
	seq.Register("world", add("world!"), rem).After("my")

	agent, _ := seq.Agent()

	// Startup sequence.
	_ = agent.Up(context.Background(), process)

	fmt.Printf("\nwords: [%s]\n\n", strings.Join(words, ", "))

	// Shutdown sequence.
	_ = agent.Down(context.Background(), process)
	fmt.Println(strings.Join(words, " "))

	// Output:
	// welcome
	// to
	// my
	// world
	//
	// Welcome to my world!
	//
	// world
	// my
	// to
	// welcome
	//
}

func add(word string) func() error {
	return func() error {
		words = append(words, word)
		return nil
	}
}

func rem() error {
	words = words[:len(words)-1]
	return nil
}

func process(p seqmgr.Progress) {
	fmt.Println("Service:", p.Service)
}
