# Variables
GO := go
DB_MIGRATE := db/migrate/migrate.go
DB_SEED := db/seeds/seeds.go

# Default target
.PHONY: all
all: migrate

# Migrate target
.PHONY: migrate
migrate:
	$(GO) run $(DB_MIGRATE)

# Seed target
.PHONY: seed
seed:
	$(GO) run $(DB_SEED)

# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migrate    - Run database migrations"
	@echo "  make seed       - Seed the database"
	