all: run

.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o bin/timetracerpt main.go

.PHONY: install
install:
	@go install

clean:
	rm -Rf bin/*
