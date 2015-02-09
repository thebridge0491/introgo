package intropractice

import ( log "github.com/alecthomas/log4go" 
)

// SwapItems ...
func SwapItems(a, b int, arr []interface{}) () {
    swap := arr[a] ; arr[a] = arr[b] ; arr[b] = swap
}

// CopyOf ...
func CopyOf(arr []interface{}) []interface{} {
    var newArr []interface{}
    
    for _, val := range arr {
        newArr = append(newArr, val)
    }
    return newArr
}

// FindIndexLp ...
func FindIndexLp(data interface{}, arr []interface{}) int {
    for i, el := range arr {
    	if data == el { return i }
    }
    return -1
}

// ReverseLp ...
func ReverseLp(arr []interface{}) () {
    log.Info("ReverseLp()")
    for i, j := 0, len(arr) - 1; j > i; i, j = i + 1, j - 1 {
    	SwapItems(i, j, arr)
    }
}
