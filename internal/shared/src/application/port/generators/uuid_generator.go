package generators

type UUIDGenerator interface {
	Generate() (string, error)
}
