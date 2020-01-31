conv
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fissue9%2Fconv%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/issue9/conv/goto?ref=master)
[![Build Status](https://travis-ci.org/issue9/conv.svg?branch=master)](https://travis-ci.org/issue9/conv)
[![codecov](https://codecov.io/gh/issue9/conv/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/conv)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
======


conv包提供了各个类型之间的转换：

```go
conv.MustInt("123", 0)  // 返回123的数值
conv.MustString(123, "")// 返回字符串123
conv.Int("123", 0)      // 返回123数值和nil的error接口
```

安装
----

```shell
go get github.com/issue9/conv
```

文档
----

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/conv)
[![GoDoc](https://godoc.org/github.com/issue9/conv?status.svg)](https://godoc.org/github.com/issue9/conv)

版权
----

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
