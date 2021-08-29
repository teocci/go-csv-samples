// Package bootseq
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package seqmgr

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)


// Agent represents the execution of a sequence of Services.
// For any sequence, there will be two agents in play: one for
// the startup sequence, and another for the shutdown sequence.
// The only difference between these two are the order in
// which the sequence is executed.
// Each Agent keeps track of its progress and handles execution
// of sequence Services.
type Agent struct {
	name            string          // Name of boot sequence.
	progressFn      func(Progress)  // Progress reporting.
	orderedServices orderedServices // Map of Service priorities, with each  containing a slice of services.

	lock   sync.Mutex // Controls access to the fields below it.
	state  state      // Current state: up/down.
	isDone bool       // Did sequence execution complete?
}

// ServiceCount returns the number of services currently registered
// with the Agent.
func (a *Agent) ServiceCount() uint16 {
	return uint16(a.orderedServices.length())
}

// String returns a string representation of the registered Services
// ordered by priority. Service names are wrapped in parentheses, and
// separated by a colon when it might run concurrently with one or more
// other services, and a right-arrow when it will run before another
// service.
// Services that have the same priority are sorted alphabetically for
// reasons of reproducibility.
func (a *Agent) String() string {
	var sequence strings.Builder

	for i := uint16(1); i <= uint16(len(a.orderedServices)); i++ {
		names := make([]string, len(a.orderedServices[i]))
		for j, service := range a.orderedServices[i] {
			names[j] = service.name
		}
		if len(names) > 1 {
			sort.Strings(names)
		}
		sequence.WriteString(fmt.Sprintf("(%s) > ", strings.Join(names, " : ")))
	}

	ret := sequence.String()

	return ret[:len(ret)-3]
}

// Up runs the startup sequence.
// Up returns an error if the Agent's current state doesn't allow
// the sequence to start.
func (a *Agent) Up(ctx context.Context, progressFn func(Progress)) error {
	a.lock.Lock()
	if a.state != stateIdle {
		msg := inProgressErrorMessage
		if a.state == stateDown {
			msg = doneErrorMessage
		}
		a.lock.Unlock()

		return InvalidStateError(msg)
	}

	a.state = stateUp
	a.isDone = false
	a.progressFn = progressFn
	a.lock.Unlock()

	return a.exec(ctx)
}

// Down runs the shutdown sequence.
// Down returns an error if the Agent's current state doesn't allow
// the sequence to start.
func (a *Agent) Down(ctx context.Context, progressFn func(Progress)) error {
	a.lock.Lock()
	if a.state != stateUp || !a.isDone {
		msg := ""
		switch a.state {
		case stateIdle:
			msg = idleErrorMessage
		case stateUp:
			msg = upErrorMessage
		case stateDown:
			msg = inProgressErrorMessage
		}
		a.lock.Unlock()

		return InvalidStateError(msg)
	}

	a.state = stateDown
	a.isDone = false
	a.progressFn = progressFn
	a.lock.Unlock()

	return a.exec(ctx)
}

// report calls the provided progressFn with the given Progress struct.
func (a *Agent) report(progress Progress) {
	if a.progressFn == nil {
		return
	}
	a.progressFn(progress)
}

// exec runs through the sequence step by step and runs the relevant
// Service Func.
// The standard behaviour is to traverse the sequence in chronological
// order and run the "up" Func. If Agent.state == downState, the traversal
// is instead done in reverse order, and the "down" Func will run instead.
// After each Service has completed, progressFn is called (if provided)
// with a Progress struct.
func (a *Agent) exec(ctx context.Context) error {
	var err error
	defer func() {
		if err == nil {
			a.lock.Lock()
			a.isDone = true
			a.lock.Unlock()
		}
	}()

	var (
		current = 0
		step    = 1
		done    = make(chan error)
	)

	if a.state == stateDown {
		current = len(a.orderedServices) + 1
		step = -1
	}

	// Iterate over priority groups. Move in the direction from
	// priority [1..n] for startup sequences, and from priority [n..1]
	// for shutdown sequences. There is no guarantee regarding order
	// of execution within each priority group. It's possible
	// to interrupt the sequence between each priority group.
	for i := 0; i < len(a.orderedServices); i++ {
		current += step

		go a.execPriority(ctx, uint16(current), done)

		select {
		case <-ctx.Done():
			err = ctx.Err()
			<-done // Wait for execPriority to finish before stopping execution.
			//a.report(Progress{Service: "", Err: err})

			return err
		case err = <-done:
			if err != nil {
				return err
			}
			continue
		}
	}

	//a.report(Progress{Service: "", Err: err})

	return err
}

// execPriority executes all Services with the same priority/order.
// execPriority creates an errgroup for a single priority level
// in the Agent's orderedServices slice and runs them.
// execPriority returns an error if any one of the Services in
// the errgroup failed.
// execPriority is uninterruptible at this level.
func (a *Agent) execPriority(ctx context.Context, priority uint16, done chan<- error) {
	grp, _ := errgroup.WithContext(ctx)

	for _, service := range a.orderedServices[priority] {
		service := service
		grp.Go(func() error {
			err := service.byState(a.state)() // Execute the Service Func.
			a.report(Progress{Service: service.name, Err: err})

			return err
		})
	}

	done <- grp.Wait()
}
