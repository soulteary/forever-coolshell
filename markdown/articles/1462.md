# CentOS上php的问题及Selinux安全设置

最近有位站长在用我们WebIM客户端的时候，无法登录我们的WebIM服务器，十分惊讶。 在我们的用户里尚属首例，其实更惊讶的是我的CentOS也遇到了同样的问题。然后分析了这位站长的HttpResponse , Shamee :( 一样的OS.

搜了一下，发现的解决方法都是在代码上。 我想可能关键词有错误，因为我坚信我的问题肯定不在代码上，应该是来自OS本身的限制。于是重新debug了一下代码，报错 permission (13) connection。然后直接在洋人的邮件列表里搜了一下。

问题确定了 是SeLinux(*http://zh.wikipedia.org/wiki/SELinux*)安全策略的限制。

这下问题明了了,执行 /usr/sbin/setenforce 0就能迅速关闭SELINUX，或者vi /etc/selinux/config 把enforcing改成permissive 然后reboot.

但是我想了一下，就算安全级别为B1的Linux被攻击的可能小，但是总会有面对这种问题的时候，况且这种解决访问本身并不优雅。

于是想了下 把Apache脱离SeLinux是一个最恰当的办法，于是执行

`sudo  setsebool -P httpd_disable_trans 1 && sudo   /etc/init.d/httpd restart`

这样就能保证在SeLinux的光环下,Web服务器行为不受控制。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
