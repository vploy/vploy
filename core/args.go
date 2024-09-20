package core

import "os"

type xArgs struct {
	//
}

func (x *xArgs) Take(i int) string {
	value := os.Args[i]
	os.Args = append(os.Args[:i], os.Args[i+1:]...)
	return value
}

func (x *xArgs) TakeFirst() string {
	return x.Take(1)
}

func (x *xArgs) TakeLast() string {
	return x.Take(x.Len() - 1)
}

func (x *xArgs) Len() int {
	return len(os.Args)
}

var Args xArgs
