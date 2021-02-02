version := "0.0.4"

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Makefile Commands:"
	@echo "----------------------------------------------------------------"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	@echo "----------------------------------------------------------------"

run:
	@go run cmd/eventgate/main.go

gen: gql helm proto

helm:
	cd chart; helm package .
	helm repo index .

patch: ## bump sem version by 1 patch
	bumpversion patch --allow-dirty

minor: ## bump sem version by 1 minor
	bumpversion minor --allow-dirty

tag: ## tag the repo (remember to commit changes beforehand)
	git tag v$(version)

push:
	git push origin v$(version)

docker-build:
	@docker build -t colemanword/eventgate:v$(version) .

docker-push:
	@docker push colemanword/eventgate:v$(version)


.PHONY: proto
proto: ## regenerate gRPC code
	@echo "generating protobuf code..."
	@docker run -v `pwd`:/defs namely/prototool:latest generate
	@go fmt ./...

.PHONY: gql
gql: ## regenerate graphql code
	@gqlgen generate

.PHONY: up
up: ## start local containers
	@docker-compose -f docker-compose.yml pull
	@docker-compose -f docker-compose.yml up -d

.PHONY: down
down: ## shuts down local docker containers
	@docker-compose -f docker-compose.yml down --remove-orphans

build: ## build the server to ./bin
	@mkdir -p bin
	@cd cmd/eventgate; gox -osarch="linux/amd64" -output="../../bin/linux/{{.Dir}}_linux_amd64"
	@cd cmd/eventgate; gox -osarch="darwin/amd64" -output="../../bin/darwin/{{.Dir}}_darwin_amd64"
	@cd cmd/eventgate; gox -osarch="windows/amd64" -output="../../bin/windows/{{.Dir}}_windows_amd64"
