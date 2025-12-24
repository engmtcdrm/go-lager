.PHONY: build runexe run rundebug runtrace

build:
	@echo "Size before build:"; \
	ls -la examples |grep examples; \
	ls -lh examples |grep examples; \
	echo "\n\nSize after build:"; \
	go build --ldflags "-s -w" -o examples/examples ./examples; \
	ls -la examples |grep examples; \
	ls -lh examples |grep examples

runexe:
	@./examples/examples

run:
	@go run ./examples/main.go

rundebug:
	@go run ./examples/main.go -d

runtrace:
	@go run ./examples/main.go -t
