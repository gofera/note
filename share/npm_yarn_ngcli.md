# 从Java角度理解Angular之入门篇：npm, yarn, Angular CLI

本系列从`Java`程序员的角度，带大家理解前端`Angular`框架。

本文是入门篇。笔者认为亲自动手写代码做实验，是最有效最扎实的学习途径，而搭建开发环境是学习一门新技术最需要先学会的技能，是入门的前提。

作为入门篇，本文重点介绍`Angular`的开发、编译工具：`npm, yarn, Angular CLI`，它们就像`Java`在中的`Maven`，同时顺便介绍一些`Angular`的概念。学习之后，希望你能够在自己的环境下练习、探索、编写出自己的第一个基于`Angular`的`Web`应用。

在开始介绍之前，先了解下一些背景知识，理解单页应用与传统基于模板的多页应用在`Web`开发思路的不同。

# 什么是单页应用（Single Page Application，SPA）
单页应用是`Web`前端发展的主流趋势。

## 服务端渲染
传统的`Web`前端是多页应用的，客户端浏览器向服务器提交请求，服务端通过处理`HTML`模板（如`PHP`,`ASP`,`JSP`等），生成`HTML`文本返回给浏览器，浏览器重新渲染并且展示。HTML完全有后端决定，因此称为“服务端渲染”。客户端每次拿到都是一个HTML页面，这样一个Web应用就有很多页面，因此叫多页应用。
```
服务器 ----(html页面)---> 浏览器
```
由于浏览器每次都要解析整个`HTML`并渲染，因此效率较低，每次即使更新一个数据也要在网络上传输整个`HTML`文本，占用更多的带宽。

## 客户端渲染
不同于传统多页应用，在`SPA`应用中，客户端浏览器向服务器提交请求，服务端返回数据（通常是`json`格式）给浏览器，浏览器中的`js`更新相应部分的`DOM`，从而更新展示。渲染的过程是在客户端浏览器完成的。
```
服务器 ----(json数据)---> 浏览器
```

优点：

1. 局部刷新。无需每次都进行完整页面请求。
2. 天然的前后端分离。后端不再有渲染展示逻辑，只要专注业务逻辑并能够向前端提供数据；而前端只关注如何展示数据、处理用户交互，提高用户体验。这很像`C/S`架构，富客户端，只不过客户端是浏览器中的`js`脚本。
3. 计算量转移。原本需要后端渲染的任务转移给了前端，减轻了服务器的压力。

缺点：

1. 首次打开较慢。由于是富客户端(js脚本)，当客户端首次打开网页，需要先下载一堆`js`和`css`后才能看到页面。对于这一点，`SPA`也有应对方案，比如可以（1）分拆打包（把`js`拆成多个包，首屏用到的先传）；（2）先传个友好的加载页面；（3）同构（第一次加载是采用后端渲染，以后用前端渲染）。
2. 不方便搜索引擎。传统的搜索引擎会从`HTML`中抓取数据，导致前端渲染的页面无法被抓取。不过这一点，随着`SPA`的流行，`Google`爬虫也可以像浏览器一样理解`js`脚本了。

## 流行的SPA框架
流行的`SPA`框架有`React`, `Vue`, `Angular`。本文基于`Angular 2/4/5+`（不是`Angular 1.x`或`AngularJS`）。

# node.js
就像`Java`开发需要`JDK`，`Angular`开发需要`node.js`。`Java`编译出来的字节码需要`Java`虚拟机`JRE`执行，而如果想在浏览器以外执行`JavaScript`代码，也需要类似的虚拟机平台，也就是`node.js`。

类似`JDK`，`node.js`下载之后也不需要安装，只要加到`PATH`路径下即可。这里需要`node.js`是因为很多前端开发工具是有`JavaScript`写成的，如`npm`，它们需要`node.js`这个虚拟机。

# 项目依赖管理工具
像`Java`中的`maven`，开发`Angular`可以使用`npm`或者`yarn`。其中`npm`是`node.js`自带的，可以直接使用。

