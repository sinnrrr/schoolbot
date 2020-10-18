package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	ClassID uint           `gorm:"not null;index"`
	Subject sql.NullString `gorm:"not null;size:255"`
	Task    sql.NullString `gorm:"not null"`
}
