# go-web-echo


Package echo implements high performance, minimalist Go web framework.

"Echo" is a web application framework for Go, commonly used for building HTTP services and APIs. It provides features to handle routing, middleware, request and response handling, and more. Developers use Echo to create web applications and RESTful APIs in Go.

"Echo" is a web application framework for Go, commonly used for building HTTP services and APIs. It provides features to handle routing, middleware, request and response handling, and more. Developers use Echo to create web applications and RESTful APIs in Go.


## Installation
```bash
// go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4
```


## Run
```bash
go get github.com/labstack/echo/v4

go mod init go-web-echo
go mod tidy
go build

http://localhost:8080
```

## Run with query params
````bash

http://localhost:8080/helloworld?hello=Hellow&world=World

````