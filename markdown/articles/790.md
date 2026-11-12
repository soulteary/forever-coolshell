# 十个你可能不曾用过的Linux命令

下面可能是你不曾用过后十个Linux的命令。相当的有用。

**1）pgrep**

pgrep名字前有个p，我们可以猜到这和进程相关，又是grep，当然这是进程相关的grep命令。不过，这个命令主要是用来列举进程ID的。如：

```shell
$ pgrep -u hchen
22441
22444

```

这个命令相当于：

`ps -ef | egrep '^hchen' | awk '{print $2}'`

**2）pstree**

这个命令可以以树形的方式列出进程。如下所示：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]$ pstree
init-+-acpid
     |-auditd-+-python
     |        `-{auditd}
     |-automount---4*[{automount}]
     |-backup.sh---sleep
     |-dbus-daemon
     |-events/0
     |-events/1
     |-hald---hald-runner---hald-addon-acpi
     |-httpd---10*[httpd]
     |-irqbalance
     |-khelper
     |-klogd
     |-ksoftirqd/0
     |-ksoftirqd/1
     |-kthread-+-aio/0
     |         |-aio/1
     |         |-ata/0
     |         |-ata/1
     |         |-ata_aux
     |         |-cqueue/0
     |         |-cqueue/1
     |         |-kacpid
     |         |-kauditd
     |         |-kblockd/0
     |         |-kblockd/1
     |         |-kedac
     |         |-khubd
     |         |-6*[kjournald]
     |         |-kmirrord
     |         |-kpsmoused
     |         |-kseriod
     |         |-kswapd0
     |         |-2*[pdflush]
     |         |-scsi_eh_0
     |         |-scsi_eh_1
     |         |-xenbus
     |         `-xenwatch
     |-migration/0
     |-migration/1
     |-6*[mingetty]
     |-3*[multilog]
     |-mysqld_safe---mysqld---9*[{mysqld}]
     |-smartd
     |-sshd---sshd---sshd---bash---pstree
     |-svscanboot---svscan-+-3*[supervise---run]
     |                     |-supervise---qmail-send-+-qmail-clean
     |                     |                        |-qmail-lspawn
     |                     |                        `-qmail-rspawn
     |                     `-2*[supervise---tcpserver]
     |-syslogd
     |-udevd
     |-watchdog/0
     |-watchdog/1
     `-xinetd

```

**3）bc**

这个命令主要是做一个精度比较高的数学运算的。比如开平方根等。下面是一个我们利用bc命令写的一个脚本（文件名：sqrt）

```shell
#!/bin/bash
if [ $# -ne 1 ]
then
    echo 'Usage: sqrt number'
    exit 1
else
    echo -e "sqrt($1)\nquit\n" | bc -q -i
fi

```

于是，我们可以这样使用这个脚本进行平方根运算：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)]$ ./sqrt 36
6
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)]$ ./sqrt 2.0000
1.4142
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)]$ ./sqrt 10.0000
3.1622

```

**4）split**

如果你有一个很大的文件，你想把其分割成一些小的文件，那么这个命令就是干这件事的了。

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) applebak]# ls -l largefile.tar.gz
-rw-r--r-- 1 hchen hchen 436774774 04-17 02:00 largefile.tar.gz 

[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) applebak]# split -b 50m largefile.tar.gz LF_

[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)]# ls -l LF_*
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:34 LF_aa
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:34 LF_ab
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:34 LF_ac
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:34 LF_ad
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:34 LF_ae
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:35 LF_af
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:35 LF_ag
-rw-r--r-- 1 hchen hchen 52428800 05-10 18:35 LF_ah
-rw-r--r-- 1 hchen hchen 17344374 05-10 18:35 LF_ai

```

文件合并只需要使用简单的合并就行了，如：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)]#  cat LF_* >largefile.tar.gz

```

**5）nl**

nl命令其它和cat命令很像，只不过它会打上行号。如下所示：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) include]# nl stdio.h | head -n 10
     1  /* Define ISO C stdio on top of C++ iostreams.
     2     Copyright (C) 1991,1994-2004,2005,2006 Free Software Foundation, Inc.
     3     This file is part of the GNU C Library. 

     4     The GNU C Library is free software; you can redistribute it and/or
     5     modify it under the terms of the GNU Lesser General Public
     6     License as published by the Free Software Foundation; either
     7     version 2.1 of the License, or (at your option) any later version.

     8     The GNU C Library is distributed in the hope that it will be useful,

```