另外，`npm`还有`maven`不具备的能力，它可以从网上下载并安装软件，类似于`Linux`中的`yum`。比如，`yarn`可以通过`npm`下载安装：
```
npm install --global yarn
```
`npm`作为项目依赖管理有个缺点，它没有本地仓库，对同一个依赖，不管其他项目是否已经下载过，只要这个项目没有，它都从网上下载。下载后存在每个项目下的`node_modules/`路径下。相比，`maven`有本地仓库，对同一个依赖只会下载一次，存在`~/.m2/repository/`下。

与`npm`不同的是，`yarn`无需互联网连接就能安装本地缓存的依赖项，它提供了离线模式。只要其他项目已经下载过，就不会上网下载，但依然会拷贝到项目下的`node_modules/`路径。另外，`yarn`的运行速度得到了显著的提升，整个安装时间也变得更少。所以推荐使用`yarn`管理项目依赖。

## 仓库
`maven`有中央仓库，也可以创建私服；`npm`,`yarn`同样都有，可以配置：
```
npm config set registry http://registry.npmjs.org/
yarn config set registry https://registry.yarnpkg.com/
```

### 科学上网
国内用户可以通过淘宝镜像提高下载速度：
```
npm config set registry https://registry.npm.taobao.org
yarn config set registry https://registry.npm.taobao.org
```

## 代理
如果上网需要代理的话，可以在`~/.bashrc`加入如下内容：
```
######################
# User Variables (Edit These!)
######################
username="myusername"
password="mypassword"
proxy="mycompany:8080"

######################
# Environement Variables
# (npm does use these variables, and they are vital to lots of applications)
######################
export HTTPS_PROXY="http://$username:$password@$proxy"
export HTTP_PROXY="http://$username:$password@$proxy"
export http_proxy="http://$username:$password@$proxy"
export https_proxy="http://$username:$password@$proxy"
export all_proxy="http://$username:$password@$proxy"
export ftp_proxy="http://$username:$password@$proxy"
export dns_proxy="http://$username:$password@$proxy"
export rsync_proxy="http://$username:$password@$proxy"
export no_proxy="127.0.0.10/8, localhost, 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16"

######################
# npm Settings
######################
npm config set registry http://registry.npmjs.org/
npm config set proxy "http://$username:$password@$proxy"
npm config set https-proxy "http://$username:$password@$proxy"
npm config set strict-ssl false
echo "registry=http://registry.npmjs.org/" > ~/.npmrc
echo "proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "strict-ssl=false" >> ~/.npmrc
echo "http-proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "http_proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "https_proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "https-proxy=http://$username:$password@$proxy" >> ~/.npmrc

######################
# yarn Settings
######################
yarn config set registry https://registry.yarnpkg.com/
yarn config set proxy "http://$username:$password@$proxy"
yarn config set https-proxy "http://$username:$password@$proxy"
yarn config set strict-ssl false
```

# 脚手架
脚手架可以简化开发，创建工程框架。比如`maven`可以使用`archetype:generate`来创建基于`maven`约定的`java`工程。

脚手架创建的项目，一般包含以下内容：
1. 编译脚本、依赖定义
2. 源代码路径及样例
3. 单元测试路径及样例
4. 资源路径
5. 环境配置（开发环境、生产环境、测试环境）
6. 编译目标路径（编译后生成）

## Angular脚手架
`Angular CLI`是`Angular`官方提供的命令行工具，可以帮助我们创建项目、编译、测试、运行，就像`Java`界的`Maven`，另外还可以创建组件、管道、指令、服务、模块、类、接口、枚举。。。

