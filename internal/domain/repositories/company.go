package repository

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/domain/entities"
)

type CompanyRepo interface {
	Create(ctx context.Context, company *entity.Company) (err error)
	Find(ctx context.Context) (companies []*entity.Company, err error)
	GetByID(ctx context.Context, id uuid.UUID) (company *entity.Company, err error)
	Update(ctx context.Context, company *entity.Company)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
