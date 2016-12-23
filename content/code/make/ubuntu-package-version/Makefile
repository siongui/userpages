PACKAGE_NAME=python
PACKAGE_VERSION=$(strip $(shell apt-cache policy $(PACKAGE_NAME) | grep Installed: | cut -d: -f2))

default:
	@echo "package name: \033[92m$(PACKAGE_NAME)\033[0m"
	@echo "package version: \033[92m$(PACKAGE_VERSION)\033[0m"
