# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)


test:
	@# -v means verbose, can see logs of t.Log
	@go test -v

fmt:
	go fmt *.go

help:
	go help
