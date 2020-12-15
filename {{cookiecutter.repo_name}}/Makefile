PROJECT_NAME := "{{cookiecutter.repo_name}}"
PKG := "{{cookiecutter.repo_path}}/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: lint test gomodgen install build clean kill frontend serve race coverage coverhtml migrate

lint: ## Lint the files
	@golangci-lint run ${${PKG}/...}

test: ## Run unittests
	@go test -short ${PKG_LIST}

install: gomodgen
	@go mod download

gomodgen:
	./gomod.sh;

build: gomodgen
	GO111MODULE=on go build -ldflags="-s -w" -o ./bin/{{cookiecutter.repo_name}} ./cmd/{{cookiecutter.repo_name}}/main.go

kill:
	@kill `cat $(PID)` || true

race: ## Run data race detector
	@go test -race ${PKG_LIST}

coverage: ## Generate global code coverage report
	@bash ./coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	@bash ./coverage.sh html;

migrate:
	@goose -dir migrations postgres "$(DSN)" $(CMD)