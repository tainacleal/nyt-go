build:
	go build -o bin/example examples/main.go

run: build
	./bin/example