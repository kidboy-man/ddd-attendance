package model

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/domain/entities"
	constant "github.com/kidboy-man/ddd-attendance/internal/domain/repositories/constants"
	schema "github.com/kidboy-man/ddd-attendance/internal/schemas"

	"gorm.io/gorm"
)

type AttendanceLog struct {
	ID           uint64    `gorm:"primaryKey"`
	EmployeeID   uuid.UUID `gorm:"column:employee_id"`
	Note         sql.NullString
	Date         sql.NullTime `gorm:"autoCreateTime;<-:create"`
	ClockedInAt  sql.NullTime
	ClockedOutAt sql.NullTime
}

func (AttendanceLog) TableName() string {
	return "attendance_logs"
}

func (a AttendanceLog) ToEntity() (en entity.AttendanceLog) {
	en.ID = a.ID
	en.EmployeeID = a.EmployeeID
	en.Note = a.Note.String
	en.Date = a.Date.Time
	en.ClockedInAt = a.ClockedInAt.Time
	en.ClockedOutAt = a.ClockedOutAt.Time
	return
}

func (a *AttendanceLog) FromEntity(en entity.AttendanceLog) {
	a.ID = en.ID
	a.EmployeeID = en.EmployeeID
	a.Note.String = en.Note
	a.Date.Time = en.Date
	a.ClockedInAt.Time = en.ClockedInAt
	a.ClockedOutAt.Time = en.ClockedOutAt
}

func (a *AttendanceLog) BeforeCreate(tx *gorm.DB) (err error) {
	if a.EmployeeID == uuid.Nil {
		err = &schema.CustomError{
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
