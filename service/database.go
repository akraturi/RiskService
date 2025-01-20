package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Risk struct {
	Id          uuid.UUID
	Title       string
	State       string
	Description string
}

type Database interface {
	GetRisks() ([]Risk, error)
	GetRiskById(id uuid.UUID) (Risk, error)
	AddRisk(risk Risk) (Risk, error)
}

type InMemoryDatabase struct {
	risks []Risk
}

func (d InMemoryDatabase) GetRisks() ([]Risk, error) {
	return d.risks, nil
}

func (d InMemoryDatabase) GetRiskById(id uuid.UUID) (Risk, error) {
	for _, risk := range d.risks {
		if risk.Id == id {
			return risk, nil
		}
	}
	return Risk{}, errors.New(fmt.Sprintf("risk with id %s not found in db", id.String()))
}

func (d InMemoryDatabase) AddRisk(risk Risk) (Risk, error) {
	risk.Id = uuid.Must(uuid.NewV6())
	d.risks = append(d.risks, risk)
	return risk, nil
}

func NewDatabase() Database {
	db := InMemoryDatabase{}
	// Initialize db with some predefined risks
	db.risks = []Risk{
		{Id: uuid.Must(uuid.NewV6()), Title: "Risk1", Description: "Description1", State: "open"},
		{Id: uuid.Must(uuid.NewV6()), Title: "Risk2", Description: "Description2", State: "closed"},
		{Id: uuid.Must(uuid.NewV6()), Title: "Risk3", Description: "Description3", State: "accepted"},
	}
	return db
}
