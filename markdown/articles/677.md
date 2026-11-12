# 用Python写NCurses UI

Ncurses是一个能提供基于文本终端窗口功能的动态库. Ncurses可以:

- 可以使用整个屏幕
- 创建和管理一个窗口
- 使用8种不同的彩色
- 为您的程序提供鼠标支持
- 使用键盘上的功能键

Ncurses可以在任何遵循ANSI/POSIX标准的Unix/Linux系统上运行，除此之外，它还可以从系统数据库中检测终端的属性,，并且自动进行调整,提供一个不受终端约束的接口。因此,Ncurses可以在不同的系统平台和不同的终端上工作的非常好。

mc工具集就是一个用ncurses写的很好的例子,而且在终端上系统核心配置的界面同样是用ncurses编写的. 下面就是它们的截图：

[![ncurses_example](https://coolshell.cn/uploads/2009/04/ncurses_example.jpg "ncurses_example")](https://coolshell.cn/wp-content/uploads/2009/04/ncurses_example.jpg)

当然，在我们这篇文章中，我们不会教你怎么写NCurses程序，我们只是想告诉你如何用Python来写Ncurses的程序，示例会非常简单，点到为止。

在此之前，我们先简单的回顾一下如何使用Python的一些简单的语法。

先看看一个最简单的Python程序：

```python
print "How easy is this?" 

x = 1
y = 2
z = x + y

print "Result of x + y is", z

```

程序很简单，我就不多说，把这个文件存成test.py，然后在命令行下调用python test.py就可以看到输出了。

下面我们再来看一个Python的函数定义——还是很简单，我也不用多说了。

```python
def saystuff(mystring):
     print "You said:", mystring 

saystuff("Bach rules")
saystuff("So does Telemann")

```

好，言归正传，让我们来看一下，如何在Python中使用NCurses，下面是一个小例程：

```python
import curses 

myscreen = curses.initscr()

myscreen.border(0)
myscreen.addstr(12, 25, "Python curses in action!")
myscreen.refresh()
myscreen.getch()

curses.endwin()

```

注意这个示例中的第一行import curses，表明使用curses库，然后这个示像在屏幕中间输出“Python curses in action!”字样，其中坐标为12, 25，注意，在字符界面下，80 x 25是屏幕大小，其用的是字符，而不是像素。下面是运行后的抓屏：

[![python_ncursespy](https://coolshell.cn/uploads/2009/04/python_ncursespy.jpg "python_ncursespy")](https://coolshell.cn/wp-content/uploads/2009/04/python_ncursespy.jpg)

最后，我们再来看一个数字菜单的示例：

```python
#!/usr/bin/env python

from os import system
import curses

def get_param(prompt_string):
     screen.clear()
     screen.border(0)
     screen.addstr(2, 2, prompt_string)
     screen.refresh()
     input = screen.getstr(10, 10, 60)
     return input

def execute_cmd(cmd_string):
     system("clear")
     a = system(cmd_string)
     print ""
     if a == 0:
          print "Command executed correctly"
     else:
          print "Command terminated with error"
     raw_input("Press enter")
     print ""

x = 0

while x != ord('4'):
     screen = curses.initscr()

     screen.clear()
     screen.border(0)
     screen.addstr(2, 2, "Please enter a number...")
     screen.addstr(4, 4, "1 - Add a user")
     screen.addstr(5, 4, "2 - Restart Apache")
     screen.addstr(6, 4, "3 - Show disk space")
     screen.addstr(7, 4, "4 - Exit")
     screen.refresh()

     x = screen.getch()

     if x == ord('1'):
          username = get_param("Enter the username")
          homedir = get_param("Enter the home directory, eg /home/nate")
          groups = get_param("Enter comma-separated groups, eg adm,dialout,cdrom")
          shell = get_param("Enter the shell, eg /bin/bash:")
          curses.endwin()
          execute_cmd("useradd -d " + homedir + " -g 1000 -G " + groups + " -m -s " + shell + " " + username)
     if x == ord('2'):
          curses.endwin()
          execute_cmd("apachectl restart")
     if x == ord('3'):
          curses.endwin()
          execute_cmd("df -h")

curses.endwin()

```

下面是运行时的显示：

[![python_ncurses](https://coolshell.cn/uploads/2009/04/python_ncurses.jpg "python_ncurses")](https://coolshell.cn/wp-content/uploads/2009/04/python_ncurses.jpg)

如果你你了解NCurses编程，你可以看看相关的Linux HOW-TO的文章，链接在这里：[Linux Documentation Project’s NCURSES Programming How To](http://www.linux.org/docs/ldp/howto/NCURSES-Programming-HOWTO/index.html)

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
