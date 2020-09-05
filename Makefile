compile:
	echo "Compiling for linux"
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/gotee ./cmd/gotee

all: compile