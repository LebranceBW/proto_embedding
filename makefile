.PHONY: pb

SRC_DIR=./proto
DST_DIR=.

.all: pb

pb:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/example.proto