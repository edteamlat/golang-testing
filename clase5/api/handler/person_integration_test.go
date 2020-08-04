package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EDteam/golang-testing/clase5/api/handler"
	"github.com/EDteam/golang-testing/clase5/api/storage"
	"github.com/labstack/echo"
)

func TestCreate(t *testing.T) {
	data := []byte(`{"name": "Alexys", "age": 40, "communities": []}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	memory := storage.NewMemory()
	p := handler.NewPerson(&memory)
	err := p.Create(&p, ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	dataStorage, _ := memory.GetAll()
	t.Logf("Memory: %v", dataStorage)
}
