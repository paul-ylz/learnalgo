package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	args := []int{5, 3, 4, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	got := bubbleSort(args)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("bubbleSort(%v) = %v got; want %v", args, got, want)
	}
}
