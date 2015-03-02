# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

all: atom2rss

example1:
	@# http://stackoverflow.com/questions/9967105/suppress-echo-of-command-invocation-in-makefile
	@go run parse-1.go

example2:
	@go run parse-2.go

example3:
	@go run parse-3.go

example4:
	@go run parse-4.go

example5:
	@go run parse-5.go

example5_2:
	@go run parse-5_2.go

example6:
	@go run parse-6.go

example7:
	@go run parse-7.go

atom2rss:
	@go run atom2rss.go

help:
	go help
