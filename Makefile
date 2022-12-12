test:
	@go test -v ./...

build_server:
	@go build -o bin/server ./cmd/server

build_client:
	@go build -o bin/client ./cmd/client

run_server: build_server
	@bin/server

run: build_client
	bin/client --addr localhost:8080 --method set --key foo --value bar
	bin/client --addr localhost:8080 --method get --key foo
	bin/client --addr localhost:8080 --method delete --key foo

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        pkg/proto/service.proto