# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
FILE=testsass
SCSSDIR=src/style


default:
	@go run $(FILE).go

fmt:
	@go fmt $(FILE).go

download:
	@[ -d $(SCSSDIR) ] || mkdir -p $(SCSSDIR)
	@[ -e $(SCSSDIR)/_normalize302.scss ] || wget -O $(SCSSDIR)/_normalize302.scss http://necolas.github.com/normalize.css/3.0.2/normalize.css
	@[ -e $(SCSSDIR)/_css4rst.scss ] || wget -O $(SCSSDIR)/_css4rst.scss https://github.com/siongui/userpages/raw/master/theme/styling/_css4rst.scss
	@[ -e $(SCSSDIR)/_pygments.scss ] || wget -O $(SCSSDIR)/_pygments.scss https://github.com/siongui/userpages/raw/master/theme/styling/_pygments.scss
	@[ -e $(SCSSDIR)/style.scss ] || wget -O $(SCSSDIR)/style.scss https://github.com/siongui/userpages/raw/master/theme/styling/style.scss

install:
	go get -u github.com/wellington/go-libsass
