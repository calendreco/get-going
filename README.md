get-going
=========

A skeleton for idiomatic go web development

Why?
----
To create a sensible starting point for futher projects

Tools
-----
* [GoDeps](http://godoc.org/github.com/tools/godep) - Fixed dependencies
* [Ginkgo](https://onsi.github.io/ginkgo/) - Test framework
* [GoImports](http://godoc.org/code.google.com/p/go.tools/cmd/goimports) - Format go files and fix common issues
* [GoDef](http://godoc.org/code.google.com/p/rog-go/exp/cmd/godef) - Used by text editors for symbol lookup
* [Gin](https://github.com/codegangsta/gin) - Live reloading

```bash
go get github.com/tools/godep github.com/onsi/ginkgo code.google.com/p/go.tools/cmd/goimports code.google.com/p/rog-go/exp/cmd/godef github.com/codegangsta/gin
```

Text editor plugins
-------------------

### Sublime 
[GoSublime](https://github.com/DisposaBoy/GoSublime)

Structure
---------
Built to support many packages with a common lib. Allows seperation of web & background processes.

Installing
----------
```bash
mkdir -p $GOPATH/src/github.com/calendreco
cd $GOPATH/src/github.com/calendreco
git clone github.com/calendreco/get-going
cd get-going
godep restore
```

Building, testing, running
--------------------------
Go commands can operate on multiple packages, for example:
```bash
go build ./...
go test ./...
```
Will build & test for all your packages (web, worker etc...)
```bash
cd packages/web
gin
```
Will begin watching,building & running the web binary

Deploying
---------
First fix your dependencies to avoid issues in production
```bash
godep save -copy=false ./...
````
This will generate Godeps that will tie packages to the revision you have.

```bash
heroku create -b https://github.com/calendreco/heroku-buildpack-go.git
```
This will create a heroku instance using the go buildpack. The buildpack will fetch dependencies according to your godeps and build a binary.
