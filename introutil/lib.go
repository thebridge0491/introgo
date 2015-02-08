package introutil

import ( "fmt" ; "math" ; "github.com/unknwon/goconfig"
	"sort" ; "container/list" ; "container/ring" //; "container/heap"
)

// MkstringInit ...
func MkstringInit(beg string, sep string, stop string,
		slice0 []interface{}) string {
	var (result string = "" ; prefix string)
	
	for _, el := range slice0 {
		if "" == result { prefix = "" } else { prefix = sep }
		result += prefix + fmt.Sprint(el)
	}
	return beg + result + stop
}

// Mkstring ...
func Mkstring(slice0 []interface{}) string {
	return MkstringInit("[", ", ", "]", slice0)
}

// IniCfgToStr ...
func IniCfgToStr(cfg *goconfig.ConfigFile) string {
    var sliceX = make([]interface{}, 0)
	
	sectList := cfg.GetSectionList()
    for _, sect := range sectList {
    	keyList := cfg.GetKeyList(sect)
    	
    	for _, key := range keyList {
    		val, err := cfg.GetValue(sect, key)
    		
    		if nil == err {
    			sliceX = append(sliceX, fmt.Sprintf("%s:%s => %s", sect, 
					key, val))
    		}
    	}
    }
    return MkstringInit("{", ", ", "}", sliceX)
}

// InEpsilon ...
func InEpsilon(tolerance float64, a float64, b float64) bool {
	var delta float64 = math.Abs(tolerance)
	//return (a - delta) <= b && (a + delta) >= b
	return !((a + delta) < b) && !((b + delta) < a)
}

// CartesianProd ...
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

// IfcArrFromInts ...
func IfcArrFromInts(arr []int) []interface{} {
	var iarr = make([]interface{}, len(arr))
	for i, el := range arr { iarr[i] = el }
	return iarr
}

// IfcArrFromFloats ...
func IfcArrFromFloats(arr []float64) []interface{} {
	var iarr = make([]interface{}, len(arr))
	for i, el := range arr { iarr[i] = el }
	return iarr
}

type Tuple2 struct {
	t1, t2 interface{}
}

func (tup Tuple2) String() string {
	return fmt.Sprintf("(%v, %v)", tup.t1, tup.t2) }

// MapToSlice ...
func MapToSlice(mapX map[string]interface{}) []interface{} {
	var sliceX = make([]interface{}, 0)
	for k, v := range mapX {
		tup := Tuple2{k, v}
		sliceX = append(sliceX, tup)
	}
	return sliceX
}

type SortableMap struct {
	Keys []string ; M map[string]interface{} }

func (sm *SortableMap) Append(k string, v interface{}) {
	sm.Keys, sm.M[k] = append(sm.Keys, k), v }

func (sm SortableMap) Len() int { // Len|Less|Swap needed for sort.Sort(sm)
    return len(sm.Keys) }
 
func (sm SortableMap) Less(i, j int) bool {
    return sm.Keys[i] < sm.Keys[j] }
 
func (sm SortableMap) Swap(i, j int) {
    sm.Keys[i], sm.Keys[j] = sm.Keys[j], sm.Keys[i] }

// SortableMapToSlice ...
func SortableMapToSlice(isSorted bool, mapX *SortableMap) []interface{} {
	var sliceX = make([]interface{}, 0)
	if isSorted {
		sort.Sort(mapX) // or: sort.Strings(mapX.Keys)
	}
	for _, k := range mapX.Keys {
		tup := Tuple2{k, mapX.M[k]}
		sliceX = append(sliceX, tup)
	}
	return sliceX
}

// ListToSlice ...
func ListToSlice(xs *list.List) []interface{} {
	var sliceX = make([]interface{}, 0)
	for el := xs.Front(); nil != el; el = el.Next() {
		sliceX = append(sliceX, el.Value)
	}
	return sliceX
}

// RingToSlice ...
func RingToSlice(ringX *ring.Ring) []interface{} {
	var sliceX = make([]interface{}, 0)
	for i, rg := 0, ringX; ringX.Len() > i; i++ {
		sliceX = append(sliceX, rg.Value)
		rg = rg.Next()
	}
	return sliceX
}

type Item struct {
	Value interface{} ; Prio int ; Idx int }

type Heap []*Item

func (h Heap) Len() int {
	return len(h) }

func (h Heap) Less(i, j int) bool {
	return h[i].Prio < h[j].Prio }

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Idx, h[j].Idx = i, j }

func (h *Heap) Push(x interface{}) {
	item := x.(*Item)
	item.Idx = len(*h)
	*h = append(*h, item) }

func (h *Heap) Pop() interface{} {
	old, n := *h, len(*h)
	item := old[n - 1]
	item.Idx = -1 ; *h = old[0 : n - 1]
	return item }

// HeapToSlice ...
func HeapToSlice(heapX *Heap) []interface{} {
	var sliceX = make([]interface{}, 0)
	for _, el := range *heapX {
		sliceX = append(sliceX, el.Value)
	}
	return sliceX
}
