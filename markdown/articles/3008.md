# Windows编程革命简史

源文：[A Brief History of Windows Programming Revolutions](http://www.drdobbs.com/windows/225701475) （Ron Burk）

首先，是 Windows API 和 [DLL Hell](https://en.wikipedia.org/wiki/DLL_hell)。（译注：DLL Hell——DLL灾难，就是微软的DLL升级时因为不同版本可能造成应用程序无法运行的灾难，首当其冲的是COM编程，相信大家都知道某些木马或是病毒更改了一些系统的DLL可以导致整个Windows不举，这就是DLL Hell） 于是，第一次革命是[DDE](https://en.wikipedia.org/wiki/Dynamic_Data_Exchange)——我们可以创建一个状态条在上面显示Microsoft的股票价格（译注：Dynamic Data Exchange，工作原理是： 甲方申请一块全局内存，然后把内存指针postmessage到乙方，乙方根据收到的指针访问那块全局内存）。

在那个时候，Microsoft 创建了 VERSIONINFO 资源来管理版本信息，当然，是用来消除DLL Hell。但是，另一个微软内部的小组发现了DDE的致命缺陷：这不是他们做的！

为了解决这个问题，他们创造了OLE（很像DDE，只是名字不一样），而且，我还记得在一次 Microsoft 大会上，某个微软的演讲者正式宣布—— Windows API 马上就会被 OLE API 所重写并取代，我还盲目地相信了这一说法。而且，所有的在图形界面的控件都会是OCX，那是OLE引入的接口，同样，其目的是为了消除DLL Hell。相信大家都记得，那个时候，我们是怎么地梦想着有一天，我们的应用程序（当然是非常大的程序）可以完全地被嵌入到Word文档中。

然而，在Microsoft的某处，Microsoft有些人开始信仰 C++，其确信MFC的出现并可以解决所有的一切问题，但是，因为历史原因，OLE并没有出局，其改了一个名字，叫COM，此时，我们立马意识到OLE（以前的DDE？）真正意味着什么——其用精心的版本管理系统来消除DLL Hell。与此同时，Microsoft的一个变节小组发现了一个MFC的致命缺陷：这不是他们做的！

当然，微软件的动作是很快的，他们立刻修正了问题——创造了ATL，有些像MFC，只是名字不同，他们想使用漂亮的ATL把那些晦涩难懂的COM的知识给隐藏住。这个动作刺激了COM团队（或是OLE团队？）改名为ActiveX，并发布了成千上万的新接口（甚至是很多版本化的接口，当然，主要目的是为了消除DLL Hell），当然，ActiveX可以让我们的程序可以从Web游览器上下载，并可以完美地和病毒一起嵌入浏览器中（哈，还不紧跟时代，感谢ATL吧）。此时，操作系统团队就像一个失宠的孩子一样，大声呼喊着“[Cairo操作系统](http://www.microsoft.com/middleeast/egypt/cmic/)来了”引起大家注意，当然有一些怪异恶心的东西连他们自己也无法解释清楚，所以，别提发布了。为了声誉，操作系统团队的确引入了“系统文件保护”的理念，当然也是为了消除DLL Hell。

这个时候，Microsoft的某个团队发现了Java的致命缺陷：这不是他们做的！于是他们创造了一个叫J，或是Jole，或是ActiveJ的东西（对不起，我真的记不起叫什么了）来挽救Java（译者：应该是Visual J++）。看起来很像Java，只是名字不同罢了。这太让人兴奋了，但是Sun使用了一些相当古老的法律条款向Microsoft提起了法律诉讼，其在一年内限制了任何一个公司可以发布类似Java的产品。这明显是抑制微软复制别人产品的一次尝试，唯一不同的，其结果导致了微软流向国会议员裤兜的现金网络的建立（在这个网络可以得到时事新闻和价值$14.75的T恤衫）。还记得 J/Jole/ActiveJ 的项目经理用他的鞋桌在敲着桌子并信誓旦旦地坚持 Microsoft 将永远不会放弃他的产品。SB！所有的这些也就仅仅意味着一件事——没有人关心ActiveX团队（或者是COM？）。令人难以置信的是，微软把这些东东全部集成起来，成了COM+（难道不应该是ActiveX+?），还有MTS（我不知道为什么没有COM和Active或是X或是+的字眼，而直接叫MTS了——我为这个名词感到实实在在地震惊！）。他们总是那么NB地为那些流行词加上“+”号。在那段时间，还有人曾叫喊着“Windows DNA”以及“Windows Washboard”，但这两个东西最终在我搞清是什么玩意的之前就夭折了。

在这一点上，Microsoft已经很不安地窥视着Internet好几年了，他们终于意识到Internet上有一个致命缺陷：嗯，你应该知道这是什么（译注：Internet不是做他们做的！）。于是他们开始培养我们和.NET约会（.NET的发音很像“doughnut”圆环图，不过，这只是他们的唯一不同），这和Internet很相似，只不过.NET有更多的印刷品。其让我们清楚再清楚地了解一件事：.NET会消除DLL Hell。.NET包含了一个新的编程语言，叫C#（为了解决已经死翘翘的Active++ J++的缺陷）。.NET还包含一个虚拟机，所有的语言都运行在上面（这主要是为了解决依赖于Intel CPU的缺陷）。.NET还包含了一个单一的登录系统（这主要是为了解决“不把口令存放在Microsoft服务器上”的缺陷）。实际上，我们更容易做的是把.NET不包含的事给列出来。.NET绝对是一个划时代地Windows编程革命……当然，仅到明年。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
