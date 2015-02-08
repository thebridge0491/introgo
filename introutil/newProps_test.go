package introutil_test

import ( "testing" ; "sort" ; "github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
	"github.com/leanovate/gopter/gen"
	//"bitbucket.org/thebridge0491/introgo/introutil"
)

func TestPropCommutAdd(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("(ints) addition is commutative", prop.ForAll(
		func(a int, b int) bool {
			return a + b == b + a
		},
		gen.IntRange(-150, 150), gen.IntRange(-150, 150),
	))
	properties.Property("(floats) addition is commutative", prop.ForAll(
		func(a float64, b float64) bool {
			return InEpsilon(epsilon * (a + b), a + b, b + a)
		},
		gen.Float64Range(-150.0, 150.0), gen.Float64Range(-150.0, 150.0),
	))
	properties.TestingRun(t)
}

func TestPropAssocAdd(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("(ints) addition is associative", prop.ForAll(
		func(a int, b int, c int) bool {
			return (a + b) + c == a + (b + c)
		},
		gen.IntRange(-150, 150), gen.IntRange(-150, 150),
		gen.IntRange(-150, 150),
	))
	properties.Property("(floats) addition is associative", prop.ForAll(
		func(a float64, b float64, c float64) bool {
			return InEpsilon(epsilon * ((a + b) + c), 
				(a + b) + c, a + (b + c))
		},
		gen.Float64Range(-150.0, 150.0), gen.Float64Range(-150.0, 150.0), gen.Float64Range(-150.0, 150.0),
	))
	properties.TestingRun(t)
}

func TestPropReverse(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSize = 1 ; parameters.MaxSize = 20
	properties := gopter.NewProperties(parameters)
	properties.Property("slice is equal reverse(reverse slice)", prop.ForAll(
		func(arr []int) bool {
			var (newArr = make([]int, len(arr)) ; res = true)
			copy(newArr, arr) ; var maxIdx = len(newArr) - 1
			for i := 0; (maxIdx >> 1) > i; i = i + 1 {
				swap := newArr[i] ; newArr[i] = newArr[maxIdx - i]
				newArr[maxIdx - i] = swap
			}
			for i := 0; (maxIdx >> 1) > i; i = i + 1 {
				swap := newArr[i] ; newArr[i] = newArr[maxIdx - i]
				newArr[maxIdx - i] = swap
			}
			for i, _ := range newArr {
				res = res && newArr[i] == arr[i]
			}
			return res
		},
		gen.SliceOf(gen.IntRange(-150, 150)),
	))
	properties.Property("slice is equal reverse slice", prop.ForAll(
		func(arr []float64) bool {
			var (newArr = make([]float64, len(arr)) ; res = true)
			copy(newArr, arr) ; var maxIdx = len(newArr) - 1
			for i := 0; (maxIdx >> 1) > i; i = i + 1 {
				swap := newArr[i] ; newArr[i] = newArr[maxIdx - i]
				newArr[maxIdx - i] = swap
			}
			for i, _ := range newArr {
				res = res && InEpsilon(epsilon * arr[i], arr[i], newArr[i])
			}
			return res
		},
		gen.SliceOf(gen.Float64Range(-150.0, 150.0)),
	))
	properties.TestingRun(t)
}

func TestPropSort(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSize = 1 ; parameters.MaxSize = 20
	properties := gopter.NewProperties(parameters)
	properties.Property("sort is equal sort(reverse slice)", prop.ForAll(
		func(arr []int) bool {
			var (newArr = make([]int, len(arr)) ; res = true)
			copy(newArr, arr) ; var maxIdx = len(newArr) - 1
			for i := 0; (maxIdx >> 1) > i; i = i + 1 {
				swap := newArr[i] ; newArr[i] = newArr[maxIdx - i]
				newArr[maxIdx - i] = swap
			}
			sort.Sort(sort.IntSlice(arr))
			sort.Sort(sort.IntSlice(newArr))
			for i, _ := range newArr {
				res = res && newArr[i] == arr[i]
			}
			return res
		},
		gen.SliceOf(gen.IntRange(-150, 150)),
	))
	properties.Property("minimum is equal sort(slice)[0]", prop.ForAll(
		func(arr []float64) bool {
			var newArr = make([]float64, len(arr))
			copy(newArr, arr) ; var minEl = arr[0]
			sort.Sort(sort.Float64Slice(newArr))
			for i := 0; len(arr) > i; i = i + 1 {
				if arr[i] < minEl { minEl = arr[i] }
			}
			return InEpsilon(epsilon * minEl, minEl, newArr[0])
		},
		gen.SliceOf(gen.Float64Range(-150.0, 150.0)),
	))
	properties.Property("minimum is equal sort(slice0 ++ slice1)[0]", prop.ForAll(
		func(arr0, arr1 []int) bool {
			var newArr0 = make([]int, len(arr0) + len(arr1))
			var newArr1 = make([]int, len(arr0) + len(arr1))
			copy(newArr0, arr0) ; newArr0 = append(newArr0, arr1...)
			copy(newArr1, arr0) ; newArr1 = append(newArr1, arr1...)
			var minEl = newArr1[0]
			sort.Sort(sort.IntSlice(newArr0))
			for i := 0; len(newArr1) > i; i = i + 1 {
				if newArr1[i] < minEl { minEl = newArr1[i] }
			}
			return minEl == newArr0[0]
		},
		gen.SliceOf(gen.IntRange(-150, 150)),
		gen.SliceOf(gen.IntRange(-150, 150)),
	))
	properties.TestingRun(t)
}
