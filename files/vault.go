package main

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Persons   []Person  `json:"persons"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault() *Vault {
	return &Vault{
		Persons:   []Person{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) addPerson(person Person) {
	vault.Persons = append(vault.Persons, person)
}

func (vault *Vault) addPersons(persons []Person) {
	vault.Persons = append(vault.Persons, persons...)
}
