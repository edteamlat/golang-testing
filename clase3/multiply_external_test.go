package clase3_test

import (
	"testing"

	"github.com/EDteam/golang-testing/clase3"
)

func TestMultiply(t *testing.T) {
	want := 6
	got := clase3.Multiply(2, 3)
	if got != want {
		t.Errorf("Se esperaba %d, se obtuvo %d", want, got)
	}
}
