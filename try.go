// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package try

func Try(
	max int,
	try func() bool,
) bool {
	if try == nil {
		return false
	}
	for i := 1; i <= max; i++ {
		if try() {
			return true
		}
	}
	return false
}
func TryI(
	max int,
	try func(int) bool,
) bool {
	if try == nil {
		return false
	}
	for i := 1; i <= max; i++ {
		if try(i) {
			return true
		}
	}
	return false
}
func TryFail(
	max int,
	try func() bool,
	fail func() bool,
) bool {
	if Try(max, try) {
		return true
	}
	if fail == nil {
		return false
	}
	return fail()
}
func TryIFail(
	max int,
	try func(int) bool,
	fail func() bool,
) bool {
	if TryI(max, try) {
		return true
	}
	if fail == nil {
		return false
	}
	return fail()
}
