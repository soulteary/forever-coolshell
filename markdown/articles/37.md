# 【引文】如何用Python往Google Spreadsheet上写数据

现代企业里，数据决定着方向，人们都想随时看到各种报表。很多项目可能都需要dashboard一类的工作，把分散的数据变成一些能随时查看实时数据的图表，这个工作有两个环节：

1. 把数据汇集起来，放入CSV或者数据库
2. 一个服务器端的程序能够读到这写数据，根据需要生成在线的图表 （离线的也可以，那样每次人们想看这些图的时候都会来麻烦你，如果你在度假，他们会想敲开你的电脑）

第一步可以通过定期跑些脚本完成，但是第二步有时候就不太容易了，如果你希望你的图表能够让所有人方便随时查看，最方便的给出一个URL能让人随时访问，Google的在线文档可以提供一个简单的解决方案。

但是，如何将数据自动弄到在Google spreadsheet 上呢？手动的copy/paste是一个方法，但是很费人工，最简单的方法就是写个脚本把这个流程自动化。如何将数据写进Spreadsheet (在线表单)呢？请参考下文：

[Write to a Google Spreadsheet from a Python script](http://www.mattcutts.com/blog/write-google-spreadsheet-from-python/ "Permanent link to Write to a Google Spreadsheet from a Python script")

注：这是个搜索方面比较大拿的Googler的博客。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
