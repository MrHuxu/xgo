package array_test

import (
	"reflect"
	"testing"

	"github.com/MrHuxu/xgo/array"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ret := []int{1, 3, 5}
	if !reflect.DeepEqual(array.Filter(arr, func(_ int, value int) bool {
		return value%2 == 1
	}), ret) {
		t.Errorf("invoke Filter on %v should return %v", arr, ret)
	}
}
