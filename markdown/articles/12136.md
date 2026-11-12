# Google Inbox如何跨平台重用代码？

原文链接《[How Google Inbox shares 70% of its code across Android, iOS, and the Web](http://arstechnica.com/information-technology/2014/11/how-google-inbox-shares-70-of-its-code-across-android-ios-and-the-web)》

[![inbox2-640x264](https://coolshell.cn/uploads/2014/11/inbox2-640x264-300x123.jpg)](https://coolshell.cn/wp-content/uploads/2014/11/inbox2-640x264.jpg)

开发一个移动应用在当下并不是一件容易的事情。如果想要获得最多的用户，你的应用通常需要覆盖 iOS, Android, 和 Web 三大平台。这就意味着同一个应用需要开发三个版本，使用 Objective-C 或者 Swift 开发 iOS 版本，使用 Java 开发 Android 版本，使用 JavaScript/CSS/HTML5 开发 Web 版本。工作量增大的同时也意味着有更多的 bug 需要修复。

这个问题也是 Google 在开发 Google Inbox 时致力要解决的。在最近发布的这款应用中，Google 使用了一些工具实现了70%的代码跨平台复用。

Google Inbox 覆盖 iOS, Android, Web 三个平台，它们使用的是同一个后台代码逻辑，只是前端的用户体验和平台相关特性的实现有所不同。Google 自主开发了一套辅助工具将 Android 版本的 Java 代码逻辑编译为 Objective-C (针对 iOS 平台) 和 JavaScript (针对 Web 浏览器)。 Java 到 JavaScript 的编译由 Google Web Toolkit SDK 完成，Java 到 Objective-C 的编译则由 J2ObjC （<j2objc.org>）来完成。

J2ObjC 是一个开源项目，由 Google 在2013年发布。Google Sheets (Google Docs 中的电子表格部分) 也使用了 J2ObjC，而 Google Inbox 则是目前使用 J2Objc 最多的 Google 项目。

Google Inbox 复用的代码逻辑包括：对话 (conversations)，提醒 (reminders)，联系人 (contacts)。还有网络相关功能和离线同步。这些代码逻辑的复用节省了大量的时间和成本。

在产品设计时，Google 将这些可复用功能划分为抽象的逻辑概念，比如：提醒的逻辑放在 “reminder.java” 中，可以被 Android UI 调用。对 iOS 版本而言，J2ObjC 将 “reminder.java” 编译成 Objective-C 代码，再由 iOS UI 调用。

Google 没有跨平台编译 UI 部分的代码，因为不同平台的UI特性各有不同，盲目统一会导致非常糟糕的用户体验。代码复用只是针对可以共享的后台逻辑，前端的UI实现是完全原生 (native) 的。这与 Xamarin (一个基于 Microsoft C# 的跨平台移动开发工具) 提出的概念类似。

跨平台代码复用通常会带来一些性能上的问题。Garrick Toubassi，Engineering Director 和 Google Inbox 项目组成员，对此表示： “性能上的影响如果有的话，也可以说是微不足道的。我们做过大量的性能测试。因为没有加入额外的中间层来处理跨平台兼容性，所有代码最后都是平台原生代码。J2ObjC 编译生成的目标代码和 Java 源代码拥有大致相同的对象数量和对象图谱复杂度 (object graph complexity) ”。

Google 使用的整套方法解决了跨平台移动开发中的一个很重要的问题，同时也推进了安卓先行 (Android-first) 的移动开发策略。

更多 Google Inbox 文章请猛戳 [Gmail 官方博客](http://gmailblog.blogspot.com.au/2014/11/going-under-hood-of-inbox.html)。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
