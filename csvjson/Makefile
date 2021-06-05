APP_NAME = glasnostic
VERSION = latest

.PHONY: clean
clean:
	@rm -rf bin profile.out coverage.txt

.PHONY: test
test:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc:./internal/app/entities \
	--proto_path=./internal/app/protos \
	./internal/app/protos/*.proto

.PHONY: build-image
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" .

.PHONY: tag-image
tag-image:
	@docker tag $(APP_NAME):$(VERSION) blackhorseya/$(APP_NAME):$(VERSION)

.PHONY: push-iamge
push-image:
	@docker push blackhorseya/$(APP_NAME):$(VERSION)