可以用`npm`安装`Angular CLI`：
```
npm install --global @angular/cli
```
创建`Angular`项目：
```
ng new <project_name>
```
脚手架会帮我们创建以下文件：
```
|-- package.json                   # 编译脚本、依赖包管理，类似maven的pom.xml。
|-- src                            # 原代码，类似maven的src/main/java
|   |-- app                        # 该应用的模块路径，angular支持模块化！
|   |   |-- app.component.css      # 组件样式（作用域只在该组件内部有效）
|   |   |-- app.component.html     # 组件模板（视图）
|   |   |-- app.component.spec.ts  # 组件单元测试
|   |   |-- app.component.ts       # 组件控制代码（控制器）
|   |   |-- app.module.ts          # 模块管理代码
|   |-- assets                     # 资源文件夹，类似maven的src/main/resources
|   |-- environments
|   |   |-- environment.prod.ts    # 生产环境配置
|   |   `-- environment.ts         # 开发环境配置
|   |-- favicon.ico                # 网页的图标
|   |-- index.html                 # 网页的HTML，使用根组件（一般无需修改）
|   |-- main.ts                    # 入口代码，引导根模块（一般无需修改）
|   |-- styles.css                 # 全局样式
|   |-- test.ts                    # 单元测试入口（一般无需修改）
|   |-- ...
|-- tsconfig.json                  # typescript配置，比如开发使用的ES版本，编译生成的目标ES版本（默认是ES5,即目前广泛的javascript，一般无需修改)
|-- tslint.json                    # typescript的语法规则，方便IDE检查（一般无需修改）
|-- node_modules                   # 项目的依赖包仓库，相当于maven的本地仓库
|-- dist                           # 编译时才生成的目标文件夹（拷贝里面内容到Web服务器即可使用），相当于maven工程下的target目录
|-- .gitignore                     # git ignore（一般无需修改，node_modules和dist以默认排除）
...
```

# 依赖管理
当我们需要更新项目依赖时，可以修改`package.json`文件，然后运行`yarn install`。它会解析`package.json`并下载还没有的依赖。这种方式跟`maven`完全一样。

## 添加/删除依赖
还可以用命令行来添加/删除依赖：
```
yarn add/remove <package_name>[@<version>] [--dev/-D]
```
`[]`表示可选的，不写为最新版本。命令执行后会自动更新`package.json`文件。`--dev`或`-D`表示包只是开发时候需要，最终产品打包是不包含的，类似与`Maven`中的`provided scope`。

# 开发阶段：运行
对于前后端分离的应用，前端开发与后端独立开，前端应用在没有启动后端服务器的情况下也应该能够运行。

`Angular CLI`提供支持，如下：
```
ng serve [--host 0.0.0.0] [--port 4200]
```
`[]`表示可选的，不写端口号默认为`4200`。打开浏览器，输入`localhost:4200`可以看到你开发的网页。

## 代码更新检测
代码文件更新不需要重启服务，当文件内容变化时，网页会自动刷新，开发者可以看到更新后的页面。

## Mock 后端服务
在后端服务还没有准备好之前，当前端需要向后端服务发出请求`（GET/POST/PUT/DELETE）`，让后端更新数据，并且返回给前端时，可以造个假的。

可以安装一个`json-server`，
```
npm install --global json-server
```
写个`json`文件，把假数据填上去。启动`json`服务：
```
json-server <your_mock_data.json> [--port 3000]
```
就可以作为服务器接受客户端的`RESTful`请求`（GET/POST/PUT/DELETE）`了。

# 单元测试
像`mvn test`，`Angular CLI`可以一键运行所有单元测试：
```
ng test
```

# 编译成目标文件
像`mvn compile`，`Angular CLI`可以一键运行编译导出目标文件（默认为`ES5`，即传统的`javascript`）：
```
ng build [-prod]
```
可选参数`-prod`表示生产环境，输出目标会小很多。输出目标为`dist`文件夹：
```
$ ls -l dist
total 413
-rw-r--r-- 1 weliu 1049089   3293 Jan 15 08:50 3rdpartylicenses.txt
-rw-r--r-- 1 weliu 1049089   5430 Jan 15 08:50 favicon.ico
-rw-r--r-- 1 weliu 1049089    597 Jan 15 08:50 index.html
-rw-r--r-- 1 weliu 1049089   1445 Jan 15 08:50 inline.08a75f8119356113a22d.bundle.js
-rw-r--r-- 1 weliu 1049089 341557 Jan 15 08:50 main.e25b64f979f240da775b.bundle.js
-rw-r--r-- 1 weliu 1049089  61268 Jan 15 08:50 polyfills.65fe1626e31e03d17f8e.bundle.js
-rw-r--r-- 1 weliu 1049089      0 Jan 15 08:50 styles.d41d8cd98f00b204e980.bundle.css
```
不如`maven compile`的地方，是`ng build`就只做编译，在编译前也不会先去下载未缓存的依赖包，因此完整的编译前需要手动调用`yarn install`。

# 部署
将编译生成的`dist`文件夹中的所有文件放在`Web`服务器即可。

也可以与服务端代码一起编译，打成一个包，方便部署。

## 与Java后端一起部署（使用Maven）
对于使用`Spring Boot`的后端`Java`代码，只提供`RESTful`的数据服务，可以打包成一个独立可执行的`jar/war`包：
```
$ cd back-end
$ mvn clean package
...
BUILD SUCCESS
...

