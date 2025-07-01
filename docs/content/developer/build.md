---
title: "Building and Contributing"
draft: false
weight: 10
---

To compile this app, Go 1.23+ is needed.

```shell
make
# make install      # install service, ..., dry-run
# make V=1 install  # Do install actions, reak mode
make release        # cross-build some targets
```

Or `go build` straightly:

```shell
go build -tags "full" # 编译完整版本
```

And cross-building by soecifying `GOOS` and `GOARCH` envvars:

```shell
GOOS=windows GOARCH=386 go build -tags "full" # windows x86
GOOS=linux GOARCH=arm64 go build -tags "full" # linux arm64
```
