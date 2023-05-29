package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid/v2"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a ULID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	ulid := ulid.Make()
	return scope.SetColumn("ID", ulid)
}
