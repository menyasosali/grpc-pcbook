gen:
	protoc --proto_path==proto/ proto/*.proto --go_out=./ --go-grpc_out=./

clean:
	del pb\*.go

test:
	go test -cover -race ./...