# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


build: fmt
	@echo "\033[92mCompiling Go to JavaScript ...\033[0m"
	gopherjs build extension/click.go -o extension/click.js
	gopherjs build extension/cc.go -o extension/cc.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt extension/*.go

install:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs
	@echo "\033[92mInstalling GopherJS Bindings for Chrome ...\033[0m"
	go get -u github.com/fabioberger/chrome
	@echo "\033[92mInstalling github.com/siongui/godom ...\033[0m"
	go get -u github.com/siongui/godom

clean:
	rm extension/*.js
	rm extension/*.js.map
