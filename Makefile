.PHONY: generate run test

generate:
	oapi-codegen -config oapi-codegen.yaml openapi.yaml

run:
	go run ./cmd/airthings_exporter

test:
	go test ./...
