tidy:
	@go mod tidy
	@go mod vendor

run:
	@go run cmd/main.go

migration:
	@migration create -ext sql -dir ./migrations -seq &(name)