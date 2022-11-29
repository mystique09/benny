clean:
	rm -rf ./tmp coverage.out

security:
	gosec -quiet -exclude-generated ./...

critic:
	gocritic check -enableAll ./...

lint: security critic
	golangci-lint run ./...

test: clean
	go test -v -cover -coverprofile=coverage.out ./...
	
coverage:
	go tool cover -html=coverage.out

run:
	go run cmd/main.go
	
build:
	rm -rf bin/benny
	go build -o bin/benny cmd/main.go

.PHONY: clean security critic lint test server coverage run build