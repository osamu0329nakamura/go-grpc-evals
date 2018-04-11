all: proto server client

proto:
	./bin/protoc --proto_path=. --go_out=plugins=grpc:./src/rpc *.proto

server: server.go
	go build server.go

client: client.go
	go build client.go
