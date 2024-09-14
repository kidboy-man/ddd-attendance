package repository

import (
	"context"

	"github.com/google/uuid"
	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
)

type CompanyRepo interface {
	Create(ctx context.Context, company entity.Company) (err error)
	Find(ctx context.Context, filter datatype.CompanyFilter) (companies []*entity.Company, err error)
	GetByID(ctx context.Context, id uuid.UUID) (company *entity.Company, err error)
	Update(ctx context.Context, company entity.Company) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}

type companyRepo struct {
}

func NewCompanyRepo() *companyRepo {
	return &companyRepo{}

}
