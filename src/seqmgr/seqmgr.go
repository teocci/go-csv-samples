// Package bootseq
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package seqmgr

import (
	"fmt"
	"sync"
)

// state represents a Manager's state. It's either:
// 1. doing nothing (stateIdle),
// 2. in the startup sequence (stateUp),
// 3. in the shutdown sequence (stateDown).
type state uint8

const (
	stateIdle state = iota
	stateUp
	stateDown
)

// Func is the type used for any function that can be executed as a
// service in a boot sequence. Any function that you wish to register
// and execute as a service must satisfy this type.
type Func func() error

// unorderedServices represents a collection of Services before
// they've been ordered.
type unorderedServices map[string]*Service

// orderedServices represents a collection of Services after
// they've been ordered.
type orderedServices map[uint16][]Service

// Manager provides registration and storage of boot sequence Services.
// Manager can instantiate an Agent, which is responsible for running
// the actual startup and shutdown sequences.
type Manager struct {
	name string

	lock     sync.Mutex // Protects field services.
	services unorderedServices
}

// New returns a new and uninitialised boot sequence Manager.
func New(name string) *Manager {
	services := make(map[string]*Service)
	mgr := Manager{lock: sync.Mutex{}, name: name, services: services}

	return &mgr
}

// Register registers a single named Service to the boot sequence,
// with the given "up" and "down" functions. If a Service with
// the given name already exists, the provided up- and down
// functions replace those already registered. Add returns a pointer
// to the added Service, that you can call After() on, in order
// to influence order of execution.
func (m *Manager) Register(name string, up, down Func) *Service {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.services) == 65535 {
		panic(panicServiceLimit)
	}

	ref := &Service{name, 0, up, down, emptyServiceName}
	m.services[name] = ref

	return ref
}

// ServiceCount returns the number of services currently registered
// with the Manager.
func (m *Manager) ServiceCount() uint16 {
	m.lock.Lock()
	defer m.lock.Unlock()

	return uint16(len(m.services))
}

// ServiceNames returns the name of each registered service, in no
// particular order.
func (m *Manager) ServiceNames() []string {
	m.lock.Lock()
	defer m.lock.Unlock()

	ns := make([]string, 0, len(m.services))

	for name := range m.services {
		ns = append(ns, name)
	}

	return ns
}

// Agent orders the registered services by priority and returns an Agent
// for controlling the startup and shutdown sequences. Agent returns
// an error if any of the registered Services refer to other Services
// that are not registered.
func (m *Manager) Agent() (agent *Agent, err error) {
	m.lock.Lock()
	if len(m.services) == 0 {
		err = EmptySequenceError(m.name)
		return
	}
	m.lock.Unlock()

	if err = m.Validate(); err != nil {
		return
	}
	agent = &Agent{}
	agent.name = m.name
	agent.orderedServices = m.services.order()
	return
}

// Validate cycles through each registered service and checks if they
// refer to other service names that don't exist, or if they refer
// to themselves. Validate returns an error if this is the case,
// or nil otherwise.
func (m *Manager) Validate() error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.services) == 0 {
		return EmptySequenceError(m.name)
	}

	for name, svc := range m.services {
		if svc.up == nil || svc.down == nil {
			return NilFuncError(svc.name)
		}
		if svc.after == emptyServiceName {
			continue
		}
		if svc.after == name {
			return SelfReferenceError(svc.after)
		}
		prev, ok := m.services[svc.after]
		if ok {
			if prev.after == svc.name {
				return CyclicReferenceError(svc.name)
			}
		} else {
			return UnregisteredServiceError(svc.after)
		}
	}

	return nil
}

// setPriority looks up the Service with the given name and attempts
// to set its priority.
// If the Service depends on another, setPriority recursively follows
// the chain of Services in order to determine priorities for the
// entire chain. setPriority returns the priority that has been
// resolved for the given Service.
func (u unorderedServices) setPriority(name string) uint16 {
	if name == emptyServiceName {
		return 0
	}
	service, ok := u[name]
	if !ok {
		panic(fmt.Sprintf("missing Service: %q, was Manager.Validate called?", name))
	}
	if service.priority > 0 {
		return service.priority
	}
	if service.after == emptyServiceName {
		service.priority = 1

		return 1
	}
	service.priority = u.setPriority(service.after) + 1

	return service.priority
}

// NoOp (no operation) is a convenience function you can use in place
// of a Service Func for when you want a function that does nothing.
func NoOp() error {
	return nil
}

// order will order each Service in unorderedServices by priority.
// order returns the same Services in order of reference.
// The algorithm is:
// 1. Services that don't come after another, receive order 1.
// 2. Services that come immediately after another, receive an order
// that is one higher than the other.
// 3. If a service refers to another which is unordered, a depth-first
// approach is taken to resolve the orders of each one.
// order assumes that each referenced service exists.
func (u unorderedServices) order() orderedServices {
	ordered := make(orderedServices, len(u))
	if len(u) == 0 {
		return ordered
	}

	var service *Service
	var priority uint16

	for name := range u {
		priority = u.setPriority(name)
		service = u[name]
		ordered[priority] = append(ordered[priority], *service)
	}

	return ordered
}

// length returns the total number of registered Services.
func (o orderedServices) length() int {
	length := 0
	for _, services := range o {
		length += len(services)
	}

	return length
}
