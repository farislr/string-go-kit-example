run:
	go run main.go

build:
	go build -o bin/core main.go

dev:
	CompileDaemon -build="go build -o bin/core main.go" -command="bin/core"

