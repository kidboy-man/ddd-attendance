package model

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/domain/entities"
	constant "github.com/kidboy-man/ddd-attendance/internal/domain/repositories/constants"
	schema "github.com/kidboy-man/ddd-attendance/internal/schemas"
	"gorm.io/gorm"
)

type Employee struct {
	ID        uuid.UUID    `gorm:"primaryKey"`
	CompanyID uuid.UUID    `gorm:"column:company_id"`
	Name      string       `gorm:"type:varchar(255)" validate:"required"`
	CreatedAt sql.NullTime `gorm:"autoCreateTime;<-:create"`
	UpdatedAt sql.NullTime `gorm:"autoUpdateTime"`
}

func (Employee) TableName() string {
	return "employees"
}

func (e Employee) ToEntity() (en entity.Employee) {
	en.ID = e.ID
	en.Name = e.Name
	en.CompanyID = e.CompanyID
	return
}

func (e *Employee) FromEntity(en entity.Employee) {
	e.ID = en.ID
	e.Name = en.Name
	e.CompanyID = en.CompanyID
}

func (e *Employee) setAttr() (err error) {
	e.Name = strings.TrimSpace(e.Name)
	if e.CompanyID == uuid.Nil {
		err = &schema.CustomError{
			Code:       constant.OrmHookValidationErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Incomplete employee data, missing required values",
		}
	}

	return
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	err = e.setAttr()
	if err != nil {
		return
	}
	e.ID = uuid.New()
	return
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) (err error) {
	err = e.setAttr()
	return
}
