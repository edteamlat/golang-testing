package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EDteam/golang-testing/clase5/api/handler"
	"github.com/EDteam/golang-testing/clase5/api/model"
	"github.com/EDteam/golang-testing/clase5/api/storage"
	"github.com/labstack/echo"
)

func TestCreate_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})
	data := []byte(`{"name": "Alexys", "age": 40, "communities": []}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := handler.NewPerson(&store)
	err := p.Create(&p, ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}
}

type responsePerson struct {
	MessageType string        `json:"message_type"`
	Message     string        `json:"message"`
	Data        model.Persons `json:"data"`
}

func TestGetAll_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})
	insertData(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := handler.NewPerson(&store)
	err := p.GetAll(&p, ctx)
	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo %v", err)
	}

	var response responsePerson
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("no se pudo hacer marshal de la respuesta: %v", err)
	}

	lenPersonsWant := 2
	lenPersonsGot := len(response.Data)
	if lenPersonsGot != lenPersonsWant {
		t.Errorf("Se esperaban %d personas, se obtuvieron %d", lenPersonsWant, lenPersonsGot)
	}
}

func insertData(t *testing.T) {
	store := storage.NewPsql()
	err := store.Create(&model.Person{Name: "Alexys", Age: 40})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

	err = store.Create(&model.Person{Name: "Matthew", Age: 4})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

	store.CloseDB()
}

func cleanData(t *testing.T) {
	store := storage.NewPsql()
	err := store.DeleteAll()
	if err != nil {
		t.Fatalf("no se pudo limpiar la tabla %v", err)
	}

	store.CloseDB()
}
