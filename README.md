# grpc-pcbook
 grpc demo

#### Makefile命令
```
gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.
clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 127.0.0.1:8080

test:
	go test -cover -race ./...

```