package array_test

import (
	"testing"

	"github.com/MrHuxu/xgo/array"
)

func TestEvery(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if array.Every(arr, func(value int) bool {
		return value%2 == 1
	}) {
		t.Errorf("invoke Every on %v should return %v", arr, false)
	}
}
