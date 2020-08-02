package clase1

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)
	if got != want {
		t.Logf("Error: Se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}

func TestAddAcum(t *testing.T) {
	want := 6
	got := AddAcum(1, 2, 3)
	if got != want {
		t.Logf("Error: Se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
