package generators

import (
	"uuid"
)

type UUIDGenerator struct{}

func (g *UUIDGenerator) Generate() (string, error) {
	newId := uuid.NewV7()

	return newId.String(), nil
}
