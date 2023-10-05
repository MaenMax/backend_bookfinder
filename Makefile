GO := go
GOFLAGS := -v -p 1

default: book_finder

all: clean default

book_finder: bin/book_finder

bin/book_finder: src/main.go
	@echo "========== Compiling $@ =========="
	$(GO) build $(GOFLAGS) -o bin/book_finder ./src
	

deploy:
	tar -cvjf "book_finder`git describe --tags --abbrev=0`.tar.bz2" bin

VERSION_DOCKER := $(shell echo `git describe`.`date +"%Y%m%d%H%M%S"` > .version)
docker: ${VERSION_DOCKER}
	docker build . -t book_finder$$(git describe --abbrev=0)

clean:
	@echo "Deleting generated binary files ..."; sh -c 'if [ -d bin ]; then  find bin/ -type f -exec rm {} \; -print ; fi; rm -Rf bin'
	@echo "Deleting generated archive files ..."; sh -c 'if [ -d pkg ]; then  find pkg -type f -name \*.a -exec rm {} \; -print ; fi;  rm -Rf pkg'
	@echo "Deleting emacs backup files ..."; find . -type f -name \*~ -exec rm {} \; -print
	@echo "Deleting log files ..."; find . -maxdepth 1 -type f \( -name \*.log.\* -o -name \*.log \) -exec rm {} \; -print
