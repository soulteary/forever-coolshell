# 9个最常见IE的Bug及其fix

![9个最常见IE的Bug及其fix](http://nettuts.s3.amazonaws.com/494_ie/images/200x200.jpg)

Internet Explorer – Web程序员的毒药。在IE上开发时间中有超过60%的时间是花在和IE的bug进行搏斗，让你的开发生产率严重下降。下面是一个教程，告诉你9个IE上最常见的BUG以及如何解决它们。

目录


- [1. 居中布局](#1_%E5%B1%85%E4%B8%AD%E5%B8%83%E5%B1%80 "1. 居中布局")
- [2. 楼梯式的效果](#2_%E6%A5%BC%E6%A2%AF%E5%BC%8F%E7%9A%84%E6%95%88%E6%9E%9C "2. 楼梯式的效果")
- [3. float元件的两倍空白](#3_float%E5%85%83%E4%BB%B6%E7%9A%84%E4%B8%A4%E5%80%8D%E7%A9%BA%E7%99%BD "3. float元件的两倍空白")
- [4. 无法设置微型高度](#4_%E6%97%A0%E6%B3%95%E8%AE%BE%E7%BD%AE%E5%BE%AE%E5%9E%8B%E9%AB%98%E5%BA%A6 "4. 无法设置微型高度")
- [5. 跨出边界](#5_%E8%B7%A8%E5%87%BA%E8%BE%B9%E7%95%8C "5. 跨出边界")
- [6. Fixing the Broken Box Model](#6_Fixing_the_Broken_Box_Model "6. Fixing the Broken Box Model")
- [7. 设置min-height和min-width](#7_%E8%AE%BE%E7%BD%AEmin-height%E5%92%8Cmin-width "7. 设置min-height和min-width")
- [8. Float 布局错误行为 Misbehaving](#8_Float_%E5%B8%83%E5%B1%80%E9%94%99%E8%AF%AF%E8%A1%8C%E4%B8%BA_Misbehaving "8. Float 布局错误行为 Misbehaving")
- [9. 在list项目门的空行](#9_%E5%9C%A8list%E9%A1%B9%E7%9B%AE%E9%97%A8%E7%9A%84%E7%A9%BA%E8%A1%8C "9. 在list项目门的空行")
- [结论](#%E7%BB%93%E8%AE%BA "结论")
#### 1. 居中布局

创建一个CSS定义把一个元素放到中间的位置，可能是每一个Web开发人员都会做的事情。最简单的做法是为你的元素增加一个*margin: auto;* ，然而 IE 6.0 会出现很多奇怪的行为。让我们来看一个例子。

```css
#container{
	border: solid 1px #000;
	background: #777;
	width: 400px;
	height: 160px;
	margin: 30px 0 0 30px;
}

#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 100px;
	margin: 30px auto;

}

```

下面是我们所期望的输出：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/1-1.png)

但IE却给我们这样的输出：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/1-2.png)

这应该是IE 6对margin的 *auto* 并没有正确的设置。但幸运的是，这是很容易被修正的。

**解决方法**

最简单的方法是在父元件中使用 *text-align: center* 属性，而在元件中使用 *text-align: left* 。

```css
#container{
	border: solid 1px #000;
	background: #777;
	width: 400px;
	height: 160px;
	margin: 30px 0 0 30px;
	text-align: center;
}

#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 100px;
	margin: 30px 0;
    	text-align: left;

}

```
#### 2. 楼梯式的效果

几乎所有的Web开发者都会使用list来创建导航条。下面是你可能会用到的代码：

```html
    <ul>
        <li><a href="#"></a></li>
        <li><a href="#"></a></li>
        <li><a href="#"></a></li>
    </ul>

```
```css
ul {
    list-style: none;
}

ul li a {
   	display: block;
   	width: 130px;
	height: 30px;
   	text-align: center;
   	color: #fff;
   	float: left;
	background: #95CFEF;
	border: solid 1px #36F;
	margin: 30px 5px;
}

```

一个符合标准的浏览器会是下面这样：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/2-1.png)

但IE却是这样的：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/2-2.png)

下面是两个解决方法

**解决方法一**

设置li元件的float属性。

```css
ul li {
	float: left;
}

```

**解决方法二**

设置 *display: inline* 属性。

```css
ul li {
	display: inline
}

```
#### 3. float元件的两倍空白

请看下面的代码：

```css
#element{
	background: #95CFEF;
	width: 300px;
	height: 100px;
	float: left;
	margin: 30px 0 0 30px;
	border: solid 1px #36F;
}

```

期望的结果是：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/3-1.png)

IE的结果是：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/3-2.png)

**解决方案**

和上面那个BUG的解决方案一样，设置 *display: inline* 属性可以解决问题。

```css
#element{
	background: #95CFEF;
	width: 300px;
	height: 100px;
	float: left;
	margin: 30px 0 0 30px;
	border: solid 1px #36F;
   	display: inline;
}

```
#### 4. 无法设置微型高度

我们发现在IE中使用 *height: XXpx* 这样的属性无法设置比较小的高度。下面是个例子（注意高度是2px）：

```css
#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 2px;
	margin: 30px 0;
}

```

期望结果： 2px的元件加1px的边框.

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/4-1.png)

IE的结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/4-2.png)

**解决方案一**

这个BUG的产生原因很简单，IE不允许元件的高度小于字体的高度，所以，下面的fix是设置上字体大小。

```css
#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 2px;
	margin: 30px 0;
    	font-size: 0;
}

```

**解决方案二**

但是最佳的解决方法是使用 *overflow: hidden* 。

```css
#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 2px;
	margin: 30px 0;
    	overflow: hidden
}

```
#### 5. 跨出边界

这个BUG是很难看的。当父元件中使用了 *overflow* 的 *auto* 属性，并且在其里放入相关元件。你会看来里面的元件会跨出来。下面是一个示例：

```html
<div id="element"><div id="anotherelement"></div></div>

```
```css
#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 150px;
	margin: 30px 0;
	overflow: auto;
}

#anotherelement{
	background: #555;
	width: 150px;
	height: 175px;
	position: relative;
	margin: 30px;
}

```

期望的结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/5-1.png)

IE的结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/5-2.png)

**解决方法**

设置 position: relative;属性

```css
#element{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 300px;
	height: 150px;
	margin: 30px 0;
	overflow: auto;
    	position: relative;
}

```
#### 6. Fixing the Broken Box Model

Internet Explorer曲解了“盒子模子”可能是最不可原谅的事情了。IE 6 这个半标准的浏览器回避了这个事情，但这个问题还是会因为IE运行在“怪异模式”下出现。

两个Div元件。一个是有fix的，一个是没有的。而他们不同的高和宽加上padding的总合却是不一样的。下图的上方是被修正的，下方则没有。

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/6.png)

