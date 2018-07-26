PREFIX = /usr/bin
MAN1 = /usr/share/man/man1
LDFLAGS = -s -w

all: build

clean:
	rm -rf httpman

build:
	go build -ldflags '$(LDFLAGS)'

man:
	scdoc < httpman.1.scd > httpman.1

install:
	install -Dm755 httpman $(PREFIX)/httpman
	install -Dm644 httpman.1 $(MAN1)/httpman.1

.PHONY: clean

.SILENT: build clean man
