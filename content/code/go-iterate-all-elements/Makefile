export GOROOT=$(realpath ../../../../go)
export GOPATH=$(realpath ../../../../nanda/tool/crmn/)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


default:
	@echo "\033[92mIterating HTML and extracting links to rst ...\033[0m"
	@go run html.go handleHtmlLink.go -input=index.html

wget:
	@wget https://siongui.github.io/2016/03/19/go-download-html-from-url/

website:
	@cd ../../.. ; make

fmt:
	@echo "\033[92mGo fmt source code ...\033[0m"
	@go fmt *.go

install:
	@echo "\033[92mInstalling golang.org/x/net/html ...\033[0m"
	go get -u golang.org/x/net/html
