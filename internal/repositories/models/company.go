package model

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	entity "github.com/kidboy-man/ddd-attendance/internal/entities"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID    `gorm:"primaryKey"`
	Name      string       `gorm:"type:varchar(255)" validate:"required"`
	CreatedAt sql.NullTime `gorm:"autoCreateTime;<-:create"`
	UpdatedAt sql.NullTime `gorm:"autoUpdateTime"`
}

func (Company) TableName() string {
	return "companies"
}

func (c Company) ToEntity() (en entity.Company) {
	en.ID = c.ID
	en.Name = c.Name
	return
}

func (c *Company) FromEntity(en entity.Company) {
	c.ID = en.ID
	c.Name = en.Name
}

func (c *Company) setAttr() {
	c.Name = strings.TrimSpace(c.Name)
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.setAttr()
	c.ID = uuid.New()
	return
}

func (c *Company) BeforeUpdate(tx *gorm.DB) (err error) {
	c.setAttr()
	return
}
