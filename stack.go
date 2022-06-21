package gots

import "runtime"

func GetCaller() (caller string) {
	caller = "(unknown)"
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return
	}

	if fpc := runtime.FuncForPC(pc); fpc != nil {
		caller = fpc.Name()
	}
	return
}
