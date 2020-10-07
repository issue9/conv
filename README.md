conv
[![Go](https://github.com/issue9/conv/workflows/Go/badge.svg)](https://github.com/issue9/conv/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/issue9/conv/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/conv)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/conv)](https://pkg.go.dev/github.com/issue9/conv)
======

conv包提供了各个类型之间的转换：

```go
conv.MustInt("123", 0)  // 返回 123 的数值
conv.MustString(123, "")// 返回字符串 123
conv.Int("123", 0)      // 返回 123 数值和 nil 的 error 接口
```

安装
----

```shell
go get github.com/issue9/conv
```

版权
----

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
