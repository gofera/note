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
## LiteIDE
GitHub地址：https://github.com/visualfc/liteide，有中文官网：http://liteide.org/cn/documents/。

解压即可运行。

## Eclipse
在Eclipse Market中搜索Go，安装插件，可参考：[Eclipse配置开发Go的插件——Goclipse](http://blog.csdn.net/linshuhe1/article/details/73473812)。

调试可以使用LiteIDE下的GDB，方法是在Debug Configuration中指定GDB的路径。


# 调试
## GDB
Windows下可以使用LiteIDE里面的GDB。

## Delve
GitHub地址：https://github.com/derekparker/delve，可以用 `go get` 下载安装：

```
go get -u github.com/derekparker/delve/cmd/dlv
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