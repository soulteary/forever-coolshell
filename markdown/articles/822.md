# Linux磁盘使用命令du的改进

我们知道，在Linux下，如果你想知道当前目录下，每个文件或子目录的尺寸，你可以使用du命令来完成这一动作。如：

```shell
$  du -sh *

```

这个命令可以以K，M，G的方式显示每个文件和子目录的大小。我们把这种方式叫做，human-readable，也就是可以让人读的方式，如下所示：

```
8.4G Desktop
2.6G Documents
12K keys
12M Pictures
536K scripts
```

  

但是，很可惜的是，我们的du并没有提供相关的排序功能，所以，如果在human-readable下，也就是-h参数下，我们很难使用sort命令来排序。因为那变成了字符串排序，小数点，数字的位数，还有单位K，M，G都会让排序变得混乱。那么，我们如何才能即有human-readble这种功能，还能有排序呢。我们得借用一些脚本语言来处理了。

下面是使用了Perl来达到这一功能：

```shell
du -sk * | sort -n |       #以 K 字节的方式排序
perl -ne '                 #使用Perl来处理 K M 和 G 单位
  ($s,$f)=split(m{\t});    #把 尺寸/文件名 分开
  for (qw(K M G)) {        #以尺寸单位循环
     if($s<1024) {         #如果 尺寸<1024 那么就输出
       printf("%.1f",$s);  #显示文件尺寸
       print "$_\t$f";     #显示文件名
       last                #换行
     };
     $s=$s/1024            #除1024然后进入下一个尺寸单位
  }
'

```

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
