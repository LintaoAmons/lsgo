# 设置要编译的目标平台和体系结构
TARGETS = darwin/arm64 linux/amd64 windows/amd64

OUTPUT_DIR = artifact

BINARY_NAME = lsgo

all: clean $(TARGETS)

$(TARGETS):
	@echo "Building $@ ..."
	GOOS=$$(echo $@ | cut -d/ -f1) GOARCH=$$(echo $@ | cut -d/ -f2) go build -o $(OUTPUT_DIR)/$(BINARY_NAME)_$$(echo $@ | tr / _) .

clean:
	rm -rf $(OUTPUT_DIR)

build:
	go build -o ~/.local/bin/lsgo main.go
	chmod u+x ~/.local/bin/lsgo

.PHONY: all clean $(TARGETS)
