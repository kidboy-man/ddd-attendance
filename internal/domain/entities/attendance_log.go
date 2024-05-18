package entity

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceLog struct {
	ID           uint64
	EmployeeID   uuid.UUID
	ClockedInAt  time.Time
	ClockedOutAt time.Time
	Date         time.Time
	Note         string
}
