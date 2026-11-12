# Unix Pipes 管道原稿

![Douglas McIlroy](https://coolshell.cn/uploads/2009/08/Douglas-McIlroy.jpg "Douglas McIlroy")

40年前，Unix操作系统横空出世，Unix不仅仅带来了一个操作系统，还创造C语言，Socket，开源，黑客等等文化，这些文化影响着整个计算机世界的文明，直到今天。

如果说Unix是计算机文明中最伟大的发明，那么，Unix下的Pipe管道就是跟随Unix所带来的另一个伟大的发明。管道的出现，解决的就是让不同功能的程序可以互相连通通讯，从而可以让软件开发，程序开发更加的“高内聚，低耦合”，从而可以让程序“Do one thing, Do it well”，从而可以让程序“Keep it Simple Stupid”等等，这一哲学引影了一代又一代的软件架构，直到今天的云计算。

管道的发名者叫，[**Malcolm Douglas McIlroy**](https://en.wikipedia.org/wiki/Douglas_McIlroy)，他也是Unix的创建者，是Unix文化的缔造者之一。他归纳的Unix哲学如下：

> 程序应该只关注一个目标，并尽可能把它做好。让程序能够互相协同工作。应该让程序处理文本数据流，因为这是一个通用的接口。

下面是管道在1964年10月11日，出现的第一个打印稿，下面是扫描件。

![Unix Pipes](https://coolshell.cn/uploads/2009/08/pipe.png "Unix Pipes")

全文如下：

```
                        - 10 -
            Summary--what's most important.
    To put my strongest concerns into a nutshell:
1. We should have some ways of connecting programs like
garden hose--screw in another segment when it becomes when
it becomes necessary to massage data in another way.
This is the way of IO also.
2. Our loader should be able to do link-loading and
controlled establishment.
3. Our library filing scheme should allow for rather
general indexing, responsibility, generations, data path
switching.
4. It should be possible to get private system components
(all routines are system components) for buggering around with.

                                                M. D. McIlroy
                                                October 11, 1964

```

我就不翻译了，因为这段文字足够的简单，就像连接花园中浇花用的软管一样，相信你不但能够读懂它，还能从中收益。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
