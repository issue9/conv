conv [![Build Status](https://travis-ci.org/issue9/conv.svg?branch=master)](https://travis-ci.org/issue9/conv)
======

conv包提供了各个类型之间的转换：
```go
conv.MustInt("123", 0)  // 返回123的数值
conv.MustString(123, "")// 返回字符串123
conv.Int("123", 0)      // 返回123数值和nil的error接口
conv.ToByte("12M")      // 返回12*1024*1024的计算结果值.
```

### 安装

```shell
go get github.com/issue9/conv
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/conv)
[![GoDoc](https://godoc.org/github.com/issue9/conv?status.svg)](https://godoc.org/github.com/issue9/conv)


### 版权

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/issue9/conv/blob/master/LICENSE)
