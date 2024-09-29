.PHONY: server client

server:
	go fmt ./cmd/server/main.go
	go vet ./cmd/server/main.go
	go build -o ./bin/server ./cmd/server/main.go
	./bin/server

client: 
	go fmt ./cmd/client/main.go
	go vet ./cmd/client/main.go
	go build -o ./bin/client ./cmd/client/main.go
	./bin/client

clean:
	rm -r ./bin

