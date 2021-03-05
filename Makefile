version := "0.13.3"

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Makefile Commands:"
	@echo "----------------------------------------------------------------"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
	@echo "----------------------------------------------------------------"

run: ## run server
	@go run cmd/stategate/main.go

patch: ## bump sem version by 1 patch
	bumpversion patch --allow-dirty

minor: ## bump sem version by 1 minor
	bumpversion minor --allow-dirty

tag: ## tag the repo (remember to commit changes beforehand)
	git tag v$(version)

push: ## push tagged version to remote repository
	git push origin v$(version)

docker-build:
	@docker build -t stategate/stategate:v$(version) .

docker-push:
	@docker push stategate/stategate:v$(version)


.PHONY: proto
proto: ## regenerate gRPC code
	@echo "generating protobuf code..."
	@docker run -v `pwd`:/defs namely/prototool:latest generate
	@go fmt ./...

.PHONY: up
up: ## start local containers
	@docker-compose -f docker-compose.yml pull
	@docker-compose -f docker-compose.yml up -d

.PHONY: down
down: ## shuts down local docker containers
	@docker-compose -f docker-compose.yml down --remove-orphans

.PHONY: gql
gql: ## regenerate graphql code
	@gqlgen generate

build: ## build the server to ./bin
	@mkdir -p bin
	@cd cmd/stategate; gox -osarch="linux/amd64" -output="../../bin/linux/{{.Dir}}_linux_amd64"
	@cd cmd/stategate; gox -osarch="darwin/amd64" -output="../../bin/darwin/{{.Dir}}_darwin_amd64"
	@cd cmd/stategate; gox -osarch="windows/amd64" -output="../../bin/windows/{{.Dir}}_windows_amd64"

test: ## run tests
	@go test -v ./...


chart:
	@cd chart; helm dependency update && helm package .
	@helm repo index .

update-helm:
	@helm repo add --username stategate https://raw.githubusercontent.com/stategate/stategate/master/
	@helm repo update