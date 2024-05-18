package entity

import "github.com/google/uuid"

type Employee struct {
	ID        uuid.UUID
	CompanyID uuid.UUID
	Name      string
	Email     string
	Password  string
}
