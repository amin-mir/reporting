package uuid

import "github.com/google/uuid"

type generator struct{}

func (g *generator) Generate() string {
	return uuid.New().String()
}

func (g *generator) Parse(uuidStr string) error {
	_, err := uuid.Parse(uuidStr)
	return err
}
