OUT_DIR=./internal/pb
PROTO_MODULE=github.com/not-hype-pruduction/bridge-protos

generate:
	PROTO_PATH=$$(go list -m -f '{{.Dir}}' $(PROTO_MODULE))/proto; \
	protoc \
		-I $$PROTO_PATH \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$$PROTO_PATH/cards/v1/cards.proto
