# 安装
绿色安装，可以去[官网](https://golang.org/dl/)下载zip版，然后按照[官网的安装方法](https://golang.org/doc/install)就可以。我在安装操作是：

解压到`$HOME/tool/go`，
```
export GOROOT=$HOME/tool/go
export PATH=$PATH:$GOROOT/bin
```
## 国内下载Go及第三方包
国内长城挡住了Go，可以去[Golang中国](https://www.golangtc.com/)下载：
1. [Go.cn](https://golang.google.cn/dl/), [国内Go下载地址](https://www.golangtc.com/download)
2. [国内第三方包下载方法](https://www.golangtc.com/download/package)

附上 golang.org/x/xxx 包的手动安装方法：

golang.org/x/xxx 这类包托管在 github.com/golang，从这里找到相应的包即可。比如 golang.org/x/crypto 包的安装，找到对应的地址为： https://github.com/golang/crypto ，运行以下命令：
```
$ cd $GOPATH/src
$ mkdir golang.org
$ cd golang.org
$ mkdir x
$ cd x
$ git clone https://github.com/golang/crypto.git
$ go install golang.org/x/crypto
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
go get = git clone + go install



go build : 编译出可执行文件

go install : go build + 把编译后的可执行文件放到$GOPATH/bin目录下

可以把$GOPATH/bin加到PATH路径，这样就可以在任何路径直接使用go get/install 下载的或自己编译出来的程序。

## test
单元测试用go test：
```
go test github.com/WenzheLiu/GoCRMS/gocrms  // run all test in package gocrms
go test -run TestCrmsd github.com/WenzheLiu/GoCRMS/gocrms  // run TestCrmsd in package gocrms
```

# 集成开发环境IDE
## JetBrains的Goland
这是Go最好的IDE，官网：https://www.jetbrains.com/go/。

非常好，缺点是收费。license server: http://idea.youbbs.org

ZuChiMa: 
1. http://blog.csdn.net/john_f_lau/article/details/78762330
2. https://www.youbbs.org/t/2115

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
Windows下可以使用LiteIDE里面的GDB。由于是GDB并不是专门为Go设计的，所以不理解Go协程，更推荐使用Delve来调试。

## Delve
Delve是一个更好用的专门为Go设计的调试工具，
GitHub地址：https://github.com/derekparker/delve，可以用 `go get` 下载安装：

```
go get -u github.com/derekparker/delve/cmd/dlv
```
可以`dlv --help`给出帮助，比如调试可以用下面命令：
```
dlv debug wenzhe/lab      # wenzhe/lab包是要调试的包，相对于$GOPATH/src的路径, 里面有可执行程序，
```
进入`debug`模式后，可以`help`列出帮助，其实与gdb的命令差不多，多了协程支持。
```
(dlv) b main.main  # 包名.方法名 或者 路径/xx.go:行号
Breakpoint 1 set at 0x401018 for main.main() /home/weliu/code/go/src/wenzhe/lab/hello.go:30
(dlv) c
> main.main() /home/weliu/code/go/src/wenzhe/lab/hello.go:30 (hits goroutine(1):1 total:1) (PC: 0x401018)
    25:		})
    26:		<-sub
    27:		fmt.Println("After <-sub")
    28:	}
    29:	
=>  30:	func main() {
    31:		a := "hello"
    32:		b := 123
    33:		fmt.Println(a, b)
    34:		//HelloRxGo()
    35:	}
(dlv) n
(dlv) n
> main.main() /home/weliu/code/go/src/wenzhe/lab/hello.go:32 (PC: 0x401047)
    27:		fmt.Println("After <-sub")
    28:	}
    29:	
    30:	func main() {
    31:		a := "hello"
=>  32:		b := 123
    33:		fmt.Println(a, b)
    34:		//HelloRxGo()
    35:	}
