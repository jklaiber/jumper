BINARY_NAME=jumper

build:
	GOARCH=amd64 GOOS=linux go build -o ./out/${BINARY_NAME} main.go

build-arm-darwin:
	GOARCH=arm64 GOOS=darwin go build -o ./out/${BINARY_NAME} main.go

run:
	./out/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ./out/${BINARY_NAME}