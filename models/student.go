package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ClassID      uint           `gorm:"not null;index"`
	FirstName    sql.NullString `gorm:"not null;size:64"`
	LastName     sql.NullString `gorm:"size:64"`
	Username     sql.NullString `gorm:"not null;size:32"`
	LanguageCode sql.NullString `gorm:"not null;size:16"`
}
