run: ## Run RozgarLink poject on host machine
	go run cmd/main.go

test: ## Run all unit tests in the project
	go test -v ./...
