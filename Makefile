OUT_DIR=./internal/pb
PROTO_MODULE=github.com/not-hype-pruduction/bridge-protos

generate:
	PROTO_PATH=$$(go list -m -f '{{.Dir}}' $(PROTO_MODULE))/proto; \
	protoc \
		-I $$PROTO_PATH \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$$PROTO_PATH/cards/v1/cards.proto \
		$$PROTO_PATH/biding/v1/biding.proto \
		$$PROTO_PATH/dds/v1/dds.proto

build-dds-macos:
	cd internal/infrastructure/dds/dds && make -f Makefile_MACOS

build-dds:
	cd internal/infrastructure/dds/dds && make -f Makefile_linux_static

