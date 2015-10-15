# Golang Try Fail
Try or Fail functions
TryI passes attempt number

Try(max int, try func() bool) bool

TryI(max int, try func(int) bool) bool

TryFail(max int, try func() bool, fail func() bool) bool

TryIFail(max int, try func(int) bool, fail func() bool) bool

