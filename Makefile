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
