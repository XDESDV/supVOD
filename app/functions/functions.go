package functions

import (
	"github.com/gofrs/uuid"
)

func NewUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
