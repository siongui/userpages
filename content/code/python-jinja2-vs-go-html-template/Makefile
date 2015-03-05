# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

# path of Google App Engine Python SDK
GAE_PY_SDK=../../../../../google_appengine

all: gonet

gonet:
	go run html-net.go

webapp2:
	$(GAE_PY_SDK)/dev_appserver.py .

webpy:
	python webpy-jinja2.py

example1:
	@# http://stackoverflow.com/questions/9967105/suppress-echo-of-command-invocation-in-makefile
	@python jinja2-example-1.py
	@go run html-template-example-1.go

help:
	go help

clean:
	rm *.pyc
