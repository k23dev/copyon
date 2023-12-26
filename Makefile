BINARY_NAME=copyon
VERSION=1.0
BUILD_DIR=./build

BINARY_NAME_WIN=copyon.exe
BUILD_DIR_WIN=./build/windows

all: build

deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go mod tidy

build-linux:
	# # create directories
	# mkdir -p ${BUILD_DIR}
	# # templates generator
	# templ generate

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -ldflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

build-win:
	# create directories
	mkdir -p ${BUILD_DIR_WIN}
	# templates generator
	templ generate

	# compile into binary file
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME_WIN} -ldflags "-w -s"
	mv ${BINARY_NAME_WIN} ${BUILD_DIR_WIN}

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