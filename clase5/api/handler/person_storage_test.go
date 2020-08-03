package handler

import (
	"errors"

	"github.com/EDteam/golang-testing/clase5/api/model"
)

type PersonStorageOKMock struct{}

func (psm *PersonStorageOKMock) Create(person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Update(ID int, person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Delete(ID int) error {
	return nil
}
func (psm *PersonStorageOKMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}
func (psm *PersonStorageOKMock) GetAll() (model.Persons, error) {
	return nil, nil
}

type PersonStorageWrongMock struct{}

func (psm *PersonStorageWrongMock) Create(person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Update(ID int, person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Delete(ID int) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error")
}
func (psm *PersonStorageWrongMock) GetAll() (model.Persons, error) {
	return nil, errors.New("error")
}
