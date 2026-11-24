package generators

type EntityType string

type SequenceNumberGenerator interface {
	Generate(entityType EntityType) (uint64, error)
}
