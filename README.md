# webdavd

---------

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

User mode WebDAV server written on Go

## License

----------

BSD-3-Clause

## Uses

-------

* log/slog
* golang.org/x/net/webdav
* github.com/go-chi/chi/v5 
* github.com/ilyakaznacheev/cleanenv 
* github.com/fatih/color (for color slog)

## Platforms

------------

Linux is webdavd' native platform.

## Compile

----------

Install Go 1.13 or higher:

* Debian/Ubuntu: `apt install golang`
* Fedora: `dnf install golang`

Then, download the source code and compile:

    $ git clone https://github.com/darkprof83/webdavd.git
    $ cd webdavd
    $ make

This will compile a static binary

## Use

------

1. Change default values

Edit config/local.yaml. Change salt, env ("prod", "dev", "local"), security ("none", "tls12") and key, cert, dir paths.

2. Generate and change hash for password
   
    $ ./cmd/wdhash/wdhash -config config/local.yaml

3. Run server
   
    $ ./cmd/webdavd/webdavd -config config/local.yaml