**解决方法**

我相信这个事情即不需要解释也不需要演示，这应该是大多数人都明白的。下面是一个很相当怪异的解决方案

```css
#element{
	width: 400px;
    	height: 150px;
	padding: 50px;
}

```

上面的定义也就是说：

```css
#element {
    width: 400px;
    height: 150px;
   \height: 250px;
   \width: 500px
}

```

是的，你要原来的长和宽上加上了padding。但这个fix只会作用于IE了的“怪异模式”，所以你不需要担心在IE6的正常模式下会有问题。

#### 7. 设置min-height和min-width

IE忽略了min-height。

**解决方法一**

这个fix由 [Dustin Diaz](http://www.dustindiaz.com/min-height-fast-hack/)提供。其利用了 *!important* 下面是代码片段：

```css
#element {
  min-height:150px;
  height:auto !important;
  height:150px;
}

```

**解决方法二**

```css
#element {
    min-height: 150px;
    height: 150px;
}

html>body #element {
	height: auto;
}

```
#### 8. Float 布局错误行为 Misbehaving

使用无table的布局最重要的就是使用CSS的float元件。在很多情况下，IE6处理起来好像在摸索阶段，有些时候，你会发现很多奇怪的行为。比如在其中有一些文本的时候。

来看一下下面这个示例：

```html
<div id="container">
	<div id="element">http://net.tutsplus.com/</div>
	<div id="anotherelement"></div>
</div>

```
```css
#element, #anotherelement{
	background: #95CFEF;
	border: solid 1px #36F;
	width: 100px;
	height: 150px;
	margin: 30px;
	padding: 10px;
	float: left;
}

#container{
	background: #C2DFEF;
	border: solid 1px #36F;
	width: 365px;
	margin: 30px;
	padding: 5px;
	overflow: auto;
}

```

期望结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/8-1.png)

IE的结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/8-2.png)

你可以看到其中的不同了

**解决方法**

要解决这个问题没有什么好的方法。只有一个方法，那就是使用 *overflow: hidden* 。

```css
#element{
	background: #C2DFEF;
	border: solid 1px #36F;
	width: 365px;
	margin: 30px;
	padding: 5px;
	overflow: hidden;
}

```
#### 9. 在list项目门的空行

先看下面的例子

```html
<ul>
 <li><a href="#">Link 1</a></li>
 <li><a href="#">Link 2</a></li>
 <li><a href="#">Link 3</a></li>
</ul>

```
```css
ul {
	margin:0;
	padding:0;
	list-style:none;
}

li a {
	background: #95CFEF;
	display: block;
}

```

期望结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/9-1.png)

IE的结果：

![Tutorial Image](http://nettuts.s3.amazonaws.com/494_ie/images/9-2.png)

Fortunately, there are a plethora of fixes you could try.

**解决方法一**

定义height来解决

```css
li a {
	background: #95CFEF;
	display: block;
    	height: 200px;
}

```

**解决方法二**

```css
li a {
	background: #95CFEF;
	float: left;
    	clear: left;
}

```

**解决方法三**

为 *li* 加上*display: inline*。

```css
li {
	display: inline;
}

```
#### 结论

调界面是一件很难的事，调一个CSS的HTML界面是一件更难的事，在IE下调一个CSS的HTML界面是难上加难的事。

文章：[来源](http://net.tutsplus.com/tutorials/html-css-techniques/9-most-common-ie-bugs-and-how-to-fix-them/)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
