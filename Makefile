APP_NAME=oidc-copycat
BUILD_DIR=build

.DEFAULT_GOAL = run

build: 
	go build -o ./${BUILD_DIR}/${APP_NAME} *.go

run: build
	./${BUILD_DIR}/${APP_NAME}

clean:
	go clean
	rm ./${BUILD_DIR}/${APP_NAME}

test:
	go test ./...

.PHONY: build clean run test