// Package seqmgr
// Created by Teocci.
// Author: teocci@yandex.com on 2021-Aug-28
package seqmgr

// Progress is the boot sequence feedback medium.
// Progress is communicated on channels returned by methods Up() and Down() and provides feedback on the current
// progress of the boot sequence. This includes the name of the Service that was last executed, along with an optional
// error if the Service Func failed. Err will be nil on success.
// Progress satisfies the error interface.
type Progress struct {
	Service string
	Err     error
}

// Error returns the error message for the receiver. Error returns an empty string if there is no error.
func (p Progress) Error() string {
	if p.Err == nil {
		return ""
	}

	return p.Err.Error()
}

// Verify that Progress satisfies the error interface.
var _ error = Progress{}