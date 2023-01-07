NAME := qas

deps:
	go mod download

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

build:
	go build -o ${NAME} .

watch:
	CompileDaemon --build="go build -o ./${NAME} ." --command="./${NAME}"
