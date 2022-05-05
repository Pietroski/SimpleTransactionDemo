// Code generated by sqlc. DO NOT EDIT.

package sqlc_device_store

import (
	"database/sql"

	"github.com/google/uuid"
)

type Devices struct {
	RowID       sql.NullInt64 `json:"rowID"`
	DeviceID    uuid.UUID     `json:"deviceID"`
	UserID      uuid.UUID     `json:"userID"`
	DeviceName  string        `json:"deviceName"`
	DeviceBrand string        `json:"deviceBrand"`
	CreatedAt   sql.NullTime  `json:"createdAt"`
	UpdatedAt   sql.NullTime  `json:"updatedAt"`
}
