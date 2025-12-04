.PHONY: swagger run build test cover

swagger:
	GOBIN=$(CURDIR)/bin go install github.com/go-swagger/go-swagger/cmd/swagger@v0.33.1
	./bin/swagger generate spec -o ./docs/swagger.yaml --scan-models

run:
	@go run cmd/api/main.go

build:
	@go build -o bin/scoreapp cmd/api/main.go

test:
	@go test -v ./...

cover: cover-profile cover-html

cover-profile:
	@go test -v -coverprofile cover.out ./...

cover-html:
	@go tool cover -html=cover.out -o cover.html