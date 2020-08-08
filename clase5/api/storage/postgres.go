package storage

import (
	"database/sql"
	"errors"
	"log"

	// Import pq library
	_ "github.com/lib/pq"

	"github.com/EDteam/golang-testing/clase5/api/model"
)

// Psql .
type Psql struct {
	db *sql.DB
}

// NewPsql .
func NewPsql() Psql {
	p := Psql{}
	p.db = getPSQLConn()

	return p
}

func getPSQLConn() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/edteam?sslmode=disable")
	if err != nil {
		log.Fatalf("no se pudo conectar a la bd: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("no se pudo hacer ping a la bd: %v", err)
	}

	return db
}

// Create .
func (psql *Psql) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	stmt, err := psql.db.Prepare(`INSERT INTO persons VALUES ($1, $2)`)
	if err != nil {
		log.Fatalf("no se pudo preparar la consulta: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(person.Name, person.Age)
	if err != nil {
		log.Fatalf("no se pudo insertar el registro: %v", err)
	}

	return nil
}

// Update actualiza una persona en el slice de memoria
func (psql *Psql) Update(ID int, person *model.Person) error {
	return errors.New("not implemented yet")
}

// Delete borra de la memoria la persona
func (psql *Psql) Delete(ID int) error {
	return errors.New("not implemented yet")
}

// GetByID retorna una persona por el ID
func (psql *Psql) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("not implemented yet")
}

// GetAll retorna todas las personas que están en la memoria
func (psql *Psql) GetAll() (model.Persons, error) {
	stmt, err := psql.db.Prepare(`SELECT name, age FROM persons`)
	if err != nil {
		log.Fatalf("no se pudo preparar la consulta: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("no se pudo consultar los registros %v", err)
	}
	defer rows.Close()

	var persons model.Persons
	for rows.Next() {
		p := model.Person{}
		err = rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalf("no se pudo mapear la persona %v", err)
		}

		persons = append(persons, p)
	}

	return persons, nil
}

// DeleteAll .
func (psql *Psql) DeleteAll() error {
	stmt, err := psql.db.Prepare(`TRUNCATE TABLE persons`)
	if err != nil {
		log.Fatalf("no se pudo preparar la consulta: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("no se pudo borrar la tabla: %v", err)
	}

	return nil
}

// CloseDB .
func (psql *Psql) CloseDB() {
	psql.db.Close()
}
