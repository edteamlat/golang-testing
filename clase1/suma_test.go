package clase1

import "testing"

func TestAdd(t *testing.T) {
	want := 9
	t.Logf("want vale: %d\n", want)
	got := Add(2, 3)
	t.Logf("got vale: %d\n", got)
	if got != want {
		t.Errorf("Se esperaba %d, se obtuvo %d", want, got)
	}
	t.Log("Terminó la prueba de Add")
}

func TestAddAcum(t *testing.T) {
	want := 6
	got := AddAcum(1, 2, 3)
	if got != want {
		t.Errorf("Se esperaba %d, se obtuvo %d", want, got)
	}
}