**6）mkfifo**

熟悉Unix的人都应该知道这个是一个创建有名管道的系统调用或命令。平时，我们在命令行上使用竖线“|”把命令串起来是使用无命管道。而我们使用mkfifo则使用的是有名管道。下面是示例：

下面是创建一个有名管道：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# mkfifo /tmp/hchenpipe

[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# ls -l /tmp
prw-rw-r-- 1 hchen  hchen  0 05-10 18:58 hchenpipe

```

然后，我们在一个shell中运行如下命令，这个命令不会返回，除非有人从这个有名管道中把信息读走。

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# ls -al > /tmp/hchenpipe

```

我们在另一个命令窗口中读取这个管道中的信息：（其会导致上一个命令返回）

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# head /tmp/hchenpipe
drwx------ 8 hchen hchen    4096 05-10 18:27 .
drwxr-xr-x 7 root  root     4096 03-05 00:06 ..
drwxr-xr-x 3 hchen hchen    4096 03-01 18:13 backup
-rw------- 1 hchen hchen     721 05-05 22:12 .bash_history
-rw-r--r-- 1 hchen hchen      24 02-28 22:20 .bash_logout
-rw-r--r-- 1 hchen hchen     176 02-28 22:20 .bash_profile
-rw-r--r-- 1 hchen hchen     124 02-28 22:20 .bashrc
-rw-r--r-- 1 root  root    14002 03-07 00:29 index.htm
-rw-r--r-- 1 hchen hchen   31465 03-01 23:48 index.php

```

**7）ldd**

这个命令可以知道你的一个可执行文件所使用了动态链接库。如：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# ldd /usr/bin/java
        linux-gate.so.1 =>  (0x00cd9000)
        libgij.so.7rh => /usr/lib/libgij.so.7rh (0x00ed3000)
        libgcj.so.7rh => /usr/lib/libgcj.so.7rh (0x00ed6000)
        libpthread.so.0 => /lib/i686/nosegneg/libpthread.so.0 (0x00110000)
        librt.so.1 => /lib/i686/nosegneg/librt.so.1 (0x009c8000)
        libdl.so.2 => /lib/libdl.so.2 (0x008b5000)
        libz.so.1 => /usr/lib/libz.so.1 (0x00bee000)
        libgcc_s.so.1 => /lib/libgcc_s.so.1 (0x00aa7000)
        libc.so.6 => /lib/i686/nosegneg/libc.so.6 (0x0022f000)
        libm.so.6 => /lib/i686/nosegneg/libm.so.6 (0x00127000)
        /lib/ld-linux.so.2 (0x00214000)

```

**8）col**

这个命令可以让你把man文件转成纯文本文件。如下示例：

```shell
# PAGER=cat
# man less | col -b > less.txt

```

**9）xmlwf**

这个命令可以让你检查一下一个XML文档是否是所有的tag都是正常的。如：

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# curl 'https://coolshell.cn/?feed=rss2' > cocre.xml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 64882    0 64882    0     0  86455      0 --:--:-- --:--:-- --:--:-- 2073k
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# xmlwf cocre.xml
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# perl -i -pe 's@<link>@<br>@g' cocre.xml
[hchen@RHELSVR5 ~]# xmlwf cocre.xml
cocre.xml:13:23: mismatched tag

```

**10）lsof**

可以列出打开了的文件。

```shell
[[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) ~]# lsof | grep TCP
httpd       548    apache    4u     IPv6   14300967    TCP *:http (LISTEN)
httpd       548    apache    6u     IPv6   14300972    TCP *:https (LISTEN)
httpd       561    apache    4u     IPv6   14300967    TCP *:http (LISTEN)
httpd       561    apache    6u     IPv6   14300972    TCP *:https (LISTEN)
sshd       1764      root    3u     IPv6       4993    TCP *:ssh (LISTEN)
tcpserver  8965      root    3u     IPv4  153795500    TCP *:pop3 (LISTEN)
mysqld    10202     mysql   10u     IPv4   73819697    TCP *:mysql (LISTEN)
sshd      10735      root    3u     IPv6  160731956    TCP 210.51.0.232:ssh->123.117.239.68:31810 (ESTABLISHED)
sshd      10767     hchen    3u     IPv6  160731956    TCP 210.51.0.232:ssh->123.117.239.68:31810 (ESTABLISHED)
vsftpd    11095      root    3u     IPv4  152157957    TCP *:ftp (LISTEN)

```

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
