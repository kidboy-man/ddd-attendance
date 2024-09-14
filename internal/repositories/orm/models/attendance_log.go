package model

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
	constant "github.com/kidboy-man/ddd-attendance/internal/repositories/constants"

	"gorm.io/gorm"
)

type AttendanceLog struct {
	ID         uint64    `gorm:"primaryKey"`
	EmployeeID uuid.UUID `gorm:"column:employee_id"`
	Note       sql.NullString
	Action     string
	CreatedAt  sql.NullTime `gorm:"autoCreateTime;<-:create"`
}

func (AttendanceLog) TableName() string {
	return "attendance_logs"
}

func (a AttendanceLog) ToEntity() (en entity.AttendanceLog) {
	en.ID = a.ID
	en.EmployeeID = a.EmployeeID
	en.Note = a.Note.String
	en.Action = datatype.Action(a.Action)
	en.CreatedAt = a.CreatedAt.Time
	return
}

func (a *AttendanceLog) FromEntity(en entity.AttendanceLog) {
	a.ID = en.ID
	a.EmployeeID = en.EmployeeID
	a.Note.String = en.Note
	a.Action = string(en.Action)
	a.CreatedAt.Time = en.CreatedAt
}

func (a *AttendanceLog) BeforeCreate(tx *gorm.DB) (err error) {
	if a.EmployeeID == uuid.Nil || a.Action == "" {
		err = &generic.CustomError{
			Code:       constant.OrmHookValidationErrCode,
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Incomplete attendance_log data, missing required values",
		}
	}
	return
}

func (e *AttendanceLog) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
