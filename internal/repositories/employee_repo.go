package repository

import (
	"context"
	"net/http"

	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
	constant "github.com/kidboy-man/ddd-attendance/internal/repositories/constants"
	model "github.com/kidboy-man/ddd-attendance/internal/repositories/models"
	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(ctx context.Context, employee entity.Employee) (created *entity.Employee, err error)
	Find(ctx context.Context, filter datatype.EmployeeFilter) (employees []*entity.Employee, err error)
}

type employeeRepo struct {
	db *gorm.DB // TODO: should be agnostic
}

func NewEmployeeRepo(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}

}

func (r *employeeRepo) Find(ctx context.Context, filter datatype.EmployeeFilter) (employees []*entity.Employee, err error) {
	var models []model.Employee
	qs := r.db.WithContext(ctx).Model(&model.Employee{})
	if len(filter.IDs) > 0 {
		qs = qs.Where("id IN (?)", filter.IDs)
	}

	if len(filter.Emails) > 0 {
		qs = qs.Where("email IN (?)", filter.Emails)
	}

	if len(filter.CompanyIDs) > 0 {
		qs = qs.Where("company_id IN (?)", filter.CompanyIDs)
	}

	err = qs.Find(&models).Error
	if err != nil {
		err = &generic.CustomError{
			Code:       constant.QueryFailureErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	for _, model := range models {
		employee := model.ToEntity()
		employees = append(employees, &employee)
	}
	return
}

func (r *employeeRepo) Create(ctx context.Context, employee entity.Employee) (created *entity.Employee, err error) {
	var model model.Employee
	model.FromEntity(employee)

	err = r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		err = &generic.CustomError{
			Code:       constant.QueryFailureErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return
	}

	result := model.ToEntity()
	return &result, nil
}