$ ls target/
...
my-web-server.war
...
```
目前这个`war`包还只是一个提供`RESTful`的数据服务器，还没有网页界面。

可以写个脚本把前端编译后`dist`中的内容，通过`jar`命令打包进上面的`war`包（或者`jar`包），脚本很简单，如下：
```
$ cat add_to_war.sh
cd dist/
jar uf ../../back-end/target/my-web-server.war *
```
这样就`war`包就带有网页界面了，可以放在`Tomcat`服务器下部署，或者直接启动（`Spring Boot`默认会内嵌`Tomcat`服务器）：
```
java -jar back-end/target/my-web-server.war
```
以上过程，可以写一个总的编译脚本`build.sh`，一键编译整个前后端应用：
```
# build back-end by maven or gradle
cd back-end
mvn clean package  # or: gradle clean build

# build back-end by yarn and Angular CLI
cd ../front-end
yarn install -prod
ng build -prod

# add to war by jar command
cd dist/
jar uf ../../back-end/target/my-web-server.war *
```

## 与Java后端一起部署（使用Gradle）
如果使用`Gradle`编译打包整个前后端应用，就更简单，因为已经有`Gradle`插件`com.moowork.node`直接支持。

在前端`Angular`项目（比如名为`front-end`)中，加入`build.gradle`文件，内容如下：
```
plugins {
  id "com.moowork.node" version "0.13"
}

apply plugin: 'java'

// configure gradle-node-plugin
node {
  version = '6.5.0'
  npmVersion = '3.10.7'
  // If true, it will download node using above parameters.
  // If false, it will try to use globally installed node.
  download = false // true
  workDir = file("${project.projectDir}/node")
}
task compileTypeScript(type: NpmTask) {
  // install the express package only
  args = ['run-script', "tsc"]
}

// clean node/node_modules/dist
task npmClean(type: Delete, group: 'node') {
  final def webDir = "${project.projectDir}"
  //delete "${webDir}/node"
  //delete "${webDir}/node_modules"
  delete "${webDir}/dist"
}

clean.dependsOn(npmClean)

task npmStart(type: NpmTask) {
  args = ['start']
  group = "node"
  dependsOn("npmInstall")
}

task npmBuild(type: NpmTask) {
  args = ['run', 'build']
  group = "node"
  dependsOn("npmInstall")
}

