package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	require := require.New(t)

	os.Setenv("SCYLLA_HOSTS", "host1.com,host2.com")
	os.Setenv("SCYLLA_MIGRATIONS_DIR", "./cql")
	os.Setenv("SCYLLA_KEYSPACE", "reporting")

	actual, err := Load()
	require.NoError(err)

	expected := Config{
		ScyllaHosts:         []string{"host1.com", "host2.com"},
		ScyllaKeyspace:      "reporting",
		ScyllaMigrationsDir: "./cql",
	}
	require.Equal(expected, actual)
}
