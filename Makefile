APP_NAME=vsps
VERSION=v0.2.2
BUILD_DIR=bin/$(VERSION)

CHECKVERSION=$(shell grep 'const version' cmd/root.go | sed 's/.*= "\(.*\)"/\1/')
check-version:
	@if [ "$(VERSION)" != "$(CHECKVERSION)" ]; then \
		echo "Version mismatch!"; \
		echo "    VERSION in Makefile: $(VERSION)"; \
		echo "    Version in cmd/root.go: $(CHECKVERSION)"; \
		exit 1; \
	else \
		echo "Version match: $(VERSION)"; \
	fi

.PHONY: build
build:
	@echo "Building $(APP_NAME) for $(OS)/$(ARCH) version $(VERSION)"
	GOOS=$(OS) GOARCH=$(ARCH) go build -o "$(BUILD_DIR)/$(APP_NAME)_$(OS)_$(ARCH)_$(VERSION)"

compress-tar:
	tar -czvf $(BUILD_DIR)/$(APP_NAME)_$(OS)_$(ARCH)_$(VERSION).tar.gz $(BUILD_DIR)/$(APP_NAME)_$(OS)_$(ARCH)_$(VERSION)

compress-zip:
	cd $(BUILD_DIR) && zip $(APP_NAME)_$(OS)_$(ARCH)_$(VERSION).zip $(APP_NAME)_$(OS)_$(ARCH)_$(VERSION)

darwin:
	$(MAKE) OS=darwin ARCH=amd64 build compress-tar

linux:
	$(MAKE) OS=linux ARCH=amd64 build compress-tar

windows:
	$(MAKE) OS=windows ARCH=amd64 build compress-zip

all: check-version clean darwin linux windows

clean:
	rm -rf $(BUILD_DIR)
