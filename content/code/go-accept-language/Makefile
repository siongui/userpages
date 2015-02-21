# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)


test:
	@# -test.v means verbose, can see logs of t.Log
	@go test -test.v

parse:
	@# http://stackoverflow.com/questions/9967105/suppress-echo-of-command-invocation-in-makefile
	@go run parse.go

help:
	go help
