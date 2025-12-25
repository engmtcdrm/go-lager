.PHONY: menv build runexe run rundebug runtrace

PARENT_DIR := $(notdir $(CURDIR))

menv:
	@echo "Current directory: $(CURDIR)"
	@echo "Parent directory name: $(PARENT_DIR)"

build:
	@cd example; \
	echo "Size before build:"; \
	ls -la |grep 'example'; \
	ls -lh |grep example; \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w"; \
	strip example; \
	ls -la |grep example; \
	ls -lh |grep example; \
	cd ..

runexe:
	@./example/example

run:
	@cd example; \
	go run main.go; \
	cd ..
