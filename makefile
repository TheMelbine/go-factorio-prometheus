
.PHONY: tests
tests:
	go test ./... --cover -coverprofile=coverage.out --covermode atomic --coverpkg=./...

.PHONY: coverage
coverage: tests
	go tool cover -html=coverage.out

docs:
	go run main.go gen-docs