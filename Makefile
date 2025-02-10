PROTO_SRC=./proto
PROTO_GEN=./proto/gen

.PHONY: proto
proto: $(PROTO_SRC)/*.proto
	mkdir -p $(PROTO_GEN)
	protoc --go_out=$(PROTO_GEN) --go-grpc_out=$(PROTO_GEN) $(PROTO_SRC)/*.proto


clean:
	rm -rf $(PROTO_GEN)
t:
	go mod tidy
server:
	go run main.go
