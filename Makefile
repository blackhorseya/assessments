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