jar {
  dependsOn("npmBuild")
  from(fileTree("dist")) {
    into "META-INF/resources"
  }
}
```
在前端`Angular`项目的根目录中直接执行：
```
gradle build
```
就可以编译、打包，在`Gradle`约定的输出目录`build/libs/`下生成目标`jar`文件，包含所有`Web`资源。
```
front-end
|   build
|   |-- libs
|   |   |-- front-end-0.0.1.jar
```

在后端`Java`项目`back-end`中的`build.gradle`中加入对前端`Angular`项目的依赖即可。
```
dependencies {
  compile project(":front-end")
  ...
```
为了一键编译、打包整个前后端工程，可以在前后端项目的父目录中加入`build.gradle`:
```
task wrapper(type: Wrapper) {
  gradleVersion = '3.0'
}
```
和`settings.gradle`：
```
rootProject.name = 'my-web-product'
include 'front-end'
include 'back-end'
```
在此根目录中，可以一键直接执行：
```
gradle build
```
这样编译、打包整个前后端工程，在后端项目中生成最终包含`Web`资源的文件：
```
my-web-product
|   back-end
|   |   build
|   |   |-- libs
|   |   |   |-- my-web-product-0.0.1.war
```

可以看到，就像`maven`或`gradle`一样，`Angular CLI`可以帮我们创建脚手架、编译、测试、运行，但她的能力可远不只这些，在开发过程中依然是个好助理，比如可以帮我们创建：组件、服务、管道、指令、模块等等。

# 组件

组件是一个独立的、可复用的、可组合的`UI`控件，网页界面就是一系列组件的有机结合。

组件由组件名、视图、控制器组成。

组件相当于面向对象中的“类”，组件实例相当于“对象”。在`Angular`中，组件就是一个装饰了`@Component`注解的类。

## 脚手架创建组件
下面语句创建一个名为`hello`的组件（同时自动将组件注册到模块`app.module.ts`中，这样同模块的其它组件才可以使用)：
```
ng g[enerate] c[omponent] hello [--inline-template] [--inline-style] [--spec false]
```
`[]`为可选，默认情况下生成下面几个文件：
```
src
|   app
|   |-- hello
|   |   |-- hello.component.css      # 组件样式（作用域只在该组件内部有效）
|   |   |-- hello.component.html     # 组件模板（视图）
|   |   |-- hello.component.spec.ts  # 组件单元测试
|   |   |-- hello.component.ts       # 组件控制代码（控制器）
```
嫌文件太多，也可以把`HTML`模板和`CSS`都内联进`TypeScript`，并且不创建单元测试：
```
src
|   app
|   |-- hello
|   |   |-- hello.component.ts       # 组件全部代码：控制器、视图、样式
```

内联了模板、`CSS`的文件`hello.component.ts`长这个样子：
```
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-hello',   // 组件选择器（组件名）
  template: `              // 组件模板（视图）
    <p>
      hello works!
    </p>
  `,
  styles: []               // 样式CSS（作用域只在该组件内部有效）
})
export class HelloComponent implements OnInit { // 组件控制代码（控制器）

  constructor() { }

  ngOnInit() {
  }

}
```
可以在模块内其它组件`HTML`模板中使用该组件：
```
<app-hello></app-hello>
```

# 服务
前端与后端交互的代码逻辑，在`Angular`中一般设计成一个服务。服务是一个可以依赖注入的（非常像`Spring`），容器管理的类。

`Angular CLI`可以帮我们生成一个服务的代码，下面的命令在`src/app/services/`目录下创建一个名为`job`的服务（同时自动将组件注册到模块`app.module.ts`中，这样同模块的其它组件才可以使用）：
```
ng g[enerate] s[ervice] services/job [--spec false]
```
创建的文件为：`src/app/services/job.service.ts`。添加代码后可以与后端进行交互（GET/POST/PUT/DELETE)。
```
import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { Server } from '../model/server';
import { JobDetail } from '../model/job.detail';

@Injectable()
export class JobService {

  private readonly headers = new Headers({
    'Content-Type': 'application/json'
  });

  constructor(private http: Http) {}

  jobs(host: string, port: number): Observable<JobDetail[]> {
    return this.get('jobs', {
      host: host,
      port: port
    });
  }

  runJob(job: string[], servers: Server[]) {
    if (job.length === 0 || servers.length === 0) {
      return;
    }
    const serverHostPorts = servers.map(server => `${server.host}:${server.port}`);
    console.log(`run job ${job} on servers ${serverHostPorts}`);
    this.post('run', {
      job: job,
      hostPorts: serverHostPorts
    }).subscribe();
  }

