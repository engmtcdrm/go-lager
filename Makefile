.PHONY: menv test testv build runexe run rundebug runtrace

PARENT_DIR := $(notdir $(CURDIR))

menv:
	@echo "Current directory: $(CURDIR)"
	@echo "Parent directory name: $(PARENT_DIR)"

test:
	@go test ./...

testv:
	@go test -v ./...

build:
	@cd examples; \
	echo "Size before build:"; \
	ls -la |grep examples; \
	ls -lh |grep examples; \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w"; \
	strip examples; \
	ls -la |grep examples; \
	ls -lh |grep examples; \
	cd ..

runexe:
	@./examples/examples

run:
	@cd examples; \
	go run .; \
	cd ..
