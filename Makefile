compile:
	echo "Compiling for linux"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/gotee ./cmd/gotee

all: compile