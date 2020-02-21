# show tab in 2 spaces
在vi中输入：
```
:set ts=2
```

# 修改VI默认配置
```
vim /etc/vimrc
```
然后输入：
```
set ts=2
```
即可把VI默认的显示tab为2个空格。

如果只改自己的用户，使用：
```
vi ~/.vimrc
```

# 显示行号
在vi中输入：
```
:set nu
```
# show current file and its folder path
current file: `%`

current folder: `%:p:h`

# open current folder
```
:E
```
# copy to clipboard and paste from clipboard

On Mac OSX
```
copy selected part: visually select text(type v or V in normal mode) and type :w !pbcopy
copy the whole file :%w !pbcopy
paste from the clipboard :r !pbpaste
```

On most Linux Distros, you can substitute:

```
pbcopy above with xclip -i -sel c or xsel -i -b
pbpaste using xclip -o -sel -c or xsel -o -b
-- Note: In case neither of these tools (xsel and xclip) are preinstalled on your distro, you can probably find them in the repos
```

# Reference
1. [Vim 中进行文件目录操作](https://www.cnblogs.com/Dev0ps/p/11661394.html)
2. [史上最全Vim快捷键键位图(入门到进阶)](https://www.runoob.com/w3cnote/all-vim-cheatsheat.html)
