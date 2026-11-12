# Waterfall 软件工程

《Royce, Winston (1970), [“Managing the Development of Large Software Systems”](http://www.cs.umd.edu/class/spring2003/cmsc838p/Process/waterfall.pdf), *Proceedings of IEEE WESCON* **26** (August): 1–9. 》，这篇文章向你说明了软件工程鼻祖“Waterfall”的工程模型，这是40年前的论文，其中的十张插图很有强大，抽出来，让我们来看看什么叫Waterfall软件工程。

首先，让我先看一下小的程序是怎么做的，呵呵，很简单，两步。

[![](https://coolshell.cn/uploads/2010/09/01.Small_.jpg "01.Small")](https://coolshell.cn/wp-content/uploads/2010/09/01.Small_.jpg)

接下来，就是我们最经典的Waterfall软件工程模型了，用户需求，软件需求，需求分析，设计，编码，测试，运维。

[![](https://coolshell.cn/uploads/2010/09/02.Large_.jpg "02.Large")](https://coolshell.cn/wp-content/uploads/2010/09/02.Large_.jpg)

为了保证每个步骤都能正确实施，于是，每个步骤之间需要有一定的交互，这是我们所希望的样子。

 [![](https://coolshell.cn/uploads/2010/09/03.Iteraction.jpg "03.Iteraction")](https://coolshell.cn/wp-content/uploads/2010/09/03.Iteraction.jpg)

然后，不幸的是，我们总是在测试的时候发现了设计甚至需求的问题，因此，不得不让我们返工。

[![](https://coolshell.cn/uploads/2010/09/04.Design.jpg "04.Design")](https://coolshell.cn/wp-content/uploads/2010/09/04.Design.jpg)

为了解决上面的“返工”问题，我们可以使用下面的几步来解决。

第一步，叫Preliminary Design，程序设计先行，确定在进入需求分析之前，我们的概要设计要完整。

[![](https://coolshell.cn/uploads/2010/09/05.01.Preliminary.Design.jpg "05.01.Preliminary.Design")](https://coolshell.cn/wp-content/uploads/2010/09/05.01.Preliminary.Design.jpg)

第二步，我们叫Document Design，书写设计文档，确认我们的设计是完整的。看到了吧，总共6个文档，1）软件需求，2）概要设计，3）接口设计，4）各种最终设计，5）测试设计/计划，6）测试结果。流程开始变得复杂了。

[![](https://coolshell.cn/uploads/2010/09/05.02.Documentation.jpg "05.02.Documentation")](https://coolshell.cn/wp-content/uploads/2010/09/05.02.Documentation.jpg)

第三步，我们叫“Do it Twice”，双保险，把文档了的东西试着预先走一遍，看看能否成为最终产品。

[![](https://coolshell.cn/uploads/2010/09/05.03.Double.Work_.jpg "05.03.Double.Work")](https://coolshell.cn/wp-content/uploads/2010/09/05.03.Double.Work_.jpg)

第四步，计划，控制和监控测试。哇，流程很乱了。

[![](https://coolshell.cn/uploads/2010/09/05.04.Test_.jpg "05.04.Test")](https://coolshell.cn/wp-content/uploads/2010/09/05.04.Test_.jpg)

第五步，用户介入，全程review各个环节。

[![](https://coolshell.cn/uploads/2010/09/05.05.Involve.Customer.jpg "05.05.Involve.Customer")](https://coolshell.cn/wp-content/uploads/2010/09/05.05.Involve.Customer.jpg) 

好了，问题解决了，让我们看看最终的“无比强大的”——Waterfall软件工程模型！

[![](https://coolshell.cn/uploads/2010/09/06.Summary.jpg "06.Summary")](https://coolshell.cn/wp-content/uploads/2010/09/06.Summary.jpg)

现在，当你在使用waterfall开发软件的时候，知道为什么痛苦了吧，40年前就已经如此了。

下面是《[Lone Star Ruby Conference 2010](http://confreaks.net/events/lsrc2010)》的一个演讲叫《[Real Software Engineer](http://confreaks.net/videos/282-lsrc2010-real-software-engineering)》，没有字幕，但我个人感觉英文很容易听懂，英文好的同学不妨看看。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
