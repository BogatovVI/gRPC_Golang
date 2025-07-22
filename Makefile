.PHONY: proto

PROTO_DIR := proto
GEN_DIR := pb

PATH_PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")
PROTO_FILES := ${shell basename ${PATH_PROTO_FILES}}

GEN_FILES := $(patsubst %.proto, $(GEN_DIR)/%.pb.go, $(PROTO_FILES))

proto: $(GEN_FILES)

$(GEN_DIR)/%.pb.go: ${PROTO_DIR}/%.proto 
	@mkdir -p $(@D)
	protoc \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$<
