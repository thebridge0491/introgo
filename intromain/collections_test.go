package intromain_test

import ( "testing" ; "fmt" ; "math" //; "sort"
	"container/list" ; "container/ring" ; "container/heap"
	util "bitbucket.org/thebridge0491/introgo/introutil"
)

type Tuple2 struct {
	t1, t2 interface{}
}

func (tup Tuple2) String() string {
	return fmt.Sprintf("(%v, %v)", tup.t1, tup.t2) }

func TestLists(t *testing.T) {
    var (numArr = []int{16, 2, 77, 29} ; nines = []int{9, 9, 9, 9})
    lst1 := list.New()
    
    if 0 != lst1.Len() && nil != lst1.Front() { t.Error("isEmpty") }
    for _, val := range numArr { lst1.PushBack(val) }
	
	if len(numArr) != lst1.Len() { t.Error("length") }
	if numArr[0] != lst1.Front().Value { t.Error("first") }
	
	for _, val := range nines { lst1.PushBack(val) }
	if (len(numArr) + len(nines)) != lst1.Len() { t.Error("append") }
	
    for el := lst1.Front(); nil != el; el = el.Next() {
		if 29 == el.Value {lst1.Remove(el)} }
	resTxt := util.Mkstring(util.ListToSlice(lst1))
	if "[16, 2, 77, 9, 9, 9, 9]" != resTxt { 
		t.Errorf("toString: %s", resTxt) }
}

func TestRings(t *testing.T) {
    var (numArr = []int{16, 2, 77, 29} ; nines = []int{9, 9, 9, 9})
    var ring1 = ring.New(len(numArr))
    
    if nil != ring1.Value { t.Error("isEmpty") }
    el := ring1
    for _, val := range numArr { el.Value = val ; el = el.Next() }
	
	if len(numArr) != ring1.Len() { t.Error("length") }
	if numArr[0] != ring1.Value { t.Error("first") }
	
	var ring2 = ring.New(len(nines))
	el2 := ring2
	for _, val := range nines { el2.Value = val ; el2 = el2.Next() }
	ring1 = ring1.Prev().Link(ring2)
	if (len(numArr) + len(nines)) != ring1.Len() { t.Error("append") }
	
	resTxt := util.Mkstring(util.RingToSlice(ring1))
	if "[16, 2, 77, 29, 9, 9, 9, 9]" != resTxt { 
		t.Errorf("toString: %s", resTxt) }
}

func TestMaps(t *testing.T) {
    var ltrArr []string = []string{"a", "e", "k", "p", "u", "k", "a"}
    map1 := make(map[string]interface{})
    
    if 0 != len(map1) { t.Error("isEmpty") }
    for i, e := range ltrArr { map1["ltr " + fmt.Sprint(i % 5)] = e }
	
	if nil == map1["ltr 2"] { t.Error("contains") }
	if "k" != map1["ltr 2"].(string) { t.Error("get") }
	delete(map1, "ltr 2")
	if nil != map1["ltr 2"] { t.Error("remove") }
	map1["ltr 2"] = "Z"
	if nil == map1["ltr 2"] { t.Error("put") }
	if len(map1) != 5 { t.Error("length") }
	
	tupArr := []Tuple2{Tuple2{"ltr 0", "k"}, Tuple2{"ltr 0", "k"}, 
		Tuple2{"ltr 1", "a"}, Tuple2{"ltr 1", "a"}, Tuple2{"ltr 2", "Z"}, 
		Tuple2{"ltr 3", "p"}, Tuple2{"ltr 4", "u"}}
	for _, tup := range tupArr { if nil == map1[fmt.Sprint(tup.t1)] { t.Error("equal") } }
}

func TestSortableMaps(t *testing.T) {
    var ltrArr []string = []string{"a", "e", "k", "p", "u", "k", "a"}
    map1 := new(util.SortableMap)
    map1.M = make(map[string]interface{})
    
    if 0 != len(map1.M) { t.Error("isEmpty") }
    for i, e := range ltrArr {
		//key1 := "ltr " + fmt.Sprint(i % 5)
		//map1.Keys = append(map1.Keys, key1) ; map1.M[key1] = e
		map1.Append("ltr " + fmt.Sprint(i % 5), e)
	}
	
	if nil == map1.M["ltr 2"] { t.Error("contains") }
	if "k" != map1.M["ltr 2"].(string) { t.Error("get") }
	delete(map1.M, "ltr 2")
	if nil != map1.M["ltr 2"] { t.Error("remove") }
	map1.M["ltr 2"] = "Z"
	if nil == map1.M["ltr 2"] { t.Error("put") }
	resTxt := util.MkstringInit("{", ", ", "}", util.SortableMapToSlice(true, map1))
	if "{(ltr 0, k), (ltr 0, k), (ltr 1, a), (ltr 1, a), (ltr 2, Z), (ltr 3, p), (ltr 4, u)}" != 
		resTxt {
			t.Errorf("toString: %s", resTxt) }
}

func TestHeaps(t *testing.T) {
    var floatArr []float32 = []float32{27.5, 0.1, 78.5, 52.3}
    var heap1 *util.Heap = &util.Heap{}
    
    heap.Init(heap1)
    
    if 0 != len(*heap1) { t.Error("isEmpty") }
    for i, val := range floatArr {
		heap.Push(heap1, &util.Item{Value: val, Prio: int(val), Idx: i}) }
    
    if len(floatArr) != len(*heap1) { t.Error("length") }
    
    var res = heap.Pop(heap1).(*util.Item) ; heap.Push(heap1, res)
    if !util.InEpsilon(epsilon * 0.1, 0.1, float64(res.Value.(float32))) {
		t.Errorf("peek: %.2f", res.Value.(float32)) }
	
	if !util.InEpsilon(epsilon * 0.1, 0.1,
		float64(heap.Pop(heap1).(*util.Item).Value.(float32))) { t.Error("pop") }
    
    heap.Push(heap1, &util.Item{Value: float32(-0.5), Prio: int(math.Round(-0.5)),
		Idx: len(*heap1)})
	res = heap.Pop(heap1).(*util.Item) ; heap.Push(heap1, res)
    if !util.InEpsilon(epsilon * -0.5, -0.5, float64(res.Value.(float32))) {
		t.Errorf("push: %.2f", res.Value.(float32)) }
	
	resTxt := util.Mkstring(util.HeapToSlice(heap1))
    if "[-0.5, 27.5, 78.5, 52.3]" != resTxt { t.Errorf("toString: %s", resTxt) }
}
