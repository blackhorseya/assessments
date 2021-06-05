APP_NAME = amazingtalker
VERSION = latest
PROJECT_ID = sean-side
NS = side
DEPLOY_TO = uat

.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint
lint:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/amazingtalker'

.PHONY: test-with-coverage
test-with-coverage:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	-f Dockerfile .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: tag-image
tag-image:
	@docker tag $(APP_NAME):$(VERSION) gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: push-image
push-image:
	@docker push gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: deploy
deploy:
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) ./deployments/helm/$(APP_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-wire gen-swagger gen-pb

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o internal/apis/docs

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./pb ./pb/*.proto

.PHONY: install-db
install-db:
	@helm --namespace $(NS) upgrade --install $(APP_NAME)-db bitnami/mariadb \
	--values ./deployments/configs/$(DEPLOY_TO)/db.yaml
