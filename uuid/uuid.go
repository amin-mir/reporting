//go:generate mockgen -source uuid.go -destination mock/uuid_mock.go -package mock

package uuid

type Generator interface {
	Generate() string
	Parse(uuid string) error
}

func New() Generator {
	return &generator{}
}
