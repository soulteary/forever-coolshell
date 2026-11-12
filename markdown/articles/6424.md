# Hash Collision DoS 问题

最近，除了国内明文密码的安全事件，还有一个事是比较大的，那就是 Hash Collision DoS （Hash碰撞的拒绝式服务攻击），有恶意的人会通过这个安全弱点会让你的服务器运行巨慢无比。**这个安全弱点利用了各语言的Hash算法的“非随机性”可以制造出N多的value不一样，但是key一样数据，然后让你的Hash表成为一张单向链表，而导致你的整个网站或是程序的运行性能以级数下降（可以很轻松的让你的CPU升到100%）**。目前，这个问题出现于[Java](http://www.java.com/), [JRuby](http://jruby.org/), [PHP](http://www.php.net/), [Python](https://python.org/), [Rubinius](http://rubini.us/), [Ruby](http://www.ruby-lang.org/)这些语言中，主要：

- [Java](http://www.java.com), 所有版本
- [JRuby](http://jruby.org/) <= 1.6.5 （目前fix在 1.6.5.1）
- [PHP](http://www.php.net/) <= 5.3.8, <= 5.4.0RC3 （目前fix在 5.3.9,  5.4.0RC4）
- [Python](https://python.org/), all versions
- [Rubinius](http://rubini.us/), all versions
- [Ruby](http://www.ruby-lang.org/) <= 1.8.7-p356 （目前fix在 1.8.7-p357, 1.9.x）
- [Apache Geronimo](http://geronimo.apache.org/), 所有版本
- [Apache Tomcat](http://tomcat.apache.org/) <= 5.5.34, <= 6.0.34, <= 7.0.22 （目前fix在 5.5.35,  6.0.35,  7.0.23）
- [Oracle Glassfish](http://glassfish.java.net/) <= 3.1.1 （目前fix在mainline）
- [Jetty](http://www.eclipse.org/jetty/), 所有版本
- [Plone](http://plone.org/), 所有版本
- [Rack](http://rack.rubyforge.org/) <= 1.3.5, <= 1.2.4, <= 1.1.2 （目前fix 在 1.4.0, 1.3.6, 1.2.5, 1.1.3）
- [V8 JavaScript Engine](https://code.google.com/p/v8/), 所有版本
- ASP.NET 没有打MS11-100补丁

注意，Perl没有这个问题，因为Perl在N年前就fix了这个问题了。关于这个列表的更新，请参看 [oCERT的2011-003报告](http://www.ocert.org/advisories/ocert-2011-003.html)，比较坑爹的是，这个问题早在2003 年就在论文《[通过算法复杂性进行拒绝式服务攻击](http://www.cs.rice.edu/~scrosby/hash/CrosbyWallach_UsenixSec2003.pdf)》中被报告了，但是好像没有引起注意，尤其是Java。

#### 弱点攻击解释

你可以会觉得这个问题没有什么大不了的，因为黑客是看不到hash算法的，如果你这么认为，那么你就错了，这说明对Web编程的了解还不足够底层。

无论你用JSP，PHP，Python，Ruby来写后台网页的时候，在处理HTTP POST数据的时候，你的后台程序可以很容易地以访问表单字段名来访问表单值，就像下面这段程序一样：

```php


$usrname = $_POST['username'];
$passwd = $_POST['password'];


```

这是怎么实现的呢？这后面的东西就是Hash Map啊，所以，我可以给你后台提交一个有10K字段的表单，这些字段名都被我精心地设计过，他们全是Hash Collision ，于是你的Web Server或语言处理这个表单的时候，就会建造这个hash map，于是在每插入一个表单字段的时候，都会先遍历一遍你所有已插入的字段，于是你的服务器的CPU一下就100%了，你会觉得这10K没什么，那么我就发很多个的请求，你的服务器一下就不行了。

举个例子，你可能更容易理解：

如果你有n个值—— v1, v2, v3, … vn，把他们放到hash表中应该是足够散列的，这样性能才高：

> 0 -> v2  
> 
> 1 -> v4  
> 
> 2 -> v1  
> 
> …  
> 
> …  
> 
> n -> v(x)

但是，这个攻击可以让我造出N个值——  dos1, dos2, …., dosn，他们的hash key都是一样的（也就是Hash Collision），导致你的hash表成了下面这个样子：

> 0 – > dos1 -> dos2 -> dos3 -> …. ->dosn  
> 
> 1 -> null  
> 
> 2 -> null  
> 
> …  
> 
> …  
> 
> n -> null

于是，单向链接就这样出现了。这样一来，O(1)的搜索算法复杂度就成了O(n)，而插入N个数据的算法复杂度就成了O(n^2)，你想想这是什么样的性能。

（关于Hash表的实现，如果你忘了，那就把大学时的《数据结构》一书拿出来看看）

#### Hash Collision DoS 详解

StackOverflow.com是个好网站， 合格的程序员都应该知道这个网站。上去一查，就看到了这个贴子“[Application vulnerability due to Non Random Hash Functions](http://stackoverflow.com/questions/8669946/application-vulnerability-due-to-non-random-hash-functions "Application vulnerability due to Non Random Hash Functions")”。我把这个贴子里的东西摘一些过来。

首先，这些语言使用的Hash算法都是“非随机的”，如下所示，这个是Java和Oracle使用的Hash函数：

```c
static int hash(int h)
{
h ^= (h >>> 20) ^ (h >>> 12);
return h ^ (h >>> 7) ^ (h >>> 4);
}
```

所谓“非随机的” Hash算法，就可以猜。比如：

1）在Java里， Aa和BB这两个字符串的hash code(或hash key) 是一样的，也就是Collision 。

2）于是，我们就可以通过这两个种子生成更多的拥有同一个hash key的字符串。如：”AaAa”, “AaBB”, “BBAa”, “BBBB”。这是第一次迭代。其实就是一个排列组合，写个程序就搞定了。

3）然后，我们可以用这4个长度的字符串，构造8个长度的字符串，如下所示：

```
"AaAaAaAa", "AaAaBBBB", "AaAaAaBB", "AaAaBBAa", 
"BBBBAaAa", "BBBBBBBB", "BBBBAaBB", "BBBBBBAa", 
"AaBBAaAa", "AaBBBBBB", "AaBBAaBB", "AaBBBBAa", 
"BBAaAaAa", "BBAaBBBB", "BBAaAaBB", "BBAaBBAa",
```

`4）同理，我们就可以生成16个长度的，以及256个长度的字符串，总之，很容易生成N多的这样的值。`

在攻击时，我只需要把这些数据做成一个HTTP POST 表单，然后写一个无限循环的程序，不停地提交这个表单。你用你的浏览器就可以了。当然，如果做得更精妙一点的话，把你的这个表单做成一个跨站脚本，然后找一些网站的跨站漏洞，放上去，于是能过SNS的力量就可以找到N多个用户来帮你从不同的IP来攻击某服务器。

#### 防守

要防守这样的攻击，有下面几个招：

- 打补丁，把hash算法改了。
- 限制POST的参数个数，限制POST的请求长度。
- 最好还有防火墙检测异常的请求。

不过，对于更底层的或是其它形式的攻击，可能就有点麻烦了。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
