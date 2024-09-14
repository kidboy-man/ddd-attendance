package domain

import (
	"context"
	"net/http"

	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	constant "github.com/kidboy-man/ddd-attendance/internal/domain/account/constants"
	helper "github.com/kidboy-man/ddd-attendance/internal/domain/account/helpers"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
)

func (d *accountDomain) Register(ctx context.Context, param datatype.Registration) (res datatype.Registered, err error) {
	cleanedEmail, err := helper.CleanEmail(param.Email)
	if err != nil {
		return
	}

	employees, err := d.employeeRepo.Find(ctx, datatype.EmployeeFilter{
		Emails: []string{cleanedEmail},
	})
	if err != nil {
		return
	}

	if len(employees) > 0 {
		err = &generic.CustomError{
			Code:       constant.EmailIsAlreadyRegisteredErrCode,
			HTTPStatus: http.StatusConflict,
			Message:    "Email already registered",
		}
		return
	}

	company, err := d.companyRepo.GetByID(ctx, param.CompanyID)
	if err != nil {
		return
	}

	hashedPassword, err := helper.HashPassword(param.Password)
	if err != nil {
		return
	}

	employee, err := d.employeeRepo.Create(ctx, entity.Employee{
		Name:      param.Name,
		CompanyID: company.ID,
		Email:     cleanedEmail,
		Password:  hashedPassword,
	})
	if err != nil {
		return
	}

	res.EmployeeID = employee.ID
	return
}
