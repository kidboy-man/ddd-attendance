package entity

import (
	"time"

	"github.com/google/uuid"
	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
)

type AttendanceLog struct {
	ID         uint64
	EmployeeID uuid.UUID
	Note       string
	Action     datatype.Action
	CreatedAt  time.Time
}
