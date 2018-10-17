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

Licence
-------

3 terms BSD licence, ©2018 Mathieu Lecarme