PROTO = ocher.proto
OUT_GO = pb/*.go
OUT_PY = python/ocher/pb/*pb2.py

all: $(OUT_GO) $(OUT_PY)

$(OUT_GO): $(PROTO)
	mkdir -p pb
	protoc -I. --go_out=plugins=grpc,paths=source_relative:pb $^

$(OUT_PY): $(PROTO)
	python scripts/generate_pb.py $^
