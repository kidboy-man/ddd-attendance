package domain

import (
	"context"

	database "github.com/kidboy-man/ddd-attendance/infrastructures/databases/psql"
	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	repository "github.com/kidboy-man/ddd-attendance/internal/repositories"
)

type AccountDomain interface {
	// Login(ctx context.Context, email, password string) (employee entity.Employee, err error)
	Register(ctx context.Context, param datatype.Registration) (res datatype.Registered, err error)
}

type accountDomain struct {
	employeeRepo repository.EmployeeRepo
	companyRepo  repository.CompanyRepo
}

func NewAccountDomain() *accountDomain {
	db := database.GetDB()
	return &accountDomain{
		employeeRepo: repository.NewEmployeeRepo(db),
	}
}
