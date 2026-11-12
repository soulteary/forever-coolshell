# Hello World 集中营

编程的人应该都知道什么是Hello World。这是一个最简单的程序，其只在屏幕上输出“Hello World”字样，这通常是初学者的在学习编程时的第一个示例。把打印出 “Hello World” 作为第一个范例程序，现在已经成为编程语言学习的传统。  

![hello_world](https://coolshell.cn/uploads/2009/03/hello_world-150x150.png "hello_world")  

“Hello World”起源于Brian Kernighan 和Dennis MacAlistair Ritchie写的计算机程序设计教程《C语言程序设计》（[*The C Programming Language*](https://en.wikipedia.org/wiki/The_C_Programming_Language "en:The C Programming Language")）而广泛流传；但这本书并不是 “hello, world” 的滥觞，虽然这是一个普遍存在的错误认知。

这范例程序最早出现于 1972 年，由贝尔实验室成员 Brian Kernighan 撰写的内部技术文件《Introduction to the Language B》之中。不久同作者于 1974 年所撰写的《Programming in C: A Tutorial》，也延用这个范例；而以本文件扩编改写的《C语言程序设计》也保留了这个範例程式。

“hello, world” 程序的标准打印内容必须满足“全小写，无惊叹号，逗点后需空一格”，不过流传至今，完全恪守传统的反而罕见。

下面我们来看几个例子：

```
#include <stdio.h>

int main(void)
{
   printf("Hello, world!n");
   return 0;
}

```
```
#include <iostream>
using namespace std;

int main()
{
    cout << "Hello, world!" << endl;
    return 0;
}

```
```
public class Hello
{
    public static void main(String[] args)
    {
        System.out.println("Hello, world!");
    }
}

```

不过，最全的Hello World的集中营在这里：（请大家围观这个网页）

[**http://www.roesler-ac.de/wolfram/hello.htm**](http://www.roesler-ac.de/wolfram/hello.htm)

这个网站很BT啊，其开始是从1994年10月3日，于1999年12月30日上互联网，2005年7月14日收集到了超过200个不同语言的Hello World，2006年12月6日达到300个，2008年2月27日达到400个。

今天这个网站有一共421个不同语言的Hello World，其中有63个来自人类的语言。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
