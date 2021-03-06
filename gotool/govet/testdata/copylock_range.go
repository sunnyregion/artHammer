// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the copylock checker's
// range statement analysis.

package testdata

import "sync"

func rangeMutex() {
	var mu sync.Mutex
	var i int

	var s []sync.Mutex
	for range s {
	}
	for i = range s {
	}
	for i := range s {
	}
	for i = range s {
	}
	for i := range s {
	}
	for _, mu = range s { // ERROR "range var mu copies lock: sync.Mutex"
	}
	for _, m := range s { // ERROR "range var m copies lock: sync.Mutex"
	}
	for i, mu = range s { // ERROR "range var mu copies lock: sync.Mutex"
	}
	for i, m := range s { // ERROR "range var m copies lock: sync.Mutex"
	}

	var a [3]sync.Mutex
	for _, m := range a { // ERROR "range var m copies lock: sync.Mutex"
	}

	var m map[sync.Mutex]sync.Mutex
	for k := range m { // ERROR "range var k copies lock: sync.Mutex"
	}
	for mu = range m { // ERROR "range var mu copies lock: sync.Mutex"
	}
	for k := range m { // ERROR "range var k copies lock: sync.Mutex"
	}
	for _, mu = range m { // ERROR "range var mu copies lock: sync.Mutex"
	}
	for _, v := range m { // ERROR "range var v copies lock: sync.Mutex"
	}

	var c chan sync.Mutex
	for range c {
	}
	for mu = range c { // ERROR "range var mu copies lock: sync.Mutex"
	}
	for v := range c { // ERROR "range var v copies lock: sync.Mutex"
	}

	// Test non-idents in range variables
	var t struct {
		i  int
		mu sync.Mutex
	}
	for t.i, t.mu = range s { // ERROR "range var t.mu copies lock: sync.Mutex"
	}
}
