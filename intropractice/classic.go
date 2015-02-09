package intropractice

import ( log "github.com/alecthomas/log4go" )

// FactLp ...
func FactLp(n int) int64 {
    var acc int64 = 1
    log.Debug("FactLp()")
    for i := n; i > 1; i-- {
    	acc *= int64(i)
    }
    return acc
}

func factIter(n int, acc int64) int64 {
	if n > 1 {
		return factIter(n - 1, acc * int64(n))
	} else { return acc }
}

// FactI ...
func FactI(n int) int64 { return factIter(n, int64(1))
}

// ExptLp ...
func ExptLp(b float64, n float64) float64 {
    var acc float64 = 1.0
    for i := n; i > float64(0.0); i-- {
    	acc *= b
    }
    return acc
}

func exptIter(b float64, n float64, acc float64) float64 {
	if n > float64(0.0) {
		return exptIter(b, n - 1, acc * b)
	} else { return acc }
}

// ExptI ...
func ExptI(b float64, n float64) float64 {
    return exptIter(b, n, float64(1.0))
}

