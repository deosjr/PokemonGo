singleplayer:
	go run src/main.go

server:
	go run src/main.go -serve

test:
	go test ./...
