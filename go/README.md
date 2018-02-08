# 安装
绿色安装，可以去[官网](https://golang.org/dl/)下载zip版，然后按照[官网的安装方法](https://golang.org/doc/install)就可以。我在安装操作是：

解压到`$HOME/tool/go`，
```
export GOROOT=$HOME/tool/go
export PATH=$PATH:$GOROOT/bin
```
## 测试是否成功
不设置`GOPATH`，缺省为`$HOME/go`，
```
$ cd $HOME/go/
$ mkdir src
$ cd src
$ mkdir hello
$ cd hello
$ vi hello.go
```
在hello.go中输入如下内容：
```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
编译：
```
$ go build
$ ls -l
total 1909
-rwxr-xr-x 1 weliu 1049089 1952256 Feb  5 08:58 hello.exe*
-rw-r--r-- 1 weliu 1049089      78 Feb  5 08:58 hello.go

$ ./hello
hello, world

```
推荐使用`go install`，这样会在`$GOPATH/bin`下生成可执行文件。
```
cd $GOPATH         # 需要在 $GOPATH 下
go install hello   # hello 为 $GOPATH/src 下的子目录，里面是编译的内容
```

# Go命令行工具
## get
下载工具，比如下载gocode：
```
go get -u github.com/nsf/gocode
go get -u golang.org/x/tools/cmd/guru
go get -u github.com/rogpeppe/godef
```

go build : 编译出可执行文件

go install : go build + 把编译后的可执行文件放到GOPATH/bin目录下

go get : git clone + go install

# 集成开发环境IDE
## JetBrains的Goland
这是Go最好的IDE，官网：https://www.jetbrains.com/go/。

非常好，缺点是收费。

## LiteIDE
第一款专门开发Go的开源免费IDE。

GitHub地址：https://github.com/visualfc/liteide，有中文官网：http://liteide.org/cn/documents/。

解压即可运行。免费软件中的首选。有简单的重构（rename），引用查找。Go非常好，不需要什么工程文件，编译文件。

## Eclipse
在Eclipse Market中搜索Go，安装插件，可参考：[Eclipse配置开发Go的插件——Goclipse](http://blog.csdn.net/linshuhe1/article/details/73473812)。

调试可以使用LiteIDE下的GDB，方法是在Debug Configuration中指定GDB的路径。

好的地方是免费，不好的地方是读代码没有tooltip显示变量类型，点击一个变量没有高亮其它引用。


# 调试
## GDB
Windows下可以使用LiteIDE里面的GDB。

## Delve
GitHub地址：https://github.com/derekparker/delve，可以用 `go get` 下载安装：

```
go get -u github.com/derekparker/delve/cmd/dlv
```

# 第三方库
安装：
```
go get -u github.com/...   （项目的github路径）
```
会下载到$GOPATH/src下，并安装到$GOPATH/pkg下。

## Go编译能不能像mvn那样下载所需依赖？（我觉得应该没有）

## [RxGo](https://github.com/ReactiveX/RxGo)
安装：
```
go get -u github.com/reactivex/rxgo
```
然后就可以在项目上使用了：
```
import (
	"github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/observer"
	"github.com/reactivex/rxgo/observable"
	//...
)
observable.Just(...).Subscribe(...)
```

# Reference
[Download](https://golang.org/dl/)

[Install](https://golang.org/doc/install)

[How to Write Go Code](https://golang.org/doc/code.html)

[Eclipse配置开发Go的插件——Goclipse](http://blog.csdn.net/linshuhe1/article/details/73473812)

[Go 语言的包依赖管理](https://io-meter.com/2014/07/30/go's-package-management/)

[使用Delve进行Golang代码的调试](https://yq.aliyun.com/articles/57578)

[Go语言实战笔记（二十三）| Go 调试](http://www.flysnow.org/2017/06/07/go-in-action-go-debug.html)

[Go语言几大命令简单介绍](http://blog.csdn.net/wuya814070935/article/details/50219915)

[windows下用eclipse+goclipse插件+gdb搭建go语言开发调试环境](http://rongmayisheng.com/post/windows%E4%B8%8B%E7%94%A8eclipsegoclipse%E6%8F%92%E4%BB%B6gdb%E6%90%AD%E5%BB%BAgo%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E8%B0%83%E8%AF%95%E7%8E%AF%E5%A2%83)

