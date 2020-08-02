package framework

import (
	"net/http"

	"github.com/labstack/echo"
)

// Person estructura de una persona
type Person struct {
	Name string
	Age  int
}

// Get retorna una persona
func Get(c echo.Context) error {
	p := Person{
		"Jhoana",
		31,
	}

	return c.JSON(http.StatusOK, p)
}
