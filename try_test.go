// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package try

import (
	"fmt"
	"github.com/sabey/unittest"
	"testing"
)

func TestTry(t *testing.T) {
	fmt.Println("TestTry()")
	unittest.Equals(t, Try(-1, func() bool {
		return true
	}), false)
	unittest.Equals(t, Try(0, func() bool {
		return true
	}), false)
	// this function isn't thread safe
	i := 0
	f := func() bool {
		i++
		if i == 4 {
			return true
		}
		return false
	}
	i = 0
	i = 0
	unittest.Equals(t, Try(3, f), false)
	i = 0
	unittest.Equals(t, Try(4, f), true)
}
func TestTryI(t *testing.T) {
	fmt.Println("TestTryI()")
	unittest.Equals(t, TryI(-1, func(int) bool {
		return true
	}), false)
	unittest.Equals(t, TryI(0, func(int) bool {
		return true
	}), false)
	// this function isn't thread safe
	i := 0
	f := func(
		attempt int,
	) bool {
		i++
		if i == 4 {
			return true
		}
		return false
	}
	unittest.Equals(t, TryI(3, f), false)
	i = 0
	unittest.Equals(t, TryI(4, f), true)
}
func TestTryFail(t *testing.T) {
	fmt.Println("TestTryFail()")
	unittest.Equals(t, TryFail(-1, func() bool {
		return true
	}, nil), false)
	unittest.Equals(t, TryFail(0, func() bool {
		return true
	}, nil), false)
	unittest.Equals(t, TryFail(-1, nil, func() bool {
		return true
	}), true)
	unittest.Equals(t, TryFail(0, nil, func() bool {
		return true
	}), true)
	// this function isn't thread safe
	i := 0
	f := func() bool {
		i++
		if i == 4 {
			return true
		}
		return false
	}
	unittest.Equals(t, TryFail(3, f, nil), false)
	i = 0
	unittest.Equals(t, TryFail(3, f, func() bool {
		return false
	}), false)
	i = 0
	unittest.Equals(t, TryFail(3, f, func() bool {
		return true
	}), true)
	i = 0
	unittest.Equals(t, TryFail(4, f, nil), true)
	i = 0
	unittest.Equals(t, TryFail(4, f, func() bool {
		return false
	}), true)
}
func TestTryIFail(t *testing.T) {
	fmt.Println("TestTryIFail()")
	unittest.Equals(t, TryIFail(-1, func(int) bool {
		return true
	}, nil), false)
	unittest.Equals(t, TryIFail(0, func(int) bool {
		return true
	}, nil), false)
	unittest.Equals(t, TryIFail(-1, nil, func() bool {
		return true
	}), true)
	unittest.Equals(t, TryIFail(0, nil, func() bool {
		return true
	}), true)
	// this function isn't thread safe
	i := 0
	f := func(
		attempt int,
	) bool {
		i++
		if i == 4 {
			return true
		}
		return false
	}
	unittest.Equals(t, TryIFail(3, f, nil), false)
	i = 0
	unittest.Equals(t, TryIFail(3, f, func() bool {
		return false
	}), false)
	i = 0
	unittest.Equals(t, TryIFail(3, f, func() bool {
		return true
	}), true)
	i = 0
	unittest.Equals(t, TryIFail(4, f, nil), true)
	i = 0
	unittest.Equals(t, TryIFail(4, f, func() bool {
		return false
	}), true)
}
