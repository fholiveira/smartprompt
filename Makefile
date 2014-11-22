GET=go get

BIN=$(GOPATH)/bin/smartprompt
GIT2GO=$(GOPATH)/src/github.com/libgit2/git2go

all: test

install:
	@$(GET) 'github.com/docopt/docopt-go'
	@$(GET) -d 'github.com/libgit2/git2go'
	@cd $(GIT2GO) && git submodule update --init 
	@cd $(GIT2GO) && $(MAKE) install

install_tests:
	@$(GET) 'github.com/stretchr/testify'
	@$(GET) 'github.com/stretchr/gorc'

clean:
	@rm -rf $(BIN)

test:
	@$(GOPATH)/bin/gorc
