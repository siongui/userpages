# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)


test:
	@# -v means verbose, can see logs of t.Log
	@go test -v po2json.go po2json_test.go

run:
	go run html.go

help:
	go help

fmt:
	@#go fmt gettext.go gettext_test.go html.go
	go fmt po2json.go po2json_test.go

install:
	go get github.com/chai2010/gettext-go/gettext
	@#go get github.com/samuel/go-gettext/gettext
	@#go get github.com/gosexy/gettext

xgettext:
	sed "s/{{gettext \(".*"\)}}/{{gettext(\1)}}/g" html.go | xgettext --no-wrap --language=c --from-code=UTF-8 --output=locale/messages.pot -

msgfmt:
	msgfmt locale/zh_TW/LC_MESSAGES/messages.po -o locale/zh_TW/LC_MESSAGES/messages.mo
	msgfmt locale/vi_VN/LC_MESSAGES/messages.po -o locale/vi_VN/LC_MESSAGES/messages.mo

clean:
	rm -rf pkg src
	rm locale/messages.pot
	rm locale/zh_TW/LC_MESSAGES/messages.mo
	rm locale/vi_VN/LC_MESSAGES/messages.mo
