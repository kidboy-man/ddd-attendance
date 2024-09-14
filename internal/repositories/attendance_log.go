package repository

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
)

type AttendanceLogRepo interface {
	Create(ctx context.Context, attendanceLog *entity.AttendanceLog) (err error)
	Find(ctx context.Context) (attendanceLogs []*entity.AttendanceLog, err error)
	GetByID(ctx context.Context, id uint64) (attendanceLog *entity.AttendanceLog, err error)
	GetByEmployeeID(ctx context.Context, employeeID uuid.UUID) (attendanceLog *entity.AttendanceLog, err error)
	Update(ctx context.Context, attendanceLog *entity.AttendanceLog)
	Delete(ctx context.Context, id uint64) (err error)
}
