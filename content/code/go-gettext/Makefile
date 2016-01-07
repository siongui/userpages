# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)


test:
	@# -v means verbose, can see logs of t.Log
	@go test -v

help:
	go help

install:
	go get github.com/chai2010/gettext-go/gettext
	@#go get github.com/samuel/go-gettext/gettext

msgfmt:
	msgfmt locale/zh_TW/LC_MESSAGES/messages.po -o locale/zh_TW/LC_MESSAGES/messages.mo
	msgfmt locale/vi_VN/LC_MESSAGES/messages.po -o locale/vi_VN/LC_MESSAGES/messages.mo

clean:
	rm -rf pkg src
	rm locale/zh_TW/LC_MESSAGES/messages.mo
	rm locale/vi_VN/LC_MESSAGES/messages.mo
