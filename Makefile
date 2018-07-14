CFG_PREFIX = /etc
PREFIX = /usr/bin
YAMLFILE = httpman_codes.yml

all: build

clean:
	rm -rf httpman

build:
	go build httpman.go

install:
	install -Dm755 httpman $(PREFIX)/httpman
	install -Dm664 $(YAMLFILE) $(CFG_PREFIX)/$(YAMLFILE)

.PHONY: clean
