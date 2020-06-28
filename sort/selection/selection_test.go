package selection

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	args := []int{29, 10, 14, 37, 14}
	want := []int{10, 14, 14, 29, 37}
	got := selectionSort(args)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("selectionSort(%v) = %v got; want %v", args, got, want)
	}
}

func TestSelectionSortReverse(t *testing.T) {
	args := []int{29, 10, 14, 37, 14}
	want := []int{37, 29, 14, 14, 10}
	got := selectionSortReverse(args)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("selectionSortReverse(%v) = %v got; want %v", args, got, want)
	}
}