(dlv) print a
"hello"
```
如果被调试程序有参数，使用`--`分隔，
```
dlv exec etcd -- --name=... port xxx  # --name，port为可执行程序etcd的参数
```
对类的方法设断点，比如：Config类的PeerURLsMapAndToken方法：
```
b embed.(*Config).PeerURLsMapAndToken
```

# Go Web开发
## Go与Angular2+
前端Angular开发，编译产生dist文件夹。
```
ng build [-prod]
```
Go实现后端提供REST服务，并作为服务器启动，代码很简单：
```
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/api/servers", getServers)           // 提供REST服务
    http.Handle("/", http.FileServer(http.Dir("./dist"))) // Angular编译出的结果文件
    http.ListenAndServe(":8080", nil)                     // 启动Web服务
}

type MyServer struct {
    Host        string `json:"host"`
    Port        int    `json:"port"`
    Status      string `json:"status"`
    IsReachable bool   `json:"isReachable"`
}

func getServers(w http.ResponseWriter, r *http.Request) {
    server := MyServer{"172.168.58.102", 8080, "new", true}
    servers := []MyServer{server}
    output, err := json.Marshal(&servers)
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(output)
}
```
进入到包含dist的目录（也就是Angular项目的目录），Go编译产生可执行文件，运行即可启动Web服务，提供angular页面并能正确展示且能与Web服务通信。
```
cd <angular-project>
go build <go-web-app>
./<go-web-app>
```
打包给别人用，只需要提供Go编译的可执行文件以及angular编译的dist文件夹。

可以与Angular前后端独立开发，对于跨域问题的解决很简单，可查看note/angular/README.md。

# 工具
## gore （交互式命令行，类似Python IDLE）
https://github.com/motemen/gore

## gophernotes （嵌入Jupyter Notebook，基于网页的交互命令行）
https://github.com/gopherdata/gophernotes

## gops
A tool to list and diagnose Go processes currently running on your system

https://github.com/google/gops/

# Mock
```
$ go get github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen
$ mkdir mock
$ mockgen -source=crms.go > gocrms/mock_crms.go
```
也可以使用 go generate 命令，比如把下面代码加入到crms.go文件的第一行，
```
//go:generate mockgen -source crms.go -destination mock/crms_mock.go github.com/WenzheLiu/GoCRMS/v2/gocrms CrmsServer
```
然后使用`go generate`即可生成（注意先创建目录）：
```
$ mkdir mock
$ go generate  github.com/WenzheLiu/GoCRMS/v2/gocrms
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

## Actor
类似于Akka Actor的[protoactor-go](https://github.com/AsynkronIT/protoactor-go)

## etcd
类似ZooKeeper，但更好用，分布式一致性库 (https://github.com/coreos/etcd)

采用raft协议，下面的动画帮助理解：http://thesecretlivesofdata.com/raft/

# 坑
## 不用在unit test的goroutine中使用testing包的方法Fatal或者Error

## 为什么值为 nil 的 error 却不等于 nil
interface 被两个元素 value 和 type 所表示。只有在 value 和 type 同时为 nil 的时候，判断 interface == nil 才会为 true。

Reference: [Golang 博主走过的有关 error 的一些坑](https://deepzz.com/post/why-nil-error-not-equal-nil.html)

# Question

## channel 可以 close(ch) 我们不用的时候需不需要同样关闭掉，不需要的话 close 主要是什么作用的？

channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放。

close可以用来通知channel接收者不会再收到数据。所以即使channel中有数据也可以close而不会导致接收者收不到残留的数据。

## run external command under another user
```
cmd := exec.Command(command, args...)
cmd.SysProcAttr = &syscall.SysProcAttr{}
cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
```

https://stackoverflow.com/questions/21705950/running-external-commands-through-os-exec-under-another-user

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

[Golang中国](https://www.golangtc.com/)

[Golang的包依赖管理 （package dependency manager）](http://blog.csdn.net/abccheng/article/details/51841754)

[Golang 包依赖管理](http://blog.csdn.net/z770816239/article/details/78909011)

[用dep代替 go get 来获取私有库](http://blog.csdn.net/jq0123/article/details/78457210?locationnum=7&fps=1)

[Golang 中的格式化输入输出](http://www.cnblogs.com/golove/p/3284304.html)

[golang使用vendor目录来管理依赖包](https://www.jianshu.com/p/e52e3e1ad1c0)
