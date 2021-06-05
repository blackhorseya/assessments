APP_NAME = ip-rate-limit
VERSION = latest
PROJECT_ID = sean-side
NS = side
DEPLOY_TO = uat

.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: test
test:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/ip-rate-limit'

.PHONY: build-image
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	-f Dockerfile .

.PHONY: tag-image
tag-image:
	@docker tag $(APP_NAME):$(VERSION) gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: push-image
push-image:
	@docker push gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: run-with-docker
run-with-docker: build-image
	@docker run -it --rm -p 8888:8080 -v "$(shell pwd)/configs/app.yaml:/app/configs/app.yaml" $(APP_NAME):$(VERSION)

.PHONY: run-with-local
run-with-local:
	@go run ./cmd/app

.PHONY: gen
gen: gen-wire gen-swagger

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/app/main.go --parseInternal -o internal/app/apis/docs

.PHONY: deploy
deploy:
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) ./deployments/$(APP_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)
