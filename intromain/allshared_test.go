package intromain_test

import ( //"testing"
	"math"
	//"bitbucket.org/thebridge0491/introgo/introutil"
)

var epsilon = 0.001

func InEpsilon(tolerance float64, a float64, b float64) bool {
	var delta float64 = math.Abs(tolerance)
	//return (a - delta) <= b && (a + delta) >= b
	return !((a + delta) < b) && !((b + delta) < a)
}

func CartesianProd(arr1 []interface{}, arr2 []interface{}) [][2]interface{} {
	var idxX int = 0
	var arrProd = make([][2]interface{}, len(arr1) * len(arr2))
	
	for idx1, el1 := range arr1 {
		for idx2, el2 := range arr2 {
			idxX = idx2 + (idx1 * len(arr2))
			arrProd[idxX][0] = el1
			arrProd[idxX][1] = el2
		}
	}
	return arrProd
}
