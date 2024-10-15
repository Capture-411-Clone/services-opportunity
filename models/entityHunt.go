package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
)

// JSONMap defined JSON data type, need to implements driver.Valuer, sql.Scanner interface
/*
 * @apiDefine: JSONMap
 */
type JSONMap map[string]interface{}

/*
 * @apiDefine: EntityHunt
 */
type EntityHunt struct {
	ID                 uint                  `gorm:"primaryKey;uniqueIndex:udx_entity_hunts" openapi:"example:1;type:integer;"`
	Details            JSONMap               `json:"details" openapi:"example:{};type:object;"`
	SolicitationNumber string                `json:"solicitation_number" openapi:"example:123456;type:string;"`
	Year               string                `json:"year" gorm:"size:4;" openapi:"example:2022;type:string;"`
	FileName           string                `json:"file_name" openapi:"example:filename;type:string;"`
	FilePath           string                `json:"file_path" openapi:"example:user/documents/report.pdf;type:string"`
	CreatedAt          time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt          time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt          soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_entity_hunts"`
}

// Value return json value, implement driver.Valuer interface
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (m *JSONMap) Scan(val interface{}) error {
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := map[string]interface{}{}
	err := json.Unmarshal(ba, &t)
	*m = JSONMap(t)
	return err
}

// MarshalJSON to output non base64 encoded []byte
func (m JSONMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	t := (map[string]interface{})(m)
	return json.Marshal(t)
}

// UnmarshalJSON to deserialize []byte
func (m *JSONMap) UnmarshalJSON(b []byte) error {
	t := map[string]interface{}{}
	err := json.Unmarshal(b, &t)
	*m = JSONMap(t)
	return err
}

// GormDataType gorm common data type
func (m JSONMap) GormDataType() string {
	return "jsonmap"
}

// GormDBDataType gorm db data type
func (JSONMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
