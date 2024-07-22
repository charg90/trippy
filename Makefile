# Variables
GO := go
DB_MIGRATE := db/migrate/migrate.go

# Default target
.PHONY: all
all: migrate

# Migrate target
.PHONY: migrate
migrate:
	$(GO) run $(DB_MIGRATE)


# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migrate    - Run database migrations"
	@echo "  make help       - Show this help message"
