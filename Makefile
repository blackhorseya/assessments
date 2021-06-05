.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint
lint:
	@golint -set_exit_status ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/assessments'

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: gen
gen: gen-wire gen-pb gen-swagger

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./internal/pkg/entity/**/*.proto

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs