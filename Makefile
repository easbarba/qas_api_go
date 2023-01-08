.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64
NAME := qas
TO := ${HOME}/.local/bin

deps:
	go mod download


build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -o ${NAME} ./main.go

lint:
	golint ./...

vet:
	go vet ./...

run:
	go run main.go

test:
	go test ./...

imports:
	goimports -l -w .

shell:
	guix shell --pure --container

watch:
	CompileDaemon --build="go build -o ./${NAME} ." --command="./${NAME}"
