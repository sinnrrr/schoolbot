package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name sql.NullString `gorm:"not null;size:255"`
}
