package main

import (
	"context"

	"github.com/cryptellation/candlesticks/dagger/internal/dagger"
)

// PostgresContainer returns a service running Postgres initialized for integration tests.
func PostgresContainer(ctx context.Context, dag *dagger.Client, sourceDir *dagger.Directory) *dagger.Service {
	// Read the SQL init file from the repo
	initSQL := sourceDir.File("deployments/docker-compose/postgresql/cryptellation.sql")

	// Start a Postgres container
	c := dag.Container().
		From("postgres:15-alpine").
		WithEnvVariable("POSTGRES_PASSWORD", "postgres").
		WithEnvVariable("POSTGRES_USER", "postgres").
		WithEnvVariable("PGUSER", "postgres").
		WithEnvVariable("PGPASSWORD", "postgres").
		WithEnvVariable("POSTGRES_DB", "postgres") // default DB

	// Copy the init SQL into the container and run it after startup
	c = c.WithMountedFile("/docker-entrypoint-initdb.d/cryptellation.sql", initSQL)

	// Expose the default Postgres port
	c = c.WithExposedPort(5432)

	return c.AsService()
}
