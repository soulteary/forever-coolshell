# Java如何取源文件中文件名和行号

如何取的Java源代码文件中文件名和行号：）

在C/C++的程序，编译器提供了两个宏来支持取得源文件中的行号和文件名，这两个宏是\_\_FILE\_\_,\_\_LINE\_\_

你可以如下的方法打印行号和文件名

```c


#include <stdio.h>
int main()
{
 fprintf(stdout,"[%s:%d] Hello World!",__FILE__,__LINE__);
 return 0;
}


```

  

但是在JAVA下没有这两个宏，那么我们如何来取得文件名和行号，翻阅JDK，我们找到StackTraceElement这个类。这个类可以从Throwable取得，另外也可以从Thread类取得，通过这些我写如下的一个打印行号的测试程序：

```java


public class LineNo {
 public static int getLineNumber() {
 return Thread.currentThread().getStackTrace()[2].getLineNumber();
 }  

 public static String getFileName() {
 return Thread.currentThread().getStackTrace()[2].getFileName();
 }
 public static void main(String args[]) {
 System.out.println("["+getFileName()+"："+ getLineNumber()+"]"+"Hello World!");
 }
}


```

留下一个问题，上面程序中的magic数字 2 代表什么含义呢？

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
