// +build ffi

package intromain_test

import ( "testing" ; "math"
	//"bitbucket.org/thebridge0491/introgo/introutil"
	lib "bitbucket.org/thebridge0491/introgo/intromain"
)

func TestFact(t *testing.T) {
    type factFunc func(int) int64
    var ( n int = 5 ; res int64 = 120 )
    
    for _, f := range []factFunc{lib.FactI, lib.FactLp} {
        if res != f(n) { t.FailNow() }
    }
}

func TestExpt(t *testing.T) {
    type exptFunc func(float64, float64) float64
    var (param1 = make([]interface{}, 3) ; param2 = make([]interface{}, 3))
    for i, el := range []float64{2.0, 11.0, 20.0} { param1[i] = el }
    for i, el := range []float64{3.0, 6.0, 10.0} { param2[i] = el }
    
	for _, row := range CartesianProd(param1, param2) {
		var ( b = row[0].(float64) ; n = row[1].(float64) )
		var ans = math.Pow(b, n)
		for _, f := range []exptFunc{lib.ExptI, lib.ExptLp} {
		    if !InEpsilon(epsilon * ans, ans, f(b, n)) {
				t.FailNow()
		    }
        }
    }
}
