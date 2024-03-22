package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt *time.Time     `json:"created_at" gorm:"index"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// JSONB is a type for JSON bytes
// it implements its valuer and scanner
type JSONB map[string]any

func (j *JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value any) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("value is not of type []byte")
	}
	return json.Unmarshal(v, &j)
}

// IsValidUUID checks if a string
// is a valid UUID or not
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
