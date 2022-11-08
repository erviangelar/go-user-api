package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BaseMode struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index" sql:"index"`
}

func (base *BaseMode) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	base.ID = uid
	base.CreatedAt = time.Now()
	return nil
}

func (base *BaseMode) BeforeUpdate(tx *gorm.DB) error {
	base.CreatedAt = time.Now()
	return nil
}

func (base *BaseMode) MarshalJSON() ([]byte, error) {
	type Alias BaseMode
	return json.Marshal(&struct {
		*Alias
		Stamp string `json:"stamp"`
	}{
		Alias: (*Alias)(base),
		Stamp: base.CreatedAt.Format("01/02/06"),
	})
}
