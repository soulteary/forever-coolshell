# chmod -x chmod的N种解法

在SlidesShare.net上有这么[一个幻灯片](http://www.slideshare.net/cog/chmod-x-chmod)，其说了如下的一个面试题：

> 如果某天你的Unix/Linux系统上的chomd命令被某人去掉了x属性（执行属性），  
> 
> 那么，你如何恢复呢？

下面是一些答案：

**1）重新安装**。对于Debian的系统：

`sudo apt-get install --reinstall coreutils`

**2）使用语言级的chmod**。

- Perl：perl-e ‘chmod 0755, “/bin/chmod”‘
- Python：python -c “import os;os.chmod(‘/bin/chmod’, 0755)”
- Node.js：require(“fs”).chmodSync(“/bin/chmod”, 0755);
- C程序：

```c
#include <sys/types.h>
#include<sys/stat.h>
void main()
{
chmod("/bin/chmod", 0000755);
}
```

**3）使用已有的可执行文件。**

```shell
$cat - > chmod.c
void main(){}
^D

$cc chmod.c
$cat /bin/chmod > a.out
$./a.out 0755 /bin/chmod

```
```shell
$cp true > new_chmod
$cat /bin/chmod > new_chmod
$./new_chmod 0755 /bin/chmod

```

**4）使用GNU tar命令**

```shell
$tar --mode 0755 -cf chmod.tar /bin/chmod
$tar xvf chmod.tar
```

`tar --mode 755 -cvf - chmod | tar -xvf -`

**5）使用cpio** （第19到24字节为file mode – <http://4bxf.sl.pt>）

```shell
echo chmod |
cpio -o |
perl -pe 's/^(.{21}).../${1}755/' |
cpio -i -u
```

**6）使用hardcore**

`alias chmod='/lib/ld-2.11.1.so ./chmod'`

**7）使用Emacs**

> Ctrl+x b > \* scratch\*  
> 
> (set-file-modes “/bin/chmod” (string-to-number “0755” 8))  
> 
> Ctrl+j

嗯，挺强大的，不过为什么不用install命令呢？

```shell
install -m 755 /bin/chmod /tmp/chmod
mv /tmp/chmod /bin/chmod
```

各位，你的方法呢？

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
