package datatype

import "github.com/google/uuid"

type EmployeeFilter struct {
	IDs        []uuid.UUID
	Emails     []string
	CompanyIDs []uuid.UUID
}

type CompanyFilter struct {
	IDs []uuid.UUID
}
