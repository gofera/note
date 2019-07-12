
# 使用plantuml生成GO的类图
## go-package-plantuml: analysis Go code and generate plantuml code
go get tool and build
```
$ go get git.oschina.net/jscode/go-package-plantuml
```
run to generate plantuml code
```
$ go-package-plantuml --codedir /home/weliu/go/src/github.com/nsqio/go-nsq --gopath /home/weliu/go --outputfile /tmp/go-nsq-result.txt
```
### Bugs
Bug 1: even set `--outputfile` but not generate, and always output to `/tmp/uml.txt`. (Fixed by me)

Bug 2: 聚合显示成关联, 组合没有线, 实现接口而不是继承

### To improve
1. 如果已经有线就不用显示field

## plantuml: convert plantuml code to image format (such as png...)
```
$ git clone https://github.com/plantuml/plantuml.git ~/code/github.com/plantuml/plantuml
$ cd ~/code/github.com/plantuml/plantuml
$ (java 8) mvn package
```
if fail with doc-gen error, comment doc-gen plugin in pom.xml.

Then run:
```
$ java -jar target/plantuml-1.2018.15-SNAPSHOT.jar /tmp/uml.txt
```
if fail with `dot` cannot find, install it: 
```
$ sudo apt install graphviz
```
Then run plantuml.jar again, then `/tmp/uml.png` will be generated.

# reference
[使用plantuml生成GO的类图](https://www.jianshu.com/p/e829f78efc20)