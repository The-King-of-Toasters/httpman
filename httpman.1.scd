HTTPMAN(1)
# NAME
httpman - manpages for HTTP status codes
# SYNOPSIS
httpman [*-ahjs*] [code ...]
# DESCRIPTION
httpman provides information about HTTP status codes in roff, *scdoc*(1)
or JSON format.
It is meant to be piped into *man*(1), *nroff*(1), *groff*(1),
*scdoc*(1) or *jq*(1) to be readable.

All codes passed are printed in numerical order.
# OPTIONS
*-a*	prints information for every code

*-h*	print usage

*-j*	output information in JSON format

*-s*	output information in *scdoc*(1) format
# INFO
Up-to-date versions of *httpman* and this man page can be found at
the upsteam url, 
_https://git.sgregoratto.me/The-King-of-Toasters/httpman_. 
A mirror on github is also available at
_https://github.com/The-King-of-Toasters/httpman_
# SEE ALSO
*groff*(7), *scdoc*(1)
# BUGS/FEATURES
If you believe you have found a bug in *httpman*
or you would like to add a feature, 
then you can submit reports, pull-requests and patches at
either the upstream url or on github.
# AUTHORS
Stephen Gregoratto _<themanhimself@sgregoratto.me>_
