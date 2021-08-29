// Package bootseq
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package seqmgr

const emptyServiceName = ""

// Service contains the functions required in order to execute a single Service Func
// in a sequence, the up() and down() functions, respectively.
type Service struct {
	name     string
	priority uint16
	up, down Func
	after    string
}

// After sets the receiver Service to be executed after the one defined by the given name.
func (s *Service) After(name string) {
	s.after = name
}

// byState returns the service function that matches the provided state.
// It panics if the state is unknown.
func (s *Service) byState(ph state) Func {
	switch ph {
	case stateUp:
		return s.up
	case stateDown:
		return s.down
	default:
		panic(panicUnknownState)
	}
}