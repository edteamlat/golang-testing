package clase3

import (
	"reflect"
	"testing"
)

func TestPerro(t *testing.T) {
	want := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criollo",
		},
	}
	got := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criolla",
		},
	}

	// t.Logf("want %p, got %p", want, got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Se esperaba: %v, se obtuvo %v", want, got)
	}
}
