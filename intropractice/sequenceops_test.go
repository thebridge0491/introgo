package intropractice_test

import ( "testing" 
	util "bitbucket.org/thebridge0491/introgo/introutil"
	lib "bitbucket.org/thebridge0491/introgo/intropractice"
)

var ints, intsRev = []int{0, 1, 2, 3, 4}, []int{4, 3, 2, 1, 0}

func TestIndexFind(t *testing.T) {
    type findIdxFunc func(interface{}, []interface{}) int
    var el int = 3
    
    for _, f := range []findIdxFunc{lib.FindIndexLp} {
        if 3 != f(el, util.IfcArrFromInts(ints)) { t.FailNow() }
        if 1 != f(el, util.IfcArrFromInts(intsRev)) { t.FailNow() }
    }
}

func TestReverse(t *testing.T) {
    type reverseFunc func([]interface{}) ()
    
    for _, f := range []reverseFunc{lib.ReverseLp} {
        var arrTmp = lib.CopyOf(util.IfcArrFromInts(ints))
        f(arrTmp)
        
        for idx, el := range util.IfcArrFromInts(intsRev) {
            if el != arrTmp[idx] { t.FailNow() }
        }
    }
}