  private get(action: string, params: any = {}): Observable<any> {
    return this.http.get(`api/${action}`, {
      headers: this.headers,
      params: params
    })
    .map(res => res.json());
  }

  private post(action: string, data: any = {}, params: any = {}): Observable<any> {
    return this.http.post(`api/${action}`, data, {
      headers: this.headers,
      params: params
    });
  }
```
虽然`Angular CLI`都会给服务类修饰`@Injectable`注解，但并不是服务的标识，也不是必要的。它的作用是让`Angular`框架能够向该类注入其它服务，比如构造函数中`Http`服务就是由`Angular`框架通过依赖注入进来的。这种机制与`Spring`非常相似。

`Angular`是基于`rxjs`的，操作`http`会得到一个可观察对象`Observerable<any>`，异步编程、事件处理都非常简单。熟悉`RxJava`的同学觉得非常亲切。

`Angular`是一个类似于`Spring`的框架，其容器管理服务和组件，所有服务都是可以依赖注入的（只提供构造器注入，这也是不可变成员变量的最佳实践）。所有，在组件或其它服务中使用服务就非常简单，只要在构造函数把服务传进来即可，下面演示了`servers`组件使用上面定义的`job`服务：
```
@Component({
  selector: 'app-servers',
  templateUrl: './servers.component.html',
  styleUrls: ['./servers.component.css']
})
export class ServersComponent {

  jobs: JobDetail[] = [];

  constructor(private jobService: JobService) {
  	jobService.jobs().subscribe(jobs => this.jobs = jobs);
  }
```
下面演示`rxjs`的好处，比如我们想每隔`5`秒，异步地去服务器请求数据，拿到后刷新`jobs`控件，一句话搞定：
```
constructor(private jobService: JobService) {
  Observable.interval(5000).subscribe(evt => 
	jobService.jobs().subscribe(jobs => this.jobs = jobs));
｝
```

## 管道
先看个例子来理解管道，比如我们在控制器`typescript`文件中有一个成员变量（字符串数组类型）：
```
command: string[] = ['java', '-jar', 'hello.jar'];
```
在视图`html`模板中显示：
```
<div>
  {{ command }}
</div>
```
会看到显示的内容是：
```
java, -jar, hello.jar
```
这是`string[]`类型的`toString()`方法输出的，但逗号很不好看，希望它输出是这样的：
```
java -jar hello.jar
```
有两种办法，第一种是直接在`html`中（`{{ }}`里面可以是一个表达式）：
```
<div>
  {{ command.join(' ') }}
</div>
```
第二种办法可以创建一个名叫`join`的管道，类型`Linux`管道一样使用：
```
<div>
  {{ command | join: ' ' }}
</div>
```
`Angular CLI`有创建管道的方法，下面的命令在src/app/pipes目录下创建了一个叫`join`的管道（同时自动将组件注册到模块`app.module.ts`中，这样同模块的其它组件才可以使用)：
```
ng g[enerate] p[ipe] pipes/join [--spec false]
```
由注解`@Pipe`修饰的类称为管道类，在生成的`src/app/pipes/join.pipe.ts`中，我们在加入`transform`方法中加入`join`的逻辑。
```
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'join'   // 管道的名字以供HTML模板使用
})
export class JoinPipe implements PipeTransform {

  transform(input: any, character: string = ''): string {
    if (!Array.isArray(input)) {
      return input;
    }
    return input.join(character);
  }
}
```
### 内置管道
当然，`Angular`提供的内置管道非常好用，看下面的代码：
```
class Server {
  host: string;
  port: number;

