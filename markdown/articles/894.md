# 将vim变得简单:如何在vim中得到你最喜爱的IDE特性

原文出处:[这里](http://arstechnica.com/open-source/guides/2009/05/vim-made-easy-how-to-get-your-favorite-ide-features-in-vim.ars)

**摘要：**  

开源的vim文本编辑器提供许多灵活而强大的功能，但是vim自身是很难被配置使用的，在本教材中，我们将向你显示通过几个简单的方式使得你的vim具有集成开发环境IDE的行为

vim是很多程序员和系统管理员最爱的文本编辑器，虽然他提供了很多优秀而灵活的功能，但是对于新手来说他依然是难于上手的。从传统集成开发环境转到vim的开发人员通常会开在发方式的转变中发现迷失了自己。

我经常收到来自于读者的邮件，他们希望能找到一种方式使得vim变得对开发者更友好。一个常见的抱怨是vim并不是自身就带有IDE的特性，并且如何来通过配置能得到等价IDE功能也不是很清晰。而揭开vim真正神奇的秘密就是利用强大的vim插件系统和对vim自身功能的改善和增强的第三方脚本。在阅你读本文之前，我已经整理好了一个vim的有用tips和插件列表，这些列表中的内容将会使那些用惯IDE功能的人们在vim上感到宾至如归的感觉。

虽然vim主要是设计给基于字符方式的文本编辑器，并且它有可能是这类编辑器中最高效的工具，但是现在在vim上也存在一些更适合新手使用的基于图形的外壳。不像运行在终端窗口上的vim，你可以尝试使用一下gvim,一个基于GUI的vim版本。gvim拥有可配置的的菜单和工具条，因此可以通过鼠标直接访问到vim的编程上的最本质的特性。gvim可以让你使用操作系统自带的文件对话框，并允许你通过鼠标点击拖拉编辑面板的能力。gvim有windows和linux的版本，等价的Mac OS X的版本是MacVim，MacVim提供了Mac机的本地Cocoa用户接口，包括菜单集成的功能。  

[![vimtxt_gvim_ars](https://coolshell.cn/uploads/2009/05/vimtxt_gvim_ars.jpg "vimtxt_gvim_ars")](https://coolshell.cn/?attachment_id=896)

我听到来自vim用户最经常被抱怨的功能是vim的编辑区列表非常麻烦，并且没有一种简单的方式可以明了的看到什么文件是打开的。在vim上有几个插件可以解决这个问题，并提供了一个额外的编辑区列表用于方便在打开文件中切换。我最喜欢的一个插件是[MiniBufExplorer](http://www.vim.org/scripts/script.php?script_id=159)，它将列表显示在窗口的头上。当MiniBufExplorer被激活时，你可以通过tab键来在列表的这些项中循环，然后通过回车键或双击鼠标来选择在编辑区显示和你要处理的文件。  

[![vimtxt_vim_minibufexplorer_ars](https://coolshell.cn/uploads/2009/05/vimtxt_vim_minibufexplorer_ars.jpg "vimtxt_vim_minibufexplorer_ars")](https://coolshell.cn/?attachment_id=898)

许多的IDE工具都有用于显示你程序项目结构和允许你通过鼠标在特定的类和方法间跳转的代码导航区。你可以通过使用流行的[Tag List 插件](http://vim-taglist.sourceforge.net/installation.html)来得到这个特性。这个插件需要[Exuberant Ctags](http://ctags.sourceforge.net/)实用工具，这个工具用于分析你的代码。TagList可以通过命令:Tlist来激活，并将你的类和方法显示在激活的区域，当你打开其他的文件或切换到其他打开文件时，新的类或方法会被加到代码导航区。在gvim中你可以通过单击方法名跳到对应方法定义。如果要使用键盘，那么通过光标键上下移光标到你希望的方法处，单击回车即可达到目标。

[![vimtxt_vim_taglist_ars](https://coolshell.cn/uploads/2009/05/vimtxt_vim_taglist_ars.jpg "vimtxt_vim_taglist_ars")](https://coolshell.cn/?attachment_id=895)

自动文本完成(**译者注**：就是eclipse，visual studio中常见的输入前几个字符后面的内容通过列表显示的功能)是另外一种在IDE工具中常用特性，并且很多用户都希望在vim中有这些特性。这个特性已经在vim7中通过[Omnicompletion system](http://vim.wikia.com/wiki/Omni_completion)被引入进来。它是可编程，这就意味着你可以通过定制，使的这个功能能在各种个样的编程语言中使用，在vim中甚至存在对动态语言python或ruby生效的自动文本完成功能。现在，自动文本完成的配置已经变成了vim包中的一个部分，所以现在你可以什么都不做就能让这个功能生效。要调出自动完成菜单(列表)，你需要敲下ctrl+x和ctrl+o键，接着你可以用ctrl+n和ctrl+p在可能完成列表中进行上下选择，当你移动到一个选项，vim将为你在另外一个Scratch区域显示带方法说明和属性的上下文帮助信息。  

[![vimtxt_vim_completion_ars](https://coolshell.cn/uploads/2009/05/vimtxt_vim_completion_ars.jpg "vimtxt_vim_completion_ars")](https://coolshell.cn/?attachment_id=897)

你可以多种方式来改善你的vim体验，[vim维基vim wiki](http://vim.wikia.com/wiki/Main_Page)和[脚本库script repository](http://www.vim.org/scripts/index.php)为你提供了可用于增强功能的第三方增强扩展集合。这些插件实现sinppet system，outlining tools，项目管理工具，和大量的其他的特性。同时还有大量的脚本实现了对某些特定编程语言和框架的增强。例如有一个[非常流行的脚本](http://www.vim.org/scripts/script.php?script_id=1567)，这个脚本将会改善你Ruby的语法高亮，并且为你Ruby on Rail的部署提供了非常方便的导航特性

同时也有一些[面向新手的脚本集合](http://cream.sourceforge.net/)，这个集合使得vim的行为变得更像一个带有简单菜单和快捷键的传统的文本编辑器。如果你对vim那些神秘键盘命名感到不舒服的话，你可以选择这个作为你使用vim的开始。

vim的多样性使得它满足不同的用户使用。对于那些没有时间，能力，和爱好去通过自己去建立一个完美vim配置的人来说，无数的第三方脚本和插件为你提供了一种简单的方式，通过这种方式你可以付出很少的努力就能得到你想要的功能和特性。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
