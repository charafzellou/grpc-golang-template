install:
	@sudo apt update
	@sudo apt install protobuf-compiler -y
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@export PATH="${PATH}:$(go env GOPATH)/bin"

gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/main.proto

run:
	@go run server/main.go
	@go run client/main.go