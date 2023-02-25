build:
	go build -o ~/.local/bin/lsgo main.go
	chmod u+x ~/.local/bin/lsgo

run:
	go run main.go
