REPO_NAME=github.com/echoturing/log



.PHONY: fmt
fmt:
	@find . -name "*.go" | xargs goimports -w -l --local $(REPO_NAME) --private "mockprivate"


.PHONY: test
test:
	@go test