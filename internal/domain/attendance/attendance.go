package domain

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
)

type AttendanceDomain interface {
	ClockIn(ctx context.Context, employeeID uuid.UUID) (err error)
	ClockOut(ctx context.Context, employeeID uuid.UUID) (err error)
	GetEmployeeAttendanceLog(ctx context.Context, employeeID uuid.UUID) (attendanceLogs []entity.AttendanceLog, err error)
}
