# Kick Ass小游戏

还记得以前那个在IE的已打开的网页上的网址里输入一段javascript的代码后，你会发现这个页面里所有的图片元素都动了起来：（只能在IE浏览器里，Chrome和Firefox无效）

> javascript:R=0; x1=.1; y1=.05; x2=.25; y2=.24; x3=1.6; y3=.24; x4=300; y4=200; x5=300; y5=200; DI=document.images; DIL=DI.length; function A(){for(i=0; i-DIL; i++){DIS=DI[ i ].style; DIS.position=’absolute’; DIS.left=Math.sin(R\*x1+i\*x2+x3)\*x4+x5; DIS.top=Math.cos(R\*y1+i\*y2+y3)\*y4+y5}R++}setInterval(‘A()’,5); void(0);

很类似一个叫Erik Rothoff Andersson的人又搞了[Kick Ass的游戏](https://erkie.github.com/)，代码如下：（用了一个js文件，所以就显得没有那么复杂了，但只能在Chrome和Firefox下有用）

> javascript:var s=document.createElement(‘script’); s.type=’text/javascript’;document.body.appendChild(s); s.src=’http://erkie.github.com/asteroids.min.js’;void(0);

在已打开的网页上输入这段代码，你会发现网页的左上角上出现了一个三角形，然后，你可以开始使用

- “左右方向键控制方向”，
- “上方向键控制前进”，
- “空格射击”，
- “B键查看有什么东西可以被射击”，
- “Esc键退出”，

于是就出现好玩的东西了。

为了方便你试验，你可以点击上面的这个链接，

[**Kiss Ass**](javascript:var%20s%20=%20document.createElement('script');s.type='text/javascript';document.body.appendChild(s);s.src='http://erkie.github.com/asteroids.min.js';void(0);)

你可以把这个链接加入收藏夹，当你需要删除某些网页上的广告或是很让你不爽的东西时，打开这个网址，就可以开始了。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
