# 使用Google API做统计图

Google提供了一个的统计图的API。你可以通过构造一个URL链接来获得Google提供的统计图方案。

比如：如果我们使用如下链接：

```html
<img src="http://chart.apis.google.com/chart?cht=p3&chd=t:60,40&chs=250x100&chl=酷壳|Cocre" alt="" />

```

我们就可能通过如下的HTML代码显示一个60:40的饼图：  

![](https://chart.apis.google.com/chart?cht=p3&chd=t:60,40&chs=250x100&chl=酷壳|Cocre)

Google的这个API支持的统计图风格相当的多。

比如：

```html
<img src="http://chart.apis.google.com/chart?chs=200x125&cht=ls&chco=0077CC&chd=t:27,25,60,31,25,39,25,31,26,28,80,28,27,31,27,29,26,35,70,25" alt="Sparkline chart in blue" />

```

![Sparkline chart in blue](https://chart.apis.google.com/chart?chs=200x125&cht=ls&chco=0077CC&chd=t:27,25,60,31,25,39,25,31,26,28,80,28,27,31,27,29,26,35,70,25)

还甚至支持有世界地图式的统计图：

```html
<img src="http://chart.apis.google.com/chart?cht=t&chs=440x220&chd=t:0,100,50,32,60,40,43,12,14,54,98,17,70,76,18,29&chco=FFFFFF,FF0000,FFFF00,00FF00&chld=DZEGMGAOBWNGCFKECGCVSNDJTZGHMZZM&chtm=africa&chf=bg,s,EAF7FE" alt="Map of Africa" />

```

![Map of Africa](https://chart.apis.google.com/chart?cht=t&chs=440x220&chd=t:0,100,50,32,60,40,43,12,14,54,98,17,70,76,18,29&chco=FFFFFF,FF0000,FFFF00,00FF00&chld=DZEGMGAOBWNGCFKECGCVSNDJTZGHMZZM&chtm=africa&chf=bg,s,EAF7FE)

更多的内容请到[http://code.google.com/apis/chart/](https://code.google.com/apis/chart/) 上查看吧。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
