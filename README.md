<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>

[![Sourcegraph](https://sourcegraph.com/github.com/labstack/echo/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/labstack/echo?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/labstack/echo)
[![Go Report Card](https://goreportcard.com/badge/github.com/labstack/echo?style=flat-square)](https://goreportcard.com/report/github.com/labstack/echo)
[![Build Status](http://img.shields.io/travis/labstack/echo.svg?style=flat-square)](https://travis-ci.org/labstack/echo)
[![Codecov](https://img.shields.io/codecov/c/github/labstack/echo.svg?style=flat-square)](https://codecov.io/gh/labstack/echo)
[![Join the chat at https://gitter.im/labstack/echo](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/labstack/echo)
[![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://forum.labstack.com)
[![Twitter](https://img.shields.io/badge/twitter-@labstack-55acee.svg?style=flat-square)](https://twitter.com/labstack)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)

## Versi Go yang didukung
Pada versi 4.0.0, Echo tersedia sebagai modul [Go module](https://github.com/golang/go/wiki/Modules). Oleh karena itu diperlukan versi Go :
- 1.9.7+
- 1.10.3+
- 1.11+

Ini adalah repository sample kerangka projek (Bahan ), dengan menggunakan teknologi Bahasa Pemrogrman Go.  framework yang saya buat sudah dilengkapi dengan fitur otentifikasi dan Priviledges. Mohon di bantu dan evaluasi apabila masih ada kekurangan dalam kerangka projek ini.  Say yes to code :)

## Requirement
- Redis Server Cache - https://redis.io/
- Go - https://golang.org/doc/install
- Mysql - https://www.mysql.com/downloads/


## Cara Menjalankan
- Import database mysql dengan keterangan sebagai berikut
```go
const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "db_jamu_golang_auth"
    DB_USER = "root"
    DB_PASS = ""
)

```
- Run Redis Service
- Run Mysql Service
- Masuk ke .App/ kemudiankan jalankan melalui CLI ( go run main.go )
- Go Framework already serve


## Author Owner
- [Dikhi Martin](https://www.linkedin.com/in/dikhi-martin/)
- [Hajik Gustian Hidayat](https://www.linkedin.com/in/hajik-gustian-hidayat-6ab575162/)
- [Adam Arya Pratama](https://www.linkedin.com/in/adam-arya-pratama-76781a140/)

## License
[MIT](https://github.com/labstack/echo/blob/master/LICENSE)
