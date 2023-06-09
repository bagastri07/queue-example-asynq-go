.PHONY: mock
mock: ; $(info $(M) generating mock...) @
	@./script/mockgen.sh

.PHONY: test
test: ; $(info $(M) start unit testing...) @
	@go test $$(go list ./... | grep -v /mocks/) --race -v -short -coverprofile=profile.cov
	@echo "\n*****************************"
	@echo "**  TOTAL COVERAGE: $$(go tool cover -func profile.cov | grep total | grep -Eo '[0-9]+\.[0-9]+')%  **"
	@echo "*****************************\n"