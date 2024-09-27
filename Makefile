# Makefile

SWAGGER_OUTPUT_DIR = ./internal/http/docs
SWAGGER_SEARCH_DIR = ./cmd,./internal/http

.PHONY: swagger build

swagger:
	@echo "Generating Swagger documentation..."
	@swag init --output $(SWAGGER_OUTPUT_DIR) --dir $(SWAGGER_SEARCH_DIR)
	@echo "Swagger documentation generated in $(SWAGGER_OUTPUT_DIR)"

build:
	@echo "Building Docker images and starting services..."
	@docker compose up --build
