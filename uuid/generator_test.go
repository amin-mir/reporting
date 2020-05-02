package uuid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	generator := New()

	uuid := generator.Generate()

	err := generator.Parse(uuid)
	require.NoError(t, err)
}
