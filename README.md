# 介绍

作者：京城郭少

利用本仓库，可以模拟键盘自动键入字符串

基于github.com/micmonay/keybd_event开发，万分感谢作者

# 注意事项

只支持英文

* 使用前请切换到英文输入法。因为本程序只支持英文（模拟的是按键，而不是传递字符串）。

Mac中使用可能会报错

* 需要用vendor，并将vendor/github.com/micmonay/keybd_event/keybd_darwin.go的```#import <Foundation/Foundation.h>``` 修改成 ```#import <Cocoa/Cocoa.h>```
* 原因请参考 [issue-44](https://github.com/micmonay/keybd_event/issues/44)

# Demo

此示例程序将会模拟键盘输入"hello world!"

```go
package main

import (
	"github.com/GuoFlight/gkeybd"
)

func main() {
	err := gkeybd.TypeStr("hello world!")
	if err != nil {
		panic(err)
	}
}
```