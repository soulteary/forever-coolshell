# xkcd 神图“Click and Drag”

[xkcd](http://xkcd.com/)对于经常浏览国外网站的朋友一定不会陌生。不过，还是先让我来介绍一下xkcd（[维基百科词条](https://en.wikipedia.org/wiki/Xkcd)）。这是一个漫画网站，它主要是发布一些很简单的随手画的漫画，它主要有四种体裁——浪漫、讽刺、数学 和 语言。也会经常出现一些和IT有关的漫画，比如下面这个漫画—— （懂Unix的人一眼就看懂了，不懂的怎么看也看不懂）

![](https://coolshell.cn/uploads/2012/10/xkcd-sandwich.png "xkcd-sandwich")

本质上来说，xkcd是一种Geek文化，里面的东西都非常的Geek和晦涩，讽刺很辛辣，但很多只有特定人群可以看得懂。而且表达的形式自由到天马行空，飘忽不定。

xkcd.com的网站创建者、所有的漫画的作者叫[Randall Munroe](https://en.wikipedia.org/wiki/Randall_Munroe "Randall Munroe")![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/f4/Randall_Munroe_ducks.JPG/230px-Randall_Munroe_ducks.JPG "Randall Munroe")，他以前在 NASA工作，是那里的Roboticist——机器人专家，80后，同样，也是一个程序员。他还会画漫画。

xkcd是他于2005年创建的，他本来只是想把他大学里在记事本里画的漫画放到他的个人主页上，但结果却搞成了一个独立的以漫画为主的网站，他用他画的这些漫画做成T恤卖。为什么要取名叫xkcd，据Munroe说，这四个字母，没有任何意义，就是为了让人不能把他们通过拼成一个单词读出来。现在他全职在搞xkcd.com。他现在一周会更新三次漫画，分别在周一，周三，和周五。

到了2007年5月，xkcd上的漫画才被广泛转载。2008年10月， *[The New Yorker](https://en.wikipedia.org/wiki/The_New_Yorker "The New Yorker")* 杂志对Munroe做了一个采访。

2010年3月，xkcd的书里的[谜底被解决了](http://forums.xkcd.com/viewtopic.php?p=2042913#p2042829)，Munroe在旧金山的金门大桥公园里给他的Fans发了255本限量版的书。

2012年4月1日愚人节，他的1037 号漫画(“Umwelt”) 会根据不同的IP，浏览器和地址显示不同的漫画。

2012年9月19号，xkcd的第1110号图问世了。

#### XKCD #1110 神图

这个图上面就是三格小漫画，一个小人拿着气球，还有两句耐人寻味的话。而**这三格漫画图的下面是一个风景图，取名 Click and Drag，也就是让你点住图片拖动。于是你就不能自拔了。**

我只所以在前面写了那么多东西，而不是把这个链接放在一开始，就是害怕你点了这个图，就再也不回来了。

好了，现在你可以点下面的链接开这个神图了 （你会发现这个图怎么也拖不完，无穷完尽的，所以，还请你先回来）

 **[Click and Drag](http://www.xkcd.com/1110/)**

**但请你一定还要回来，本文后面还有精彩内容!**

**这个图一发布，几乎全世界的各大论坛都在疯狂的转载，很多媒体都关注这个漫画，各种技术社区如：reddit 在疯狂地讨论着这个图是怎么实现的，有多大？还有很多人再分析这个图里的内容，这个图里隐藏着很多很有意思的东西，《有2001太空漫游》，有《星球大战》，还有《超级马丽》等等。**

**几乎整个互联网都沸腾了，但好像中国社区对此事完全不知。**

网上出现了很多相关的blog和站点来分析这个图片。如果你在Google里搜xkcd 1110，你会发现很多内容。

#### 这个图有多大

- 这个图可以分解成 2592 个 2048 x 2048 像素的图。

- 但其中只有 225 个 2048 x 2048 的PNG 图片文件。而剩下的2337 基本上是纯黑的或是纯白的块。比如地下和天空。

- 整个图横向有81个2048 x 2048的图（左边有33个，右边有48个），纵向有32个 2048 x 2048个图（天上有13个，地下有19个）

- 老大当晚Release的全尺寸的大图（比现在你看到的还要大），不算空白处，图片共有60G的像素，而如果要算上整个图将会是T级别的像素。现在你看到版本已被做过优化，不算空白处，只有1G的像素，而算上全图有10G的像素。 (2048x2048x225 = 943,718,400 和 2048x2048x2592 = 10,871,635,968).

- 如果我们按比例来看的话，图中的32个象素对应于现实世界的5英尺，那么，这个图的宽有25920英尺（7.9公里），高有10240英尺（3.1公里）。

- 如果每个 2048 x 2048 的PNG图可以被打印成一个300 dpi的宣传画，那么，这个宣传画基本上是14.05米宽，5.55米高的图。现在的PNG被调整过了，只有72dpi左右。

有人说，创作这么这个大图很费时间。不过我觉得这对于Geek来说不是问题，因为这应该是可以通过矢量图的拼装来搞定。

[![xkcd 1110全景缩略图（点击看大缩略图）](https://coolshell.cn/uploads/2012/10/xkcd1110-1024x346.png "xkcd 1110全景缩略图（点击看大缩略图）")](https://coolshell.cn/wp-content/uploads/2012/10/xkcd1110.png)

xkcd 1110全景缩略图（点击看大缩略图）

#### 看看技术宅们干了什么

下面我只记录了些不完全的技术宅们的因为这个画搞出来的东西。大家可以补充。

1）如果你用鼠标翻得不爽的话，你可以[看看这篇文章](http://www.potch.me/blog/press-and-hold.html)，在你的Chrome下按Ctrl+Shift+I，然后到Javascript控制台里，粘贴文中的代码，于是，你就可以用键盘的光标键移动并浏览整个世界了。

2）这是个全屏版的：[http://ares.aylett.co.uk/xkcd/](https://ares.aylett.co.uk/xkcd/)

3）如果你要下载所有的图，你可以使用这个[Python脚本](http://lebbeo.us/static/get-xkcd-1110.py)来完成（[转自这篇文章](http://lebbeo.us/2012/09/19/not-bbq-fetching-component-images-of-xkcd-comic-1110/)）

4）还有人把它搞成了像Google Map一样的东西。 你可以访问下面的链接：

> - <http://xkcd-map.rent-a-geek.de/>
> - <http://xkcdmap.webege.com/>
> 
> 5）看看Hacker News的讨论贴吧，什么都有了（<http://news.ycombinator.com/item?id=4542367>）

当然，对于这个图最强的一个站点如下，解释了所有和这个图有关信息，包括图中的各种文字和图案的意思。

<http://www.explainxkcd.com/wiki/index.php?title=1110:_Click_and_Drag>

看到这个图后，我陷入了深深地沉思，我在想。是什么样的动力能让人干出这样的事来？兴趣，还是为了好玩。还就是为了证明他能干一些让人拍案叫绝的东西？**这可能就是一种Geek精神吧。就是为了能做出让世人冿冿乐道的东西**。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
