build: 
	@go build -o bin/LiveDb cmd/main.go

run:
	@./bin/LiveDb

test: 
	@go test -v ./...
	