
jock:
	go build jock.go

plugin:
	go build -o example ./example

pb:
	protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      proto/kv.proto

all:
	protoc --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  proto/jock.proto && \
	go build jock.go && go build -o example ./example