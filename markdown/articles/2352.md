# telnet的一个Bug

下面这个链接是Linux分发包Ubuntu的关于Telnet命令的Man Page，

<http://manpages.ubuntu.com/manpages/karmic/man1/telnet-ssl.1.html>

打开这个Man Page，把页面拉到最后一行，你会看到下面这个BUG（“BUGS：源代码不易读！”）

```
     The source code is not comprehensible.
```

Telnet的源代码在这里：<http://packages.ubuntu.com/source/dapper/netkit-telnet>，下载下来一看，还真是不易读，简单地看了一下代码，发现至少有这样一些问题：

- 空格和Tab键混用的缩进，导致很多代码在没有缩进。
- 大量的#if #else以及大量的各种预编译宏。以及一些怪异的宏。如：

#ifndef B19200  

#define B19200 B9600  

#endif

#ifndef B38400  

#define B38400 B19200  

#endif

- 什么叫在C中写C++，第一次见。（在terminal.cc中间居然出现了几个class）
- 变量命名很不直观，大量的old, tmp, c1, c2, s1, s2, s3 等学校里用的变量名，只有作者自己知道是什么意思。函数命令的风格也不一致，编程风格也很不一致，基本没有编程规范。

的确很不易读。不管怎么样，很欣赏在man page中把源码的易读性列为BUG的这种作法。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
