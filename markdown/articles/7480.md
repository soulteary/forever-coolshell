# Javascript 中的 var

MelonCard发布了一篇文章——”[how one missing var ruined our launch](http://blog.meloncard.com/post/12175941935/how-one-missing-var-ruined-our-launch)“（”少写了一个var毁了我的网站”），这篇文章是说MelonCard用Node.js做后台，因为出了一个小高峰——有50-100人注册，结果整个网站都不响应了，而且还出现了很多奇怪的问题。当他们调查到问题的要源的时候，他们发现下面的代码少写了一个var。

[javascript]app.all(‘/apps/:user\_id/status’, function(req, res, next) {  

// …  

initial = extractVariables(req.body);  

});[/javascript]

为什么inital少写一个var会引发这个问题呢？因为如果你不写var，这个局部的变量会被javascript当成全局变量，而这个变量又是一个函数，所以，当多用户并发的时候，这个本应该在不同用户下互不干扰的变量，成了各个用户共享的东西。试想，用户A的数据被用户B覆盖了，用户A和B的数据还没处理完，结果被新的C给搞乱了，程序的逻辑自然出现了问题。

在stackoverflow.com上有[这么一个贴子说明了“有var”和“无var”](http://stackoverflow.com/questions/1470488/difference-between-using-var-and-not-using-var-in-javascript)的差别：

```js
// These are both globals
var foo = 1;
bar = 2;

function test()
{
    var foo = 1; // Local
    bar = 2;     // Global

    // Execute an anonymous function
    (function()
    {
        var wibble = 1; // Local
        foo = 2; // Inherits from scope above (creating a closure)
        moo = 3; // Global
    }())
}
```

上面这个示例告诉我们，如果你不用var，那么这个js引擎会一层一层地向上找父作用域中的变量，如果找到了，就用，如果找不到了，就会帮你定义一个全局的变量。上面这个例子充分说明了这一点。所以，**如果你想在当前的作用域用声明变量，你一定要用var**。这对于一些乱写javascript代码的程序员要注意了。这里再给大家介绍一个工具——

**JSLint( <http://www.jslint.com/> )**，一个JS代码质量的分析工具，我们把上述stackoverflow的代码copy到JSLint这个在线工具中，我们可以看到下面的报告：

![](https://coolshell.cn/uploads/2012/05/jslint.jpg "jslint")

这个报告说明了源码中的那些变量的情况。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
