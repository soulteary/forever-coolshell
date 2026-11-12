# 简单实用的Code Review工具

![](http://www.review-board.org/media/rbsite/images/logo.png?1238930581 "Code Review")Code Review中文应该译作“代码审查”或是“代码评审”，这是一个流程，当开发人员写好代码后，需要让别人来review一下他的代码，这是一种有效发现BUG的方法。由此，我们可以审查代码的风格、逻辑、思路……，找出问题，以及改进代码。因为这是代码刚刚出炉的时候，所以，这也是代码重构，代码调整，代码修改的最佳时候。所以，Code Review是编码实现中最最重要的一个环节。

长时间以来，Code Review需要有一些有效的工具来支持，这样我们就可以更容易，更有效率地来进行代码审查工作。下面是5个开源的代码审查工具，他们可以帮助你更容易地进行这项活动。

**1. [Review board](http://www.review-board.org/):**  

[Review board](http://www.review-board.org/) 是一个 基于web 的工具，是由 [django](https://www.djangoproject.com/) 和[python](https://www.python.org/)设计的。 [Review board](http://www.review-board.org/) 可以帮助我们追踪待决代码的改动，并可以让Code-Review更为容易和简练。尽管[Review board](http://www.review-board.org/) 最初被设计在[VMware](http://www.vmware.com/)项目中使用，但现在其足够地通用。当前，其支持这些代码版本管理软件： [SVN](http://subversion.tigris.org/), CVS, [Perforce](http://www.perforce.com/), [Git](http://git-scm.com/), [Bazaar](http://bazaar-vcs.org/), 和[Mercurial](http://www.selenic.com/mercurial/wiki/).

Yahoo 是[review-board](http://www.review-board.org/)的其中一个用户。

“[Review board](http://www.review-board.org/) 已经改变了代码评审的方式，其可以强迫高质量的代码标准和风格，并可以成为程序员编程的指导者。每一次，当你访问search.yahoo.com 时，其代码都是使用 [Review board](http://www.review-board.org/)工具Review过的。 We’re great fans of your work!” – Yahoo! Web Search

### [Detailed review requests](http://www.review-board.org/media/screenshots/2009/02/02/review-requests.png)

[![Powerful diff viewer](http://www.review-board.org/media/screenshots/2009/02/02/diffviewer_thumb.png)](http://www.review-board.org/media/screenshots/2009/02/02/diffviewer.png)

**2. [Codestriker](http://codestriker.sourceforge.net/):**  

[Codestriker](http://codestriker.sourceforge.net/) 也是一个基于Web的应用，其主要使用 GCI-Perl 脚本支持在线的代码审查。[Codestriker](http://codestriker.sourceforge.net/) 可以集成于CVS, [Subversion](http://subversion.tigris.org/), [ClearCase](http://www-01.ibm.com/software/awdtools/clearcase/), [Perforce](http://www.perforce.com/) 和Visual SourceSafe。并有一些插件可以提供支持其它的源码管理工具。

David Sitsky 是 [Codestriker](http://codestriker.sourceforge.net/) 的作者，并也是最活跃的开发人员之一。 Jason Remillard 是另一个活路的开发者，并给这个项目提供了最深远最有意义的贡献。大量的程序员贡献他们的代码给 [Codestriker](http://codestriker.sourceforge.net/) 项目，导致了这个项目空前的繁荣。

![http://codestriker.sourceforge.net/viewtopicdetail.png](http://codestriker.sourceforge.net/viewtopicdetail.png)

**3. [Groogle](http://groogle.sourceforge.net/):**  

[Groogle](http://groogle.sourceforge.net/) 是一个基于WEB的代码评审工具。 [Groogle](http://groogle.sourceforge.net/) 支持和 [Subversion](http://subversion.tigris.org/) 集成。它主要提供如下的功能：

- 各式各样语言的语法高亮。
- 支持整个版本树的比较。
- 支持当个文件不同版本的diff功能，并有一个图形的版本树。
- 邮件通知所有的Reivew的人当前的状态。
- 认证机制。

![Screenshot](http://sourceforge.net/dbimage.php?id=218190)

**4. [Rietveld](https://code.google.com/p/rietveld/):**  

[Rietveld](https://code.google.com/p/rietveld/) 由Guido van Rossum 开发（他是Python的创造者，现在是Google的员工），这个工具是基于Mondrian 工具，作者一开始是为了Google 开发的，并且，它在很多方面和[Review board](http://www.review-board.org/) 很像。它也是一个基于Web的应用，并在[Google App Engine](https://code.google.com/appengine/) 上。它使用了目前最流行的Web开发框架 [django](https://www.djangoproject.com/) 并支持 [Subversion](http://subversion.tigris.org/) 。当前，任何一个使用 Google Code 的项目都可以使用 [Rietveld](https://code.google.com/p/rietveld/) 并且使用 [python](https://www.python.org/) [Subversion](http://subversion.tigris.org/) 服务器。当然，它同样支持其它的Subversion服务器。

[![](http://info-database.csdn.net/Upload/2008-11-13/Reviewboard.jpg "下一张")](javascript:dyn.onClickNextTbn())

 

**5. [JCR](http://jcodereview.sourceforge.net/)**  

[JCR](http://jcodereview.sourceforge.net/) 或者叫做 JCodeReview 也是一个基于WEB界面的最初设计给Reivew Java 语言的一个工具。当然，现在，它可以被用于其它的非Java的代码。

[JCR](http://jcodereview.sourceforge.net/) 主要想协助：

- **审查者**。所有的代码更改都会被高亮，以及大多数语言的语法高亮。Code extracts 可以显示代码评审意见。如果你正在Review Java的代码，你可以点击代码中的类名来查看相关的类的声明。
- **项目所有者**。可以 轻松创建并配置需要Review的项目，并不需要集成任何的软件配置管理系统（SCM）。
- **流程信仰者**。 所有的评语都会被记录在数据库中，并且会有状态报告，以及各种各样的统计。
- **架构师和开发者**。 这个系统也可以让我们查看属于单个文件的评语，这样有利于我们重构代码。

[JCR](http://jcodereview.sourceforge.net/) 主要面对的是大型的项目，或是非常正式的代码评审，从这方面看来，他并不像上面的那些工具。

![Screenshot](http://sourceforge.net/projects/jcodereview/screenshots/242251)

**[Jupiter](https://code.google.com/p/jupiter-eclipse-plugin/)**：最后我们要提一下[Jupiter](https://code.google.com/p/jupiter-eclipse-plugin/)，这是另一个代码review的工具你可以去考虑使用的，它是一个Eclipse IDE 的插件。

文章：[来源](http://open-tube.com/easy-code-review-tools/)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