  toString(): string {         // 作用完全等同Java的toString方法
  	return `${host}:${port}`;  // 与Kotlin一样也支持模板字符串语法
  }
}
myServer: Server = {           // 用json直接赋值，比Java方便不少吧
  host: "127.0.0.1",           // 如果变量名或者赋值类型的值写错了，类型系统会提示出错
  port: 8080
}
```
HTML模板：
```
<div>{{ myServer }}</div>
```
显示结果为：
```
127.0.0.1:8080
```
如果想要把下面的对象`myServer`按照`json`格式打印出来，可以用`json`管道：
```
<div>{{ myServer | json }}</div>
```
显示结果为：
```
{
  "host": "127.0.0.1",
  "port": 8080
}
```
如果上面的`json`是服务端发送过来的，那么结果是一个`rxjs`可观察对象`Observable<Server>`，
```
myServer$: Observable<Server> = http.get('api/server').map(res => res.json());
```
`Observable`类型的变量名以`$`结尾，是`Angular`的一种编程习惯，因为`rxjs`使用频率太高，这是为了区分普通对象和可观察对象。

可以不需要在代码中手动订阅，也不需要考虑什么时候去解除订阅，在`HTML`模板中用`async`管道，
```
<div>{{ myServer$ | async }}</div>
```
当可观察对象有新值(`onNext`)时会在网页中更新，显示结果为：
```
127.0.0.1:8080
```
就像`Linux`中的管道一样，`Angular`的管道也可以级联，上面的结果很容易以`json`格式显示出来：
```
<div>{{ myServer$ | async | json }}</div>
```
显示结果为：
```
{
  "host": "127.0.0.1",
  "port": 8080
}
```

# 模块
可以看到，`Angular CLI`可以帮我们创建组件、服务、管道、指令等，非常方便，它们都有共同的特点，就是都会更改模块定义文件（`app.module.ts`）。

那么，什么是模块？

类似于`Java 9`或`OSGi`的模块化，`Angular`应用天生就是模块化，概念都是差不多的。模块，就是在类和包之上在做封装，通常把协同完成某个功能的一组组件、服务作为一个整体，就是一个模块。打个比方，如果把模块想象成一个班级，那么组件、服务就是班上的同学。

模块具有封装性和隔离性，模块外部无法访问模块内部的组件、服务，除非导入该模块。这一点，用过`OSGi`或`Eclipse RCP`很熟悉，一个模块相对于一个`Eclipse`插件。如果想在一个插件中使用另一个插件的类，常用的做法是在插件声明文件`MANIFEST.MF`中声明为`Require-Bundle`（当然`Eclipse`不只这一种方法，比如还可以动态导入、或只`import`包名让`OSGi`容器自动去匹配插件，目前`Angular`模块化还做不到这一点）。

`Angular CLI`脚手架创建的`src/app/app.module.ts`，定义了一个由注解`@NgModule`装饰的类（模块类），用于管理该模块。这个文件就好比`Eclipse`插件中的`MANIFEST.MF`。一个模块类的代码如下：

```
import { BrowserModule } from '@angular/platform-browser';

@NgModule({                   // 模块类的标识由该注解决定
  declarations: [             // 声明组件、管道、指令，类型由具体类的注解来区分，用Angular CLI创建时会自动填入
    AppComponent, MyPipe, MyDirective ...
  ],
  imports: [                  // imports指定要导入的模块，这样才可以使用其它模块内的组件、指令、管道、服务
    BrowserModule, ...
  ],
  providers: [                // 声明服务，只有声明的服务才会被`Angular`框架创建出来，才能注入到其它服务或组件
    MyService, ...            // 用Angular CLI创建时会自动填入
  ],
  bootstrap: [AppComponent]   // 引导组件
})
export class AppModule { }    // 模块类定义
```

一个`Angular`应用中可以有多个模块，用`Angular CLI`可以很容易创建一个模块：
```
ng g module your-new-module-name
```

# `Angular CLI`创建`TypeScript`类
创建普通类：
```
ng g class my-new-class
```
创建接口：
```
ng g interface my-new-interface
```
创建枚举：
```
ng g enum my-new-enum
```

# IDE
免费版本推荐使用微软的`Visual Studio Code (VSCode)`, 很好用，但有条件一定要买付费的`IDE`：`JetBrain`的`WebStorm`。原因就跟`Java IDE`的选择一样，要免费的`Eclipse`，还是商业使用需付费的`IntelliJ IDEA`，看你或你的公司是不是土豪。

笔者是穷屌丝，使用`VSCode`，编写`Angular`代码，推荐安装下面的插件：

1. Angular 5 Snippets：`Angular`开发必备，这里我使用的是`Angular 5`。
2. TSLint：`TypeScript`基于最佳实践的编程规范检测插件（类似于`Java`的`CheckStyle`），以红线实时提示你的不规范代码。
3. TypeScript Importer：在编写`TypeScript`代码时，当你使用一个新的类，自动插入`import`语句。

`Typescript`类型系统非常好的地方是`IDE`的提示非常好，比如能提示出类有哪些方法、哪些属性，再加上语法和方法与`Java`太像，使用起来非常自然、亲切。

# 总结

读到这里，相信你已经有了`Angular`的开发环境，并且对一些`Angular`的概念有了初步理解，能够与服务端一起编译并部署到`Web`服务器上。

也许你会问我，有那么多`Web`前端框架，为什么偏爱`Angular`？

## 选择`Angular`的原因
笔者作为一个`Java`背景的软件开发人员，选择`Angular`的原因有：

### 使用`TypeScript`编程
`TypeScript`，其实并不是一种新语言，她仍然还是`JavaScript`，是一种带有类型系统的、面向对象版本的`JavaScript`，是`ES6/7/8+`的超集:

1. 语法类似`Java/Kotlin`，很多语言特性、关键字、类方法名都是一样的；
2. 类型系统可以在编译期做检查，避免敲错字；
3. 类型系统可以帮助`IDE`分析代码，代码跳转、引用分析、出错实时提醒等；
4. 类型系统可以提高代码的可读性、可维护性、类型安全性，调用方法时可以避免传入不期望的类型；
5. `TypeScript`是`ES6+`的超集，是`ES6`规范的实现，在未来的浏览器很可能直接支持；
6. `TypeScript`可以直接编译成`ES5`运行于旧浏览器，在现有版本的`Angular`是默认编译目标，无需引入任何编译依赖和配置；
7. `TypeScript`与其编译出的`ES5 javascript`代码是一一对应的，很多时候就是类型擦除，无需额外库，不像`Kotlin.js`那么重。
8. 拼爹时代，这点还是要考虑的：微软的`TypeScript`，加上谷歌的`Angular`，非常看好它的前景。
9. `JavaScript`编写的库可以在`TypeScript`使用，缺乏类型声明部分，社区以提供大量第三方`js`库的类型定义，一般不需要自己做，当然，想自定义个类型也不难。

如果你是`Java`程序员，应该会赞成这句话：“`TypeScript`一点都不难，`JavaScript`才叫难！”。

当然如果你真的不习惯写类型，那不写也可以，因为`TypeScript`是`ES6+`的超集，`ES6+`的代码在`TypeScript`完全合法，不过就没有了编译时类型检查，`IDE`也不会帮你跳转和引用分析，也不会做方法提示，你能接受缺乏类型安全机制的痛苦吗？我不能！

### 很像`Spring Boot`框架
`Angular`框架很像`Spring Boot`框架(都无需额外配置)，有组件、服务、依赖注入、基于注解的配置，容器管理组件、服务的生命周期等特性；

### 基于`rxjs`
`rxjs`可以方便编写异步响应式代码、处理事件总线、大量可组合的操作符、等等（其实就是流行的`RxJava`的`TypeScript`版本）。

### 模块化
模块化对于构建大型复杂的`Web`系统来说尤为重要。

## 所以

对于`Java`、`Kotlin`等静态强类型语言背景的开发人员来说，`Angular`学习门槛较低，因为很多特性、概念都太`Java`了。

而对于`JavaScript`、`Python`等动态语言背景的开发人员，入手`React`, `Vue`可能更容易，因为它们都是基于`ES6`的无类型系统。



——- 本博客所有内容均为原创，转载请注明作者和出处 ——-

作者：刘文哲

联系方式：liuwenzhe2008@qq.com

博客：http://blog.csdn.net/liuwenzhe2008
