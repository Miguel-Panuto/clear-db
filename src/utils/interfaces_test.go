package utils

import "testing"

func TestToInterfaceArr(t *testing.T) {
	got := ToInterfaceArr([]string{"name", "id"})
	want := []interface{}{"name", "id"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}

func TestToStringArr(t *testing.T) {
	got := ToStringArr([]interface{}{"name", "id", 0, 2.54, true})
	want := []string{"name", "id", "0", "2.54", "true"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}
