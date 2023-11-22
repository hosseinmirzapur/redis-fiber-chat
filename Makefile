build:
	go build -o bin/rethink cmd/main.go

run: build
	./bin/rethink