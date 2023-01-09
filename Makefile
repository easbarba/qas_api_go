.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64
NAME := qas
MAIN := ./main.go

deps:
	go mod download
	go mod verify


build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -o ${NAME} ${MAIN}

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
