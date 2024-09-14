package datatype

import "github.com/google/uuid"

type Registration struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CompanyID uuid.UUID `json:"company_id"`
}

type Registered struct {
	EmployeeID uuid.UUID `json:"employee_id"`
	Token      string    `json:"token"`
}
