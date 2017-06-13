[Makefile] Rules for Linters of Golang Source Code
##################################################

:date: 2017-05-10T16:37+08:00
:tags: Bash, Makefile, Commandline, Go, Golang
:category: Bash
:summary: Example of Makefile rules for Go linters.
:og_image: https://image.slidesharecdn.com/writemicroserviceusinggolang-160821082419/95/write-microservice-in-golang-79-638.jpg

Example [1]_ of Makefile_ rules for Go linters. Make it a post for search engine
friendliness.

.. raw:: html

  <div class="reddit-embed" data-embed-media="www.redditmedia.com" data-embed-parent="false" data-embed-live="false" data-embed-uuid="10d5c777-cbb9-4ae5-a4c3-76127d21d247" data-embed-created="2017-05-10T08:25:21.134Z"><a href="https://www.reddit.com/r/golang/comments/621gk6/linters_do_not_get_attention_they_deserve/dfjub71/">Comment</a> from discussion <a href="https://www.reddit.com/r/golang/comments/621gk6/linters_do_not_get_attention_they_deserve/">Linters do not get attention they deserve</a>.</div><script async src="https://www.redditstatic.com/comment-embed.js"></script>

.. code-block:: makefile

  PACKAGES=$(shell go list ./... | grep -v /vendor)

  check:  ## Check the code using various linters and static checkers.
  	@echo "Running gofmt..."
  	@gofmt -d $(shell find . -type f -name '*.go' -not -path "./vendor/*")

  	@echo "Running go vet..."
  	@for package in $(PACKAGES);  do \
  		go vet -v $$package || exit 1; \
  	done

  	@echo "Running golint..."
  	@for package in $(PACKAGES); do \
  		golint -set_exit_status $$package || exit 1; \
  	done

  	@echo "Running errcheck..."
  	@for package in $(PACKAGES); do \
  		errcheck -ignore 'Close' -ignoretests $$package || exit 1; \
  	done

References:

.. [1] `Orange_Tux comments on Linters do not get attention they deserve <https://www.reddit.com/r/golang/comments/621gk6/linters_do_not_get_attention_they_deserve/dfjub71/>`_
.. [2] `Linters do not get attention they deserve : golang <https://www.reddit.com/r/golang/comments/621gk6/linters_do_not_get_attention_they_deserve/>`_

.. _Makefile: https://www.google.com/search?q=Makefile
