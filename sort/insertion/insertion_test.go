package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	items := []int{29, 10, 14, 37, 14}
	itemsCopy := make([]int, len(items))
	copy(itemsCopy, items)
	want := []int{10, 14, 14, 29, 37}
	insertionSort(items)
	if !reflect.DeepEqual(want, items) {
		t.Errorf("insertionSort(%v) = %v got; want %v", itemsCopy, items, want)
	}
}
