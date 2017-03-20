export PATH := $(PATH):$(realpath ../../../../../go_appengine/)
PROJECT_DIR=$(CURDIR)
PROJECT_ID=golden-operator-130720
PROJECT_VERSION=blogexample

default:
	@echo "\033[92mRun development web server ...\033[0m"
	@goapp serve

fmt:
	@echo "\033[92mGo fmt source code ...\033[0m"
	@goapp fmt *.go

deploy:
	goapp deploy -application ${PROJECT_ID} -version ${PROJECT_VERSION} ${PROJECT_DIR}
	@echo "\033[92mDeployed URL: http://${PROJECT_VERSION}.${PROJECT_ID}.appspot.com/\033[0m"
