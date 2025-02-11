PROTO_SRC=./proto
PROTO_GEN=./proto/gen
server =./main.go
client =./client/main.go
.PHONY: proto clean t server
	
proto:
	protoc --proto_path=$(PROTO_SRC) \
	       --go_out=${PROTO_GEN} --go_opt=paths=source_relative \
	       --go-grpc_out=${PROTO_GEN} --go-grpc_opt=paths=source_relative \
	       $(PROTO_SRC)/*.proto

clean:
	rm -rf $(PROTO_GEN)
t:
	go mod tidy
server:
	go run $(server)

client:
	go run $(client)
