build:
	mkdir -p ./bin
	go build -o bin/main src/*.go

run:
	./bin/main

clean:
	rm -rf ./bin

