# fast input history command
Press `Ctrl R`

# show sub folders' size
```
du --max-depth=1 . -m
```
Here, `.` means the current folder, `-m` means the unit is MB, `--max-depth=1` means only shows the first level sub folders.

# 压缩 tar.gz
```
tar -czvf xxx.tar.gz xxx_folder
```

# 解压 tar.gz
```
tar -xzvf xxx.tar.gz
```

# 解压tar.bz2
```
tar -xf **.tar.bz2
```

# 解压tar.xz
```
tar xvJf  ***.tar.xz
```

# ruptime
[显示网络上每个主机的状态](https://www.ibm.com/support/knowledgecenter/zh/ssw_aix_61/com.ibm.aix.cmds4/ruptime.htm)

# SSH Passwordless Login Using SSH Keygen in 5 Easy Steps
The following example shows from fnode403 to visit fnode401/fnode400 without password
```
[weliu@fnode403 pycli]$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/home/weliu/.ssh/id_rsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /home/weliu/.ssh/id_rsa.
Your public key has been saved in /home/weliu/.ssh/id_rsa.pub.
The key fingerprint is:
5c:99:06:86:08:b2:32:53:e9:b1:3b:b8:18:d5:60:38 weliu@fnode403
The key's randomart image is:
+--[ RSA 2048]----+
|..oo . .o        |
|E+= . .. . o     |
|=+ =      =      |
|.o+ .  . o       |
| o .    S        |
|o o              |
|.o .             |
|o                |
|                 |
+-----------------+
[weliu@fnode403 pycli]$ ls ~/.ssh
id_rsa  id_rsa.pub  known_hosts
[weliu@fnode403 pycli]$ cp ~/.ssh/id_rsa.pub ~/.ssh/authorized_keys
[weliu@fnode403 pycli]$ ssh fnode401
[weliu@fnode401 ~]$ exit
logout
Connection to fnode401 closed.
[weliu@fnode403 pycli]$ ssh fnode400
Warning: Permanently added 'fnode400' (RSA) to the list of known hosts.
Last login: Mon Mar 26 20:11:34 2018 from dn121201.10g.tflex.briontech.com
[weliu@fnode400 ~]$ exit
logout
Connection to fnode400 closed.
```

Reference: https://www.tecmint.com/ssh-passwordless-login-using-ssh-keygen-in-5-easy-steps/

# SSH command without prompting the message for ssh key save or cancel options
## Solution 1: command line
```
ssh -q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null my-server-leaf-5 echo "5" >> ~/tmp/a
```
You are looking to disable "Host Key Verification" and you need the following SSH options:
```
StrictHostKeyChecking no
UserKnownHostsFile /dev/null
```
If adding them to the command (rather than your ssh config file) then use
```
-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null
```
after the -q in your example command.

(Seems -q can be removed, not sure by wenzhe)

## Solution 2: configure file

In your ~/.ssh/config (if this file doesn't exist, just create it):
```
Host *
    StrictHostKeyChecking no
```
This will turn it off for all hosts you connect to. You can replace the * with a hostname pattern if you only want it to apply to some hosts.

Make sure the permissions on the file restrict access to yourself only:
```
sudo chmod 400 ~/.ssh/config
```
# lsof
```
lsof -p 744158 | grep ESTABLIS | wc -l
```

# ulimit
```
ulimit -a
```

# sysctl
```
sysctl -l
sysctl limit
cat /etc/sysctl.conf 
```

# gstack
```
gstack 84966
```

# pstack
```
pstack 744158 | grep Thread | wc -l
```

# strace 
```
strace -p 84966
```

# /proc/<pid>
```
cat /proc/85712/limits
```

# watch
```
watch "lsof -p 85712 | grep ESTABLIS | wc -l"
```

# pstree
Show process tree:
```
ps -p <pid>
```

# ps display pid,ppid,pgid,sid
```
$ ps xao pid,ppid,pgid,sid,comm
  PID  PPID  PGID   SID COMMAND
    1     0     1     1 init
 8593     1  8509  8509 dbus-activation
 8597  8593  8509  8509 hud-service
 9360     1  9360  9360 bash
12638     1 12638 12638 bash
21799 12638 21799 12638 gore
22126  9360 22126  9360 testkill
22139     1 21799 12638 gocode
22161  9360 22161  9360 ps
```

# How do I find all files containing specific text on Linux?
```
grep --include=ngen-leaf-*.log -rnw ~/.gocrms/log/ -e 'etcdserver: too many requests'
```
Refer:
```
grep -rnw '/path/to/somewhere/' -e 'pattern'
-r or -R is recursive,
-n is line number, and
-w stands for match the whole word.
-l (lower-case L) can be added to just give the file name of matching files.
```
Along with these, --exclude, --include, --exclude-dir flags could be used for efficient searching:

This will only search through those files which have .c or .h extensions:
```
grep --include=\*.{c,h} -rnw '/path/to/somewhere/' -e "pattern"
```
This will exclude searching all the files ending with .o extension:
```
grep --exclude=*.o -rnw '/path/to/somewhere/' -e "pattern"
```
For directories it's possible to exclude a particular directory(ies) through --exclude-dir parameter. For example, this will exclude the dirs dir1/, dir2/ and all of them matching `*.dst/`:
```
grep --exclude-dir={dir1,dir2,*.dst} -rnw '/path/to/somewhere/' -e "pattern"
```
This works very well for me, to achieve almost the same purpose like yours. For more options check man grep.

# 使用grep排除一个或多个字符串
```
grep -v 'xxx'
```

# mount 其它机器的路径
比如要把 146.106.207.104:/h/user mount到本机器的 /h/user 目录，可以：
```
$ sudo su
# vi /etc/fstab
```
文件最后加入一行：
```
146.106.207.104:/h/user /h/user nfs proto=tcp,nfsvers=3,hard,bg,intr  0 0
```
生效方式：
```
mount -a
```

# curl: web client
http get:
```
curl localhost:7001/api/hello
```
https get:
```
curl -k https://localhost:443/findcode/api/hello
```
# Create the home directory while creating a user

For command line, these should work:
```
useradd -m USERNAME
```
You have to use -m, otherwise no home directory will be created. If you want to specify the path of the home directory, use -d and specify the path:
```
useradd -m -d /PATH/TO/FOLDER USERNAME
```
You can then set the password with:
```
passwd USERNAME
```
All of the above need to be run as root, or with the sudo command beforehand. For more info, run man adduser.

# Create default home directory for existing user in terminal
```
mkhomedir_helper username
```

# solution to no X (graphic) environment with error: xauth not creating .Xauthority file

Follow these steps to create a $HOME/.Xauthority file.

Log in as user and confirm that you are in the user's home directory.

## Rename the existing .Xauthority file by running the following command
```
mv .Xauthority old.Xauthority 
```

## xauth with complain unless ~/.Xauthority exists
```
touch ~/.Xauthority
```
## only this one key is needed for X11 over SSH 
```
xauth generate :0 . trusted 
```
## generate our own key, xauth requires 128 bit hex encoding
```
xauth add ${HOST}:0 . $(xxd -l 16 -p /dev/urandom)
```

## To view a listing of the .Xauthority file, enter the following 
```
xauth list 
```
After that no more problems with .Xautority file since then.

# disk usage (also for mounted usage)
```
$ df
Filesystem             1K-blocks       Used Available Use% Mounted on
/dev/mapper/LOCAL-ROOT
                         1998672    1853340     40476  98% /
tmpfs                    1962148        216   1961932   1% /dev/shm
/dev/sda1                 194241      80309    103692  44% /boot

```

# list size of sub folders
```
$ cd <YOUR_DIR>
$ du --max-depth=1 -h
```

# Bash for loop with range
```
for i in {1..3};
do
    echo "Iteration $i"
done
```

# Apache httpd
Task: Start httpd server:
```
# service httpd start
```
Task: Restart httpd server:
```
# service httpd restart
```
Task: Stop httpd server:
```
# service httpd stop
```
It is also good idea to check configuration error before typing restart option:
```
# httpd -t
# httpd -t -D DUMP_VHOSTS
```

# uninstall package
删除软件及其配置文件
```
sudo apt-get --purge remove <package>
```
删除没用的依赖包
```
sudo apt-get autoremove <package>
```
此时dpkg的列表中有“rc”状态的软件包，可以执行如下命令做最后清理：
```
sudo dpkg -l |grep ^rc|awk '{print $2}' |sudo xargs dpkg -P
```

# grep
```
weliu@x1:~$ grep -brin -e "gopher.*china" ~/ppt --exclude-dir=.git
/home/weliu/ppt/github.com/bigwhite/talks/gopherchina/2017/go-coding-in-go-way-cn.slide:2:20:GopherChina 2017
/home/weliu/ppt/github.com/bigwhite/talks/gopherchina/2017/go-coding-in-go-way-cn.slide:426:10691:  buf.WriteString("gopherchina ")

weliu@x1:~$ grep -rine 'python vs go' ~/ppt --exclude-dir=.git
/home/weliu/ppt/learnt_from_gopherchina2009/learnt.slide:175:* Python VS Go

weliu@x1:~$ head -175 /home/weliu/ppt/learnt_from_gopherchina2009/learnt.slide | grep -c '\* '
23

weliu@x1:~$ grep -rine 'argError' ~/ppt --exclude-dir=.git
/home/weliu/ppt/share_go/goexamples/errors.go:15:type argError struct {
/home/weliu/ppt/share_go/goexamples/errors.go:20:func (e *argError) Error() string {
/home/weliu/ppt/share_go/goexamples/errors.go:26:		return -1, &argError{arg, "can't work with it"}
/home/weliu/ppt/share_go/goexamples/errors.go:49:	if ae, ok := e.(*argError); ok { // check error type

weliu@x1:~$ grep -rin 'goexamples/errors.go' /home/weliu/ppt/share_go/ --exclude-dir=.git --include=*.slide
/home/weliu/ppt/share_go/goexamples.slide:343:.code -edit -numbers goexamples/errors.go /START1 OMIT/,/END1 OMIT/
/home/weliu/ppt/share_go/goexamples.slide:346:.play -edit -numbers goexamples/errors.go /START2 OMIT/,/END2 OMIT/

weliu@x1:~$ head -343 /home/weliu/ppt/share_go/goexamples.slide | grep -c '\* '
59


weliu@x1:~$ grep -nbri -e 'import "os"' ~/ppt --exclude-dir=.git
/home/weliu/ppt/github.com/davecheney/presentations/gopher-puzzlers/missing-panic2.go:3:14:import "os"
/home/weliu/ppt/presentations/gopher-puzzlers/missing-panic2.go:3:14:import "os"
/home/weliu/ppt/share_go/goexamples/string-formatting.go:5:42:import "os"
/home/weliu/ppt/share_go/goexamples/defer.go:3:26:import "os"

weliu@x1:~$ grep -brin -e 'defer.go' ~/ppt --exclude-dir=.git --include=*.slide
/home/weliu/ppt/share_go/goexamples.slide:349:8319:.play -edit -numbers goexamples/defer.go
```
# install Ubuntu in VirtualBox
[基于VirtualBox虚拟机安装Ubuntu图文教程](https://blog.csdn.net/u012732259/article/details/70172704)

安装VirtualBox虚拟机增强功能:
```
$ sudo apt install gcc
$ sudo apt-get install build-essential virtualbox-guest-dkms
$ cd /media/weliu/VBox_GAs_6.1.2
$ sudo ./VBoxLinuxAdditions.run
```
