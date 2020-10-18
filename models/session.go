package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	StudentID   uint `gorm:"not null;index"`
	Language    sql.NullInt32 `gorm:"not null"`
	Preferences sql.NullString
}
