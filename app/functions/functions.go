package functions

import (
	"github.com/gofrs/uuid"
)

// NewUUID ...
func NewUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
