package state_machine

import "testing"

func TestCheck(t *testing.T) {
	check(nil)
}

func TestSliceString(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	if getSliceString(arr) != "1, 2, 3, 4" {
		t.Fatal("Got wrong string")
	}

	arr_1 := []int{1, 2, 3, 4}
	arr_2 := []int{1, 2, 3, 4}
	if getSliceString(arr_1) != getSliceString(arr_2) {
		t.Fatal("Func makes different strings")
	}
}
