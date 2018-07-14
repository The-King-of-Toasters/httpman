# httpman

[![asciicast](https://asciinema.org/a/191692.png)](https://asciinema.org/a/191692)

`httpman` is your manpage for HTTP. It prints info about 
HTTP status codes in troff markup, and can be piped into (n|g)roff for viewing.

## General use

```sh
httpman 404 [200 100 ...] | nroff -man

httpman -a | nroff -man
```