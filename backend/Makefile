include .env

MIGRATION_DIR=./migrations
DB_DRIVER=postgres


run:
	go run cmd/main.go

build:
	go build -o app cmd/main.go

# Cria uma nova migration SQL: make create-migration NAME=nome_da_migration
create-migration:
	@if [ -z "$(NAME)" ]; then echo "Erro: Variavel NAME nao definida. Use: make create-migration NAME=exemplo"; exit 1; fi
	goose -dir $(MIGRATION_DIR) create $(NAME) sql

# Aplica todas as migrations: make migrate-up
migrate-up:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" up

# Desfaz a ultima migration: make migrate-down
migrate-down:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" down

migrate-down-0:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" down-to 0

# Para limpeza de banco de desenvolvimento
migrate-reset:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" down-to 0
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" up

# Mostra o status das migrations: make migrate-status
migrate-status:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) status