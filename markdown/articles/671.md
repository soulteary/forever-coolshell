# Python调用C语言函数

使用Python的[ctypes](https://docs.python.org/library/ctypes.html)，我们可以直接调用由C直接编译出来的函数。其实就是调用动态链接库中的函数。为什么我们需要这样做呢，因为有些时候，我们可能需要一个性能上比较讲究的算法，有些时候，我们可以在Python中使用已经有了的现成的被封闭在动态链接库中的函数。下面是如何调用的示例。

首先，我们用一个乘法来表示一个算法功能。下面是C的程序：

```c
int
multiply(int num1, int num2)
{
    return num1 * num2;
}

```

如果在Windows下，你可能需要写成下面这个样子：

```c
#include <windows.h>


BOOL APIENTRY
DllMain(HANDLE hModule, DWORD dwReason, LPVOID lpReserved)
{
    return TRUE;
}

__declspec(dllexport) int
multiply(int num1, int num2)
{
    return num1 * num2;
}

```

然后，自然是把这个C文件编成动态链接库：

Linux下的编译：

```shell
gcc -c -fPIC libtest.c
gcc -shared libtest.o -o libtest.so

```

Windows下的编译：

```shell
cl -LD libtest.c -libtest.dll

```

于是在我们的Python中可以这样使用：  

(其中的libtest.so在Windows下改成libtest.dll即可)

```python
>>> from ctypes import *
>>> import os
>>> libtest = cdll.LoadLibrary(os.getcwd() + '/libtest.so')
>>> print libtest.multiply(2, 2)
4

```

注意：上面的Python脚本中需要把动态链接库放到当前目录中。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
