NAME := qas

watch:
	CompileDaemon --build="go build -o ./${NAME} ." --command="./${NAME}"
