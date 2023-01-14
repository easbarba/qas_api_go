.DEFAULT_GOAL := build

NAME := qas_api
MAIN := ./main.go
DEST := ${HOME}/.local/bin

OS :=linux
ARCH := amd64

deps:
	go mod download
	go mod verify


build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -o ${NAME} ${MAIN}

install: build
	mv -v ${NAME} ${DEST}/${NAME}

lint:
	golint ./...

vet:
	go vet ./...

run:
	go run ${MAIN}

test:
	go test ./...

imports:
	goimports -l -w .

clean:
	rm ${NAME}
	go mod tidy


shell:
	guix shell --pure --container

watch:
	CompileDaemon --build="go build -o ./${NAME} ${MAIN}" --command="./${NAME}"

image:
	podman build --file ./Dockerfile --tag ${USER}/${NAME}:$(shell cat .version)
