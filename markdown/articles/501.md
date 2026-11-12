# Ubuntu的并行启动

如果你在使用多核处理器，那么你可以配置Ubuntu的一个参数来使用Ubuntu的启动并行，以加快其启动速度。

这个参数在文件*/etc/init.d/rc*中，其参数名是CONCURRENCY默认值是none，你可以把这个参数改成如下所示。于是，你就开启了Ubuntu的并行启动的功能。

CONCURRENCY=shell

但是，这个参数会导致dbus, hal和gdm的产生“race condition”竞争条件，所以，这三个程序的启动顺序非常的关键。其必需保证这个顺序：dbus -> hal -> gdm。这个顺序在Ubuntu的Hardy，Intrepid 或Jaunty中是没有问题的。但是，我们不排除在别的版本中会有问题。

所以，在开启“并行启动”时，你需要去检查一下dbus，hal和gdm的启动顺序，其启动顺序你可以在*/etc/rc2.d/*目录下，查看一个这三个程序的S后面的编号顺序。如果你看到下面的这个顺序，那么你就需要做出调整了。

s12dbus  

s13gdm  

s24hal

调整也很简单，就是改一下S后面的数字就行了，如下所示：

mv s24hal s13hal  

mv s13gdm s14gdm

关于更多详细的情况，请查看这个[BUG报告](https://bugs.launchpad.net/ubuntu/+source/hal/+bug/149881) 。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
