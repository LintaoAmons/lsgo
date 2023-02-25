# 设置要编译的目标平台和体系结构
TARGETS = darwin/amd64 darwin/arm64 linux/amd64 windows/amd64

OUTPUT_DIR = artifact

BINARY_NAME = lsgo

all: clean $(TARGETS)

$(TARGETS):
	@echo "Building $@ ..."
	GOOS=$$(echo $@ | cut -d/ -f1) GOARCH=$$(echo $@ | cut -d/ -f2) go build -o $(OUTPUT_DIR)/$(BINARY_NAME)_$$(echo $@ | tr / _) .

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: all clean $(TARGETS)
