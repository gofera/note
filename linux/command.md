# fast input history command
Press `Ctrl R`

# show sub folders' size
```
du --max-depth=1 . -m
```
Here, `.` means the current folder, `-m` means the unit is MB, `--max-depth=1` means only shows the first level sub folders.

# 解压tar.bz2
```
tar -xf **.tar.bz2
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
