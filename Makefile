PROTO = ocher.proto

PB_GO_DIR = internal/pb
PB_GO = $(PB_GO_DIR)/ocher.pb.go $(PB_GO_DIR)/ocher_grpc.pb.go

PB_PY_DIR = python/ocher/pb/
PB_PY = $(PB_PY_DIR)/ocher_pb2.py $(PB_PY_DIR)/ocher_pb2_grpc.py

all: $(PB_GO) $(PB_PY)

$(PB_GO) &:: $(PROTO)
	mkdir -p $(PB_GO_DIR)
	protoc -I. --go_out=paths=source_relative:$(PB_GO_DIR) --go-grpc_out=$(PB_GO_DIR) --go-grpc_opt=paths=source_relative $^

$(PB_PY) &:: $(PROTO)
	python scripts/generate_pb.py $^
