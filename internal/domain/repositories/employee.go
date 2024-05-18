package repository

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/domain/entities"
)

type EmployeeRepo interface {
	Create(ctx context.Context, employee *entity.Employee) (err error)
	Find(ctx context.Context) (employees []*entity.Employee, err error)
	GetByID(ctx context.Context, id uuid.UUID) (employee *entity.Employee, err error)
	GetByCompanyID(ctx context.Context, companyID uuid.UUID) (employee *entity.Employee, err error)
	Update(ctx context.Context, employee *entity.Employee)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
