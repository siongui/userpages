# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

all: example3

example1:
	@# http://stackoverflow.com/questions/9967105/suppress-echo-of-command-invocation-in-makefile
	@go run parse-1.go

example2:
	@go run parse-2.go

example3:
	@go run parse-3.go

help:
	go help
