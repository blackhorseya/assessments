.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint
lint:
	@golint -set_exit_status ./...

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh
