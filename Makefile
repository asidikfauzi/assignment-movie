#Go Params
GO_APP=go
GO_BUILD=$(GO_APP) build
GO_MOD=$(GO_APP) mod

APP_DIR=cmd
BIN_NAME=app
BIN_MIGRATE_NAME=migrate
BIN_SEED_NAME=seed
CRON_BIN_NAME=cron

MIG_DIR=migrate
BIN_MIGRATE=migrate

all: mod build-main run
build: build-main
reload: build-main run
migrate: mod migrate-run run-migrate
seed: mod seed-run run-seed
mod:
	$(GO_MOD) vendor -v
build-main:
	rm -f $(APP_DIR)/$(BIN_NAME)/$(BIN_NAME)
	$(GO_BUILD) -o $(APP_DIR)/$(BIN_NAME)/$(BIN_NAME) $(APP_DIR)/${BIN_NAME}/app.go
run:
	./$(APP_DIR)/$(BIN_NAME)/$(BIN_NAME)
migrate-run:
	rm -f $(APP_DIR)/$(BIN_MIGRATE_NAME)/$(BIN_MIGRATE_NAME)
	$(GO_BUILD) -o $(APP_DIR)/$(BIN_MIGRATE_NAME)/$(BIN_MIGRATE_NAME) $(APP_DIR)/${BIN_MIGRATE_NAME}/migrate.go
run-migrate:
	./$(APP_DIR)/$(BIN_MIGRATE_NAME)/$(BIN_MIGRATE_NAME)
seed-run:
	rm -f $(APP_DIR)/$(BIN_MIGRATE_NAME)/$(BIN_MIGRATE_NAME)
	$(GO_BUILD) -o $(APP_DIR)/$(BIN_SEED_NAME)/$(BIN_SEED_NAME) $(APP_DIR)/${BIN_SEED_NAME}/seed.go
run-seed:
	./$(APP_DIR)/$(BIN_SEED_NAME)/$(BIN_SEED_NAME)