# install virtual box in MacOS
1. restart: command+R
2. menu, list, open ternimal
3. input: csrutil disable
4. restart mac

# MacBook 跳到行尾、行首、Home和end快捷键
在Windows系统中，如果你想跳到行首、行尾直接点击home、end键就可以了，但MacBook的相关快捷键就有些区别了，相关快捷键如下：
```
Ctrl+A：到行首（达到Home键的效果）
Ctrl+E：到行尾（达到End键的效果）
Ctrl+N：到下一行
Ctrl+P：到上一行
Ctrl+K：从光标处开始删除，知道行尾
fn键+左方向键是HOME
fn键+右方向键是END
fn+上方向键是page up
fn+下方向键是page down
```

# MacOS X终端为ls命令设置颜色输出
Ref: [MacOS X终端为ls命令设置颜色输出](https://www.jianshu.com/p/62cdec0fa0a1)

Edit in `~/.bashrc`:
```
export CLICOLOR=1
export LSCOLORS=fxbxcxdxcxegedabagacad
```

