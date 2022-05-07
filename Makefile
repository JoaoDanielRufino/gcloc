.PHONY:build
build:
	@echo "building the gcloc executable"
	go build -o bin/gcloc cmd/main.go
	@echo "executable created bin/gcloc"

.PHONY:clean
clean:
	@echo "cleaning the gcloc package"
	rm -rf bin coverage.out coverage.html

.PHONY:unit-test
unit-test:
	@echo "running unit tests"
	go test -v `go list ./... | grep -v test/`

.PHONY:coverage
coverage:
	@echo "coverage report"
	go test -cover -v -timeout 60s -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
