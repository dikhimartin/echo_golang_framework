# GOLANG | Echo Framework v.2
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

## Opsi 
### 1. Docker (portablize)

#### Intro 
Docker adalah sebuah platform open source untuk menyatukan file-file yang dibutuhkan sebuah software sehingga menjadi menjadi satu kesatuan yang lengkap dan berfungsi secara portabel dan bersifat virtual server. 

#### Requirement install
- Docker Engine
https://docs.docker.com/engine/install/
- Docker Compose
https://docs.docker.com/compose/install/

#### Cara Menjalankan 
``` shell
cp app/.env.example app/.env
cd app/
docker-compose up -d
docker-compose exec go bash
```
```shell
go run main.go
```

#### Server Running
```shell
http://localhost:8888/
```

#### Stop Server
```shell
docker-compose down
```
#### Manage Redis Cache
```shell
docker-compose exec redis redis-cli
```

#### Manage Mysql Database
```shell
docker-compose exec mysql sh 
mysql -u user -p
Enter password: user
```

#### Sumary

### 2. Bash (manual native) 

#### Requirement install
- GO            => go1.15.1
- Redis Server  => 4.0.9)
- Mysql         => (mysql  Ver 14.14 Distrib 5.7.30, for Linux)


#### Cara Menjalankan 
 ``` shell
 cd app/
 go run main.go 
 ```
 
#### Server Running
```shell
http://localhost:8888/
```




## Author Owner
- [Dikhi Martin](https://www.linkedin.com/in/dikhi-martin/)

## License
[MIT](https://github.com/labstack/echo/blob/master/LICENSE)

