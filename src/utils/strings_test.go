package utils

import "testing"

func TestSplit(t *testing.T) {
	got := Split("Name:Miguel", ":")
	want := []string{"Name", "Miguel"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}
func TestSplitWithEmptyString(t *testing.T) {
	got := Split("", " ")

	if len(got) > 0 {
		t.Errorf("should be 0 got len of %d", len(got))
	}
}

func TestTrimSplit(t *testing.T) {
	got := TrimSplit("Name : Miguel", ":")
	want := []string{"Name", "Miguel"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}

func TestMultipleSplit(t *testing.T) {
	got := MultipleSplit("Name : Miguel;Panuto", ":", ";")
	want := []string{"Name", "Miguel", "Panuto"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
	got = MultipleSplit("Name : MiguelPanuto", ":", ";")
	want = []string{"Name", "MiguelPanuto"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}

func TestSubstring(t *testing.T) {
	got := SubString("{Name}", "{", "}")
	want := "Name"

	if got != want {
		t.Errorf("not expected value got: %s - want: %s", got, want)
	}
}

func TestSubSplit(t *testing.T) {
	got := SubSplit("{Name : Miguel}", "{", "}", ":")
	want := []string{"Name", "Miguel"}

	if len(got) != len(want) {
		t.Errorf("they has to be the same size: %d - want: %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("not expected value got: %s - want: %s", got[i], want[i])
		}
	}
}

func TestLowerPrefix(t *testing.T) {
	got := VerifyLowerPrefix(" Use certain db", "use")
	want := true

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}

	got = VerifyLowerPrefix(" Use certain db", "notuse")
	want = false

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}
}

func TestLower(t *testing.T) {
	got := VerifyLower(" Use ", "use")
	want := true

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}

	got = VerifyLower(" Use ", "notuse")
	want = false

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}
}

func TestMakeStringArr(t *testing.T) {
	testObj := []interface{}{1, 2, 3.53, "name"}
	got := MakeStringArr(testObj)
	want := []string{"1", "2", "3.53", "name"}

	if len(got) != len(want) {
		t.Errorf("different sizes got %d want %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("they are different got: %s - want: %s", got[i], want[i])
		}
	}
}

func TestContainsMany(t *testing.T) {
	got := ContainsMany("{Place}", "{", "}")
	want := true

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}

	got = ContainsMany("{Place", "{", "}")
	want = false
	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}
}

func TestContainsInside(t *testing.T) {
	got := ContainsInside([]string{"required"}, "required")
	want := true

	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}

	got = ContainsInside([]string{"required"}, "unique")
	want = false
	if got != want {
		t.Errorf("not expected value got: %t - want: %t", got, want)
	}
}
