package handler

import "github.com/labstack/echo"

var (
	NewPerson = newPerson
)

type Person = person

func (p *Person) Create(pp *person, c echo.Context) error {
	return pp.create(c)
}

func (p *Person) GetAll(pp *person, c echo.Context) error {
	return pp.getAll(c)
}
