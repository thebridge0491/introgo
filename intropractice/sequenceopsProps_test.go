package intropractice_test

import ( "testing" ; "math/rand" ; "time" ; "github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
	"github.com/leanovate/gopter/gen"
	util "bitbucket.org/thebridge0491/introgo/introutil"
	lib "bitbucket.org/thebridge0491/introgo/intropractice"
)

func TestPropIndexFind(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSize = 1 ; parameters.MaxSize = 20
	properties := gopter.NewProperties(parameters)
	type findIndexFunc func(interface{}, []interface{}) int
	rand.Seed(time.Now().Unix())
	
    properties.Property("findIndex", prop.ForAll(
		func(iarr []interface{}) bool {
			var (ndx = rand.Intn(len(iarr)) ; res = true)
			for _, f := range []findIndexFunc{lib.FindIndexLp} {
				res = res && -1 != f(iarr[ndx], iarr) && 
					ndx >= f(iarr[ndx], iarr)
			}
			return res
		},
		gen.SliceOf(gen.IntRange(-150, 150)).Map(
			func(arr []int) []interface{} {return util.IfcArrFromInts(arr)}),
	))
	properties.TestingRun(t)
}

func TestPropReverse(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSize = 0 ; parameters.MaxSize = 20
	properties := gopter.NewProperties(parameters)
	type reverseFunc func([]interface{}) ()
	
    properties.Property("reverse", prop.ForAll(
		func(iarr []interface{}) bool {
			var (newIarr = make([]interface{}, len(iarr)) ; res = true)
			copy(newIarr, iarr) ; var maxIdx = len(newIarr) - 1
			for i := 0; (maxIdx >> 1) >= i; i = i + 1 {
				lib.SwapItems(i, maxIdx - i, newIarr)
			}
			for _, f := range []reverseFunc{lib.ReverseLp} {
				//var tmpIarr = make([]interface{}, len(iarr))
				//copy(tmpIarr, iarr)
				var tmpIarr = lib.CopyOf(iarr)
				f(tmpIarr)
				for i, _ := range newIarr {
					res = res && tmpIarr[i] == newIarr[i]
				}
			}
			return res
		},
		gen.SliceOf(gen.IntRange(-150, 150)).Map(
			func(arr []int) []interface{} {return util.IfcArrFromInts(arr)}),
	))
	properties.TestingRun(t)
}
