# Gonban

Cross-platform [kanban board](https://github.com/huytd/kanelm) with [jsonbin.io](https://jsonbin.io/) as backend. Inspired by [huytd/kanban-app](https://github.com/huytd/kanban-app).

## Status

It is a proof of concept. I wanted to test out the native webview approach using [zserge/webview](https://github.com/zserge/webview).

## How to build and package

### Before building

You need these resources:

* yarn
* libwebkit2gtk-dev
* elm: `yarn global add elm`
* [go-bindata](https://github.com/go-bindata/go-bindata): `go get -u -v github.com/go-bindata/go-bindata/...`
* a [jsonbin.io](https://jsonbin.io/) account

### Build

*This instructions have been tested only under Linux.*

```
go get -u -v github.com/imdario/gonban

# kanelm
git clone https://github.com/huytd/kanelm
cd kanelm
yarn install
elm-package
yarn build

# gonban
cd $GOPATH/src/github.com/imdario/gonban
cp $OLDPWD/dist/dist.js www
go-bindata www/dist.js
go build
```