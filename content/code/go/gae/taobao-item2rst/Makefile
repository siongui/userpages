export PATH := $(PATH):$(realpath ../../../../go_appengine/)
PROJECT_DIR=$(CURDIR)
PROJECT_ID=golden-operator-130720
PROJECT_VERSION=taobao-item2rst

default:
	@echo "\033[92mRun development web server ...\033[0m"
	@cd ../; goapp serve ${PROJECT_DIR}

fmt:
	@echo "\033[92mGo fmt source code ...\033[0m"
	@goapp fmt *.go

deploy:
	cd ../; appcfg.py -A ${PROJECT_ID} -V ${PROJECT_VERSION} update ${PROJECT_DIR}
	@echo "\033[92mDeployed URL: http://${PROJECT_VERSION}.${PROJECT_ID}.appspot.com/\033[0m"

install:
	@echo "\033[92mInstall golang.org/x/net/html ...\033[0m"
	@goapp get -u golang.org/x/net/html
	@echo "\033[92mInstall google.golang.org/appengine ...\033[0m"
	@goapp get -u google.golang.org/appengine
