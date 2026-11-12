# 写HTML和CSS的新方法

**[Zen Coding](https://code.google.com/p/zen-coding/)** 一个用来简化编写 HTML，XML， XSL （或是其它一些诸如此类格式的编辑器）。其主要是用一种缩写方式的语法来书写大量重复和无味的HTML，很像CSS语法。下面是一个例子：
```
div#page>div.logo+ul#navigation>li*5>a
```
展开后会成为下面这个样子：
```html
<div id="page">
        <div></div>
        <ul id="navigation">
                <li><a href=""></a></li>
                <li><a href=""></a></li>
                <li><a href=""></a></li>
                <li><a href=""></a></li>
                <li><a href=""></a></li>
        </ul>
</div>

```
可以看出来，#代表ID，>代表下一层。

如果你写下：
```
select>option#item-$*3
```
那么将会得到：
```html


<select>
	<option id="item-1"></option>
	<option id="item-2"></option>
	<option id="item-3"></option>
</select>

```
看上去很不错吧。目前，其支持如下的编辑器：

 

- [AptanaHowToEn](https://code.google.com/p/zen-coding/wiki/AptanaHowToEn)
- **TextMate** (Mac). Available in two flavors: basic snippets (Zen HTML and Zen CSS) and full-featured plugin (ZenCoding for TextMate). Bundles > Zen Coding menu item
- **Coda** (Mac) — [external download](https://github.com/sergeche/tea-for-coda/downloads), via [TEA for Coda](http://onecrayon.com/tea/). Plug-ins > TEA for Coda > Zen Coding menu item
- **Espresso** (Mac) — [external download](https://github.com/sergeche/tea-for-espresso/downloads), via [TEA for Espresso](http://onecrayon.com/tea/). Zen Coding is bundled with Espresso by default, but you should upgrade ZC to latest version. Actions > HTML menu item
- **Komodo Edit/IDE** (crossplatform) — [external download](http://community.activestate.com/xpi/zen-coding). Tools > Zen Coding menu item
- **Notepad++** (Windows). Zen Coding menu item
- **PSPad** (Windows). Scripts > Zen Coding menu item
- **<textarea>** (browser-based). See [online demo](http://zen-coding.ru/textarea/).
- **editArea** (browser-based). See [online demo](http://zen-coding.ru/demo/).

还有下面这些第三方的插件：

- **Dreamweaver** (Windows, Mac)
- **Sublime Text** (Windows)
- **UltraEdit** (Windows)
- **TopStyle** (Windows)
- **GEdit** (crossplatform) — [Franck Marcia’s plugin](https://github.com/fmarcia/zen-coding-gedit), [Mike Crittenden’s plugin](https://github.com/mikecrittenden/zen-coding-gedit)
- **BBEdit/TextWrangler** (Mac) — [external download](http://www.angelwatt.com/coding/zen-coding_bbedit.php)
- **Visual Studio** (Windows) — [external download](http://zencoding.codeplex.com/)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
