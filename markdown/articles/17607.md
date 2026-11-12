# 从 MongoDB “赎金事件” 看安全问题

![](https://coolshell.cn/uploads/2017/01/MongoDB-360x200.jpg)今天上午（2017年1月7日），我的微信群中同时出现了两个MongoDB被黑掉要赎金的情况，于是在调查过程中，发现了这个事件。这个事件应该是2017年开年的第一次比较大的安全事件吧，发现国内居然没有什么报道，国内安全圈也没有什么动静（当然，他们也许知道，只是不想说吧），Anyway，让我这个非安全领域的人来帮补补位。

#### 事件回顾

这个事情应该是从2017年1月3日进入公众视野的，是由安全圈的大拿 Victor Gevers （网名：[0xDUDE](https://twitter.com/0xDUDE)，[GDI.foundation](http://GDI.foundation "http://GDI.foundation") 的Chairman），其实，他早在2016年12月27日就发现了一些在互联网上用户的MongoDB没有任何的保护措施，被攻击者把数据库删除了，并留下了一个叫 WARNING 的数据库，这张表的内容如下：

```
{
    "_id" : ObjectId("5859a0370b8e49f123fcc7da"),
    "mail" : "[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)",
    "note" : "SEND 0.2 BTC TO THIS ADDRESS 13zaxGVjj9MNc2jyvDRhLyYpkCh323MsMq AND CONTACT THIS EMAIL WITH YOUR IP OF YOUR SERVER TO RECOVER YOUR DATABASE !"
}
```

基本上如下所示：

![MongoDB ransom demand (via Victor Gevers)](https://coolshell.cn/uploads/2017/01/MongoDB-ransom.png)

MongoDB ransom demand (via Victor Gevers)

说白了就是黑客留下的东西——**老子把你的MongoDB里的数据库给转走了，如果你要你的数据的话，给我0.2个的比特币（大约USD200）**。然后，他的twitter上不断地发布这个“赎金事件”的跟踪报道。与此同时，中国区的V2EX上也发现了相关的攻击问题 《[自己装的 mongo 没有设置密码结果被黑了](https://www.v2ex.com/t/331887)》

然后，在接下来的几天内，全球大约有1800个MongoDB的数据库被黑，这个行为来自一个叫 Harak1r1 的黑客组织（这个组织似乎就好黑MongoDB，据说他们历史上干了近8500个MongoDB的数据库，几乎都是在祼奔的MongoDB）。

不过，这个组织干了两天后就停手了，可能是因为这事已经引起了全球科技媒体的注意，产生了大量的报道（如果你在Google News里查一下“[mongodb ransom](https://www.google.com/webhp?sourceid=chrome-instant&ion=1&espv=2&ie=UTF-8#q=mongodb+ransom&newwindow=1&tbm=nws)”，你会看到大量的报道（中文社区中，只有[台湾有相关的报道](https://unwire.pro/2017/01/05/2000-mongodb-ransom/security/)）），他们也许是不敢再搞下去了。

不过，很快，有几个copycats开始接着干，

马上跟进的是 own3d ，他们留下的数据库的名字叫 WARNING\_ALERT，他们至少干掉了 930个MongoDB，赎金0.5个比特币（USD500），至少有3个用户付费了

然后是0704341626asdf，他们留下的数据库名字叫PWNED，他们至少干掉了740个MongoDB，赎金0.15个比特币（USD150），看看他们在数据库里留下的文字——**你的MongoDB没有任何的认证，并且暴露在公网里（你TMD是怎么想的？）……**

![0704341626asdf group ransom note (via Victor Gerves)](https://coolshell.cn/uploads/2017/01/MongoDB-Group-3.jpg)

0704341626asdf group ransom note (via Victor Gerves)

就在这两天，有两个新的黑客也来了

- 先是kraken0，发现到现在1天了，干了13个MongoDB，赎金 0.1个比特币。
- 然后是 3lix1r，发现到现在5个小时，干了17个MongoDB，赎金0.25比特币。

BBC新闻也于昨天报道了这一情况——《[Web databases hit in ransom attacks](http://www.bbc.com/news/technology-38521973)》，现在这个事情应该是一个Big News了。

#### 关于MongoDB的安全

安全问题从来都是需要多方面一起努力，但是安全问题最大的短板就是在用户这边。这次的这个事，说白了，就是用户没有给MongoDB设置上用户名和口令，然后还把服务公开到了公网上。

是的，这个安全事件，相当的匪夷所思，为什么这些用户要在公网上祼奔自己的数据库？他们的脑子是怎么想的？

让我们去看一下Shodan上可以看到的有多少个在暴露在公网上而且没有防范的MongoDB？我了个去！**4万7千个，还是很触目惊心的**（下图来自我刚刚创建的 [Shodan关于MongoDB的报表](https://www.shodan.io/report/h0bgF6zM)）

![](https://coolshell.cn/uploads/2017/01/MongoDB_Shodan-1024x485.png)

那么，怎么会有这么多的对外暴露的MongoDB？看了一下Shodan的报告，发现主要还是来自公有云平台，Amazon，Alibaba，Digital Ocean，OVH，Azure 的云平台上有很多这样的服务。不过，像AWS这样的云平台，有很完善的默认安全组设置和VPC是可以不把这样的后端服务暴露到公有云上的，为什么还会有那么多？

![](https://coolshell.cn/uploads/2017/01/MongoDB_Org.png)

这么大量的暴露在公网上的服务是怎么回事？有人发现（参看这篇文章《[It’s the Data, Stupid!](https://blog.shodan.io/its-the-data-stupid/)》 ），MongoDB历史上一直都是把侦听端口绑在所有的IP上的，这个问题在5年前（2011年11月）就报给了MongoDB ([SERVER-4216](https://jira.mongodb.org/browse/SERVER-4216))，结果2014年4月才解决掉。所以，他觉得可能似乎 MongoDB的 2.6之前的版本都会默认上侦听在0.0.0.0 。

于是我做了一个小试验，到我的Ubuntu 14.04上去 `apt-get install mongodb`（2.4.9版），然后我在`/etc/mongodb.conf` 文件中，看到了默认的配置是127.0.0.1，mongod启动也侦听在了127.0.0.1这台机器上。一切正常。不过，可能是时过境迁，debain的安装包里已加上了这个默认配置文件。不管怎么样，MongoDB似乎是有一些问题的。

再到Shodan上看到相关的在公网裸奔的MongoDB的版本如下，发现3.x的也是主流：

![](https://coolshell.cn/uploads/2017/01/MongoDB_Version.png)

虽然，3.x的版本成为了主流，但是似乎，还是有很多人把MongoDB的服务开到了互联网上来，而且可以随意访问。

**你看，我在阿里云随便找了几台机器，一登就登上去了。**

![](https://coolshell.cn/uploads/2017/01/MongoDB_Aliyun.png)

真是如那些黑客中的邮件所说的：WTF，你们是怎么想的？

#### 后续的反思

为什么还是有这么多的MongoDB在公网上祼奔呢？难道有这么多的用户都是小白？这个原因，是什么呢？我觉得可能会是如下两个原因：

1）一是技术人员下载了mongod的软包，一般来说，mongodb的压缩包只有binary文件 ，没有配置文件 ，所以直接解开后运行，结果就没有安全认证，也绑在了公网上。也许，MongoDB这么做的原因就是为了可以快速上手，不要在环境上花太多的时间，这个有助于软件方面的推广。但是，这样可能就坑了更多的人。

2）因为MongoDB是后端基础服务，所以，需要很多内部机器防问，按道理呢，应该绑定在内网IP上，但是呢，可能是技术人员不小心，绑在了0.0.0.0的IP上。

那么，这个问题在云平台上是否可以更好的解决呢？

**关于公网的IP。**一般来说，公有云平台上的虚拟主机都会有一个公网的IP地址，老实说，这并不是一个好的方法，因为有很多主机是不需要暴露到公网上的，所以，也就不需要使用公网IP，于是，就会出现弹性IP或虚拟路由器以及VPC这样的虚拟网络服务，这样用户在公有云就可以很容易的组网，也就没有必要每台机器都需要一个公网IP，使用云平台，最好还是使用组网方案比较好的平台。

**关于安全组**。在AWS上，你开一台EC2，会有一个非常严格的安全组——只暴露22端口，其它的全部对外网关闭。这样做，其实是可以帮用户防止一下不小心把不必要的服务Open到公网上。按道理来说，AWS上应该是帮用户防了这些的。但是，AWS上的MongoDB祼奔的机器数量是最多的，估计和AWS的EC2的 基数有关系吧（据说AWS有千万台左右的EC2了）

最后，提醒大家一下，被黑了也不要去付赎金，因为目前来说没有任何证据证明黑客们真正保存了你的数据，因为，被黑的服务器太多了，估计有几百T的数据，估计是不会为你保存的。下面也是Victor Gevers的提示：

![](https://coolshell.cn/uploads/2017/01/MongoDB_Twitter.png)

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
