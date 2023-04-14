SHELL := /bin/sh
BIN := /usr/local/bin

PROTO_LOCATION := $(shell find proto -iname "proto" -exec echo "-I="{} \;)

.PHONY: clean proto

clean:
	rm -rf gen

## Compile all proto files with buf
proto: clean
		mkdir -p gen # create the empty directory for proto targets.
		buf generate --verbose .

## Compile all proto files with protoc - this is currently not in use because we have buf, but leaving for knowledge purpose
proto-protoc: clean
	mkdir -p gen # create the empty directory for proto targets.

        # Explanation of flags:
                # --proto_path specifies all the directories to search for "import" files inside the proto files
                # --go_grpc_out specifies where to put the generated files
                # --go_out specifies where to put the generated files
                # The last arg is the input directory of which proto files to build
                # Note: we have to build the message and service proto files separately,
                # because of the required gRPC plugin.
	protoc $(PROTO_LOCATION) \
		--go_out=paths=source_relative:./gen/ --go-grpc_out=paths=source_relative:./gen/ \
		proto/*.proto

	protoc $(PROTO_LOCATION) \
        --grpc-gateway_out=paths=source_relative:./gen/ \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt generate_unbound_methods=true \
        --grpc-gateway_opt paths=source_relative \
        proto/*.proto
