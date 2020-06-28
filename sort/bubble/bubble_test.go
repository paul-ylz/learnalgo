package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	args := []int{29, 10, 14, 37, 14 }
	want := []int{10, 14, 14, 29, 37}
	got := bubbleSort(args)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("bubbleSort(%v) = %v got; want %v", args, got, want)
	}
}
