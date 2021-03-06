package introforeignc

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L.
// #cgo pkg-config: intro_c-practice
/*
#include <stdlib.h>
#include "intro_c/classic.h"
*/
import ( "C" )

// FactI ...
func FactI(n int) int64 {
    //return int64(C.fact_i(C.long(n)))
    var res C.long = C.fact_i(C.long(n))
    return int64(res)
}

// FactLp ...
func FactLp(n int) int64 {
    return int64(C.fact_lp(C.long(n)))
}

// ExptI ...
func ExptI(b float64, n float64) float64 {
    return float64(C.expt_i(C.float(b), C.float(n)))
}

// ExptLp ...
func ExptLp(b float64, n float64) float64 {
    return float64(C.expt_lp(C.float(b), C.float(n)))
}
