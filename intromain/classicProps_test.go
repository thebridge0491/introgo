// +build ffi

package intromain_test

import ( "testing" ; "math" ; "github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
	"github.com/leanovate/gopter/gen"
	//"bitbucket.org/thebridge0491/introgo/introutil"
	lib "bitbucket.org/thebridge0491/introgo/intromain"
)

func TestPropFact(t *testing.T) {
	properties := gopter.NewProperties(nil)
	type factFunc func(int) int64
	properties.Property("fact n", prop.ForAll(
		func(num int) bool {
			var (ans int64 = 1 ; res = true)
			for n := int64(num); n > 1; n = n - 1 {
				ans = ans * n
			}
			for _, f := range []factFunc{lib.FactI, lib.FactLp} {
				res = res && ans == f(num)
			}
			return res
		},
		gen.IntRange(0, 18),
	))
	properties.TestingRun(t)
}

func TestPropExpt(t *testing.T) {
	properties := gopter.NewProperties(nil)
	type exptFunc func(float64, float64) float64
	properties.Property("expt b n", prop.ForAll(
		func(b, n float64) bool {
			var (ans = math.Pow(b, n) ; res = true)
			for _, f := range []exptFunc{lib.ExptI, lib.ExptLp} {
				res = res && InEpsilon(epsilon * ans, ans, f(b, n))
			}
			return res
		},
		gen.IntRange(2, 20).Map(func(x int) float64 {return float64(x)}), 
		gen.IntRange(2, 10).Map(func(x int) float64 {return float64(x)}),
	))
	properties.TestingRun(t)
}
