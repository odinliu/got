# got
[![Build Status](https://travis-ci.org/odinliu/got.svg)](https://travis-ci.org/odinliu/got)

用trick的方式解决`golang.org/x`在国内被墙无法用go get的问题。

# usage

```
$ mkdir -p $GOPATH/src/golang.org
$ mkdir -p $GOPATH/src/github.com/golang
$ ln -s $GOPATH/src/github.com/golang $GOPATH/src/golang.org/x
$ go get -u -v github.com/odinliu/got
$ got -u -v golang.org/x/net/trace
```

# known issues

- 由于只是简单的修改`golang.org/x`为`github.com/golang`，所以里面递归下载的无法替换
- 会有类似`package github.com/golang/net/trace: code in directory $GOPATH/src/github.com/golang/net/trace expects import "golang.org/x/net/trace"`的错误日志
