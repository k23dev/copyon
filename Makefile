BINARY_NAME=goshareit
BUILD_DIR=./build

all: build

deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go mod tidy

build:
	# create directories
	mkdir -p ${BUILD_DIR}
	# templates generator
	templ generate

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -ldflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

templates:
	templ generate

dev:
	templ generate
	air

run:
	templ generate
	go run .

test:
	go test ./tests

clean:
	go clean
	rm -rf ${BUILD_DIR}