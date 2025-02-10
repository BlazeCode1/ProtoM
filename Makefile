PROTO_SRC=./proto
PROTO_GEN=./proto/gen

.PHONY: proto tidy
proto:
	protoc --go_out=$(PROTO_GEN) --go-grpc_out=$(PROTO_GEN) $(PROTO_SRC)/*.proto


tidy:
	go mod tidy
