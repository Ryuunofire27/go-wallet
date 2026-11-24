package generators

import (
	"fmt"
	"go-wallet/internal/shared/src/application/port/generators"
	apperror "go-wallet/internal/shared/src/domain/error"

	"gorm.io/gorm"
)

type PostgresSequenceGenerator struct {
	nameService string
	db          *gorm.DB
	sequences   map[generators.EntityType]string
}

func NewPostgresSequenceGenerator(
	db *gorm.DB,
	sequences map[generators.EntityType]string,
) *PostgresSequenceGenerator {
	return &PostgresSequenceGenerator{
		nameService: "PostgresSequenceGenerator",
		db:          db,
		sequences:   sequences,
	}
}

func (g *PostgresSequenceGenerator) Generate(
	entityType generators.EntityType,
) (uint64, error) {
	sequence, ok := g.sequences[entityType]
	if !ok {
		return 0, apperror.New(
			apperror.Internal,
			"SEQUENCE_NOT_CONFIGURED",
			"sequence not configured for entity: "+string(entityType),
			fmt.Sprintf("%s.%s", g.nameService, "Generate"),
		)
	}

	var id uint64

	err := g.db.Raw(
		"SELECT nextval(?::regclass)",
		sequence,
	).Scan(&id).Error

	if err != nil {
		return 0, apperror.Wrap(
			apperror.Internal,
			"SEQUENCE_GENERATION_FAILED",
			"failed to generate sequence number",
			fmt.Sprintf("%s.%s", g.nameService, "Generate"),
			err,
		)
	}

	return id, nil
}
