include .env
include .secret.env
export

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: name
name: ## display app name
	@echo ${APP_NAME}

.PHONY: aqua
aqua: ## insatll aqua
	@brew install aquaproj/aqua/aqua

.PHONY: tool
tool: ## install tool
	@aqua i
	@curl https://get.okteto.com -sSfL | sh

.PHONY: compile
compile: ## go compile
	@go build -v ./... && go clean

.PHONY: fmt
fmt: ## go format
	@go fmt ./...

.PHONY: lint
lint: ## go lint
	@golangci-lint run --fix

.PHONY: update
update: ## go modules update
	@go get -u -t ./...
	@go mod vendor

.PHONY: mod
mod: ## go mod tidy & go mod vendor
	@go mod tidy
	@go mod vendor

.PHONY: test
test: ## unit test
	@$(call _test,${c})

define _test
if [ -z "$1" ]; then \
	go test ./internal/... ; \
else \
	go test ./internal/... -count=1 ; \
fi
endef

.PHONY: up
up: ## docker compose up with air hot reload
	@docker compose --project-name ${APP_NAME} --file ./.docker/docker-compose.yaml up -d

.PHONY: down
down: ## docker compose down
	@docker compose --project-name ${APP_NAME} down --volumes

.PHONY: log
log: ## docker log
	@docker logs ${APP_NAME}-app

.PHONY: ymlfmt
ymlfmt: ## yaml file format
	@yamlfmt

.PHONY: ymlint
ymlint: ## yaml file lint
	@yamlfmt -lint && actionlint

.PHONY: key
key: ## generate age key
	@age-keygen

.PHONY: secret
secret: ## Create kubernates secret yaml. ex) make secret secret=password
	@script/secret.sh ${secret}

.PHONY: encrypt
encrypt: ## Encrypt kubernates secret. ex) make encrypt secret=password
	@script/encrypt.sh ${secret}

.PHONY: decrypt
decrypt: ## Decrypt kubernates secret. ex) make decrypt secret=password
	@script/decrypt.sh ${secret}

.PHONY: sops
sops: ## sops kubernates secret. ex) make sops secret=password
	@sops --decrypt --in-place k8s/secret/${secret}.yaml

.PHONY: mongo
mongo: ## mongo
	@go run cmd/mongo/main.go

.PHONY: linetest
linetest: ## linetest
	@go run cmd/line/main.go

.PHONY: runtest
runtest: ## run test
	@go run cmd/app/main.go
