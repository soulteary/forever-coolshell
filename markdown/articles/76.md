# 怎样做一个 Program Manager

我个人认为，这是一篇不错的文章，虽然我不是Program Mananger，但是我几乎在做着和这个职位很相似的工作。在这里，我把这篇文章推荐给所有的程序员，我相信，这篇文章会让你明白，只有技术是远远不够的，因为没有Program Manager这个角色，程序员们只不过一些手中拿着利器却不知所措的散兵游勇。我希望我的导读和原文能给所有的程序带来启示。

**原文在这里：**  

“How to be a program manager”  

<http://www.joelonsoftware.com/items/2009/03/09.html>

[![09meeting-thumbnail](https://coolshell.cn/uploads/2009/03/09meeting-thumbnail-300x198.jpg "09meeting-thumbnail")](https://coolshell.cn/wp-content/uploads/2009/03/09meeting-thumbnail.jpg)这篇文章的作者叫Joel Spolsky，在Microsoft做过Program Manager，这篇文章非常值得一读。下面是我给大家做的一个导读：

首先，他讲了两个人，一个是负责WYSIWYG 字处理的天才级的Program Manager——Charles Simonyi，第二个是上世纪80年代的负责Mac OS上的Excel项目的程序员Jabe Blumenthal，他发现了程序员和市场人员的代沟，Marketing的人很难通过把MBA-Speaking翻译成实际的Feature，并且，有太多的和编码不相关的工作，比如说，和用户交谈，运行usability测试，Reivew竞争者的产品，并且得冥思苦想怎么能让事情变得更简单，而我们的程序员通常来说即不具备这样的时间，也不具备这样的能力。于是，Jabe开始了他的Program Manager的生涯。

**工作范围**

作者在第二节里说了一个PM主要负责哪些事务：

1. Design UIs （用户界面的设计）
2. Write functional specs （书写功能规格说明书）
3. Coordinate teams （团队协调）
4. Serve as the customer advocate, and （从用户角度思考问题）
5. Wear Banana Republic chinos （Banana Republic是一个服装品牌，意思是作者在调侃PM需要衣冠楚楚，而不像程序员们只有T恤或牛仔裤）

接下来，作者讲述了他第一份Program Manager工作的经历，非常有意思，那是一个关于Excel 用户定制化的项目（耗子注：应该是在Excel中加入VBScript的项目吧，就是所谓的宏）。

第一个阶段

- 首先，作者找了很多很多的用户谈论了这个什么是最有用最合理的实现，这是一个非常巨大的工作，花费了非常多的精力和时间。
- 然后，作者找到了Visual Basic团队询问了是否可能给Excel提供一个编译器和代码编辑器，以便实现“宏”。
- 接着，作者查看了一下Apple上面的AppleScript这种宏，取了取经。
- 最后，作者同 Word, Access, Project, 和Mail团队们讨论了很多很多。

作者说，这个阶段的工作让他满是伤痕，他甚至害怕听到手机铃响。

第二个阶段

- 确定大方向。他开始写下Visual Baisc应该怎么样在Excel里面工作的文档。并提供了一些简单的宏的样子。这应该是high-level的Functional Spec。
- 当大的方向确定后，他开始了一些更为细节的功能规格说明的书写。这就是所谓的Functional Specification. (耗子注：这份文档应该只是说明从用户的角度上来看这个产品长成什么样，而不是实现)
- 虽然FS并不需要说明怎么去实现，但这份文档应该是需要非常详细地说明整个Excel和VBScript怎么相互交互的，这是其中最重要的部分。
- 当作者把FS的一个初始化版本发给开发团队（Ben Waldman）时，开发团队非常快地实现出了一个原型，并提供了面向对象的相关接口。但可惜的是，那并不是Program Manger所想要的。
- 作者描述了一个细节如果帮助开发团队解决技术难点的例子。那是关于把一个Excel中的一个cell的值取出来的例子。当时，developer团队认为这是一个难点，因为这个值可能是任意类型的。而VB中却需要先声明变量的类型。后来，作者找到了VB的开发团队，了解到了Variants 和IDispatch可以做到这个。

我们可以看到，FS在这样反复地和developer 团队推敲，甚至去帮助程序员解决技术难题，之后最终才能确定下来。一旦FS确定后，program manger需要做两件事：

1. 负责解释相关的问题。
2. 组织并形成相关的design。

也就是说，除了对FS解释外，需还需要把What needs to do 变成 How to do的设计文档。另外，Program Manager可能会有下面的工作：

- 测试人员会对FS有很多很多疑问，因为他们需要知道怎么样去测试这些FS中所包含的东西。
- 和文档团队商讨如何写一个好的教程或是一个参考文档。
- 和localization 团队制定localization 的策略。
- 和市场人员说明VBA的优势和功能。

我们可以看到，作者有太多，太多的会议和太多的与人沟通的事务，真是一个不简单的工作啊。

**冲突管理**

后面，作者着重讲了“Conflicts”冲突，这可能是所有的团队都会有的问题。而我们的Program Manager因为要和那么多的人沟通交流，所以，必然会需要有一种超人的能力去管理与人的发生的观点上的冲突。作者，在这里说了和程序员发生的很多争论，因为Program Manager是从用户的角度出发，而我们程序员总是从技术和实现的角度出发，不同的角度必然会引发冲突。作者举了一个例子，他说，用户们喜欢一个“心灵感应”的界面和一个30英寸的显示器，而我们的程序员喜欢的只是用Python搞的命令行接口。呵呵。另外，作者引用了一个Excel中的“pivot tables ”所引发的一个历时最长的争议作为案例。

最后，作者讨论了，争论是一个很好的事，就好像法院里的原告和被告都有自己的辩护律师一样，这有助于人们逼近事物的真相。对于软件开发也一样，良好的争论其实是对产品有好处的。我们应该在争论中关注事。

当在讨论到和程序相处的过程，作者说到了和程序员相外并不是一件很容易的事，因为你并不编码而也没有技术能力，通常会受到程序员的冷眼。所以在和程序沟通的过程中需要保证两件事：1）确信自己的正确的。2）让程序员尊敬自己。而对于第二点，如何让程序员尊敬自己，作者发表了自己的见解：1）demonstrate intelligence（展示自己的才华），2）open-mindedness（心胸宽阔），3）fairness（公平，正直）。千万不要搞办公室政治，或是开私密的经理会，等等。不然的话，你必然受到排挤。

**推荐读物**

最后作者给大家推荐了一些很不错的读物：

- [Making Things Happen](https://www.amazon.com/gp/product/0596517718?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0596517718 "blocked::http://www.amazon.com/gp/product/0596517718?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0596517718") （经理一般都在干什么？）
- [Don’t Make Me Think](https://www.amazon.com/gp/product/0321344758?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0321344758 "blocked::http://www.amazon.com/gp/product/0321344758?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0321344758")  （如果你要写FS或UI设计，你应该看看这本书）
- [User Interface Design for Programmers](https://www.amazon.com/gp/product/1893115941?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=1893115941 "blocked::http://www.amazon.com/gp/product/1893115941?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=1893115941"). （作者自己的书，关于UI设计）
- [How to Win Friends & Influence People](https://www.amazon.com/gp/product/0671027034?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0671027034 "blocked::http://www.amazon.com/gp/product/0671027034?ie=UTF8&tag=joelonsoftware&linkCode=as2&camp=1789&creative=390957&creativeASIN=0671027034") （在人际关系方面，需要看看这本书）

（完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
