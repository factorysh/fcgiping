Fcgiping
========

Debug fastcgi things.

Build
-----

    make

Usage
-----

    docker run -ti --rm -p 9000:9000 -u 33 bearstech/php:7.2

In another terminal :

    ./bin/fcgiping fcgi://127.0.0.1:9000/__status

Thanks
------

The hard job is done by [Caddy's fastcgi client](https://github.com/mholt/caddy/caddy/caddyhttp/fastcgi/fcgiclient.go)
(Copyright 2015 Light Code Labs, LLC)

Licence
-------

3 terms BSD licence, ©2018 Mathieu Lecarme