.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64
NAME := qas
BIN := ./main.go
deps:
	go mod download

build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -o ${NAME} ${BIN}

lint:
	golint ./...

vet:
	go vet ./...

run:
	go run ${BIN}

test:
	go test ./...

imports:
	goimports -l -w .

shell:
	guix shell --pure --container

watch:
	CompileDaemon --build="go build -o ./${NAME} ${BIN}" --command="./${NAME}"
