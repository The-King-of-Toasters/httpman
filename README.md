# httpman

[![asciicast](https://asciinema.org/a/191692.png)](https://asciinema.org/a/191692)
[Link for gitea users](https://asciinema.org/a/191692) 

`httpman` is your manpage for HTTP. It prints info about 
HTTP status codes in troff markup, and can be piped into (n|g)roff for viewing.

**NEW:** Output formats also include JSON and
[scdoc](https://git.sr.ht/~sircmpwn/scdoc/)

## General use

```sh
httpman 404 [200 100 ...] | nroff -man
httpman -a | nroff -man
httpman -aj | jq
httpman 404 | scdoc > 404.7
```

## Compilation and Installation

```sh
go get -d git.sgregoratto.me/The-King-of-Toasters/httpman
cd $GOPATH/src/git.sgregoratto.me/The-King-of-Toasters/httpman
# Apply YAMLFILE.patch if needed
make
sudo make install
```

## Credits

The idea for httpman was stolen from [statcode](https://github.com/shobrook/statcode).
I was trying to immerse myself into go after reading 
[The Go Programming Language](https://www.gopl.io) and thought a simple program
like httpman help me with that.
