# 理解Javascript的闭包

【感谢 Neo 投递本文 – 微博帐号：[\_锟\_](http://weibo.com/gandalfthegrey "_锟_") 】

**前言：还是一篇入门文章。**Javascript中有几个非常重要的语言特性——对象、原型继承、闭包。其中闭包对于那些使用传统静态语言C/C++的程序员来说是一个新的语言特性。本文将以例子入手来介绍Javascript闭包的语言特性，并结合一点ECMAScript语言规范来使读者可以更深入的理解闭包。

注：**本文是入门文章，例子素材整理于网络**，如果你是高手，欢迎针对文章提出技术性建议和意见。本文讨论的是Javascript，不想做语言对比，如果您对Javascript天生不适，请自行绕道。****

目录


- [什么是闭包](#%E4%BB%80%E4%B9%88%E6%98%AF%E9%97%AD%E5%8C%85 "什么是闭包")
- [ECMAScript闭包模型](#ECMAScript%E9%97%AD%E5%8C%85%E6%A8%A1%E5%9E%8B "ECMAScript闭包模型")
- [闭包的样列](#%E9%97%AD%E5%8C%85%E7%9A%84%E6%A0%B7%E5%88%97 "闭包的样列")
- [闭包的应用](#%E9%97%AD%E5%8C%85%E7%9A%84%E5%BA%94%E7%94%A8 "闭包的应用")
#### **什么是闭包**

闭包是什么?闭包是Closure，这是静态语言所不具有的一个新特性。但是闭包也不是什么复杂到不可理解的东西，简而言之，闭包就是：

- **闭包就是函数的局部变量集合，只是这些局部变量在函数返回后会继续存在。**
- **闭包就是就是函数的“堆栈”在函数返回后并不释放，我们也可以理解为这些函数堆栈并不在栈上分配而是在堆上分配**
- **当在一个函数内定义另外一个函数就会产生闭包**

上面的第二定义是第一个补充说明，抽取第一个定义的主谓宾——闭包是**函数的‘局部变量’集合**。只是这个局部变量是可以在函数返回后被访问。（这个不是官方定义，但是这个定义应该更有利于你理解闭包）

做为局部变量都可以被函数内的代码访问，这个和静态语言是没有差别。闭包的差别在于局部变变量可以在函数执行结束后仍然被函数外的代码访问。这意味着函数必须返回一个指向闭包的“引用”，或将这个”引用”赋值给某个外部变量，才能保证闭包中局部变量被外部代码访问。当然包含这个引用的实体应该是一个对象，因为在Javascript中除了基本类型剩下的就都是对象了。可惜的是，ECMAScript并没有提供相关的成员和方法来访问闭包中的局部变量。但是在ECMAScript中，函数对象中定义的**内部函数(inner function)**是可以直接访问外部函数的局部变量，通过这种机制，我们就可以以如下的方式完成对闭包的访问了。

[javascript]  

function greeting(name) {  

var text = ‘Hello ‘ + name; // local variable  

// 每次调用时，产生闭包，并返回内部函数对象给调用者  

return function() { alert(text); }  

}  

var sayHello=greeting("Closure");  

sayHello() // 通过闭包访问到了局部变量text  

[/javascript]

上述代码的执行结果是：Hello Closure，因为sayHello()函数在greeting函数执行完毕后，仍然可以访问到了定义在其之内的局部变量text。

好了，这个就是传说中闭包的效果，闭包在Javascript中有多种应用场景和模式，比如Singleton，Power Constructor等这些Javascript模式都离不开对闭包的使用。

#### **ECMAScript闭包模型**

ECMAScript到底是如何实现闭包的呢？想深入了解的亲们可以获取[ECMAScript 规范](http://www.ecma-international.org/publications/files/ECMA-ST/Ecma-262.pdf)进行研究，我这里也只做一个简单的讲解，内容也是来自于网络。

在ECMAscript的脚本的函数运行时，每个函数关联都有一个执行上下文场景(Execution Context) ，这个执行上下文场景中包含三个部分

- 文法环境（The LexicalEnvironment）
- 变量环境（The VariableEnvironment）
- this绑定

其中第三点this绑定与闭包无关，不在本文中讨论。文法环境中用于解析函数执行过程使用到的变量标识符。我们可以将文法环境想象成一个对象，该对象包含了两个重要组件，环境记录(Enviroment Recode)，和外部引用(指针)。环境记录包含包含了函数内部声明的局部变量和参数变量，外部引用指向了外部函数对象的上下文执行场景。全局的上下文场景中此引用值为NULL。这样的数据结构就构成了一个单向的链表，每个引用都指向外层的上下文场景。

例如上面我们例子的闭包模型应该是这样，sayHello函数在最下层，上层是函数greeting，最外层是全局场景。如下图：  

![](https://coolshell.cn/uploads/2012/03/closure.png)  

因此当sayHello被调用的时候，sayHello会通过上下文场景找到局部变量text的值，因此在屏幕的对话框中显示出”Hello Closure”  

变量环境(The VariableEnvironment)和文法环境的作用基本相似，具体的区别请参看ECMAScript的规范文档。

#### **闭包的样列**

前面的我大致了解了Javascript闭包是什么，闭包在Javascript是怎么实现的。下面我们通过针对一些例子来帮助大家更加深入的理解闭包，下面共有5个样例，例子来自于[JavaScript Closures For Dummies(](http://blog.morrisjohns.com/javascript_closures_for_dummies.html)[镜像](http://web.archive.org/web/20080209105120/http://blog.morrisjohns.com/javascript_closures_for_dummies)[)](http://blog.morrisjohns.com/javascript_closures_for_dummies.html)。  

**例子1:闭包中局部变量是引用而非拷贝**

[javascript]  

function say667() {  

// Local variable that ends up within closure  

var num = 666;  

var sayAlert = function() { alert(num); }  

num++;  

return sayAlert;  

}

var sayAlert = say667();  

sayAlert()  

[/javascript]

因此执行结果应该弹出的667而非666。

**例子2：多个函数绑定同一个闭包，因为他们定义在同一个函数内。**

[javascript]  

function setupSomeGlobals() {  

// Local variable that ends up within closure  

var num = 666;  

// Store some references to functions as global variables  

gAlertNumber = function() { alert(num); }  

gIncreaseNumber = function() { num++; }  

gSetNumber = function(x) { num = x; }  

}  

setupSomeGlobals(); // 为三个全局变量赋值  

gAlertNumber(); //666  

gIncreaseNumber();  

gAlertNumber(); // 667  

gSetNumber(12);//  

gAlertNumber();//12  

[/javascript]

**例子3：当在一个循环中赋值函数时，这些函数将绑定同样的闭包**

[javascript]  

function buildList(list) {  

var result = [];  

for (var i = 0; i < list.length; i++) {  

var item = ‘item’ + list[i];  

result.push( function() {alert(item + ‘ ‘ + list[i])} );  

}  

return result;  

}

function testList() {  

var fnlist = buildList([1,2,3]);  

// using j only to help prevent confusion – could use i  

for (var j = 0; j < fnlist.length; j++) {  

fnlist[j]();  

}  

}  

[/javascript]

testList的执行结果是弹出item3 undefined窗口三次，因为这三个函数绑定了同一个闭包，而且item的值为最后计算的结果，但是当i跳出循环时i值为4，所以list[4]的结果为undefined.

**例子4：外部函数所有局部变量都在闭包内，即使这个变量声明在内部函数定义之后。**

[javascript]  

function sayAlice() {  

var sayAlert = function() { alert(alice); }  

// Local variable that ends up within closure  

var alice = ‘Hello Alice’;  

return sayAlert;  

}  

var helloAlice=sayAlice();  

helloAlice();  

[/javascript]

执行结果是弹出”Hello Alice”的窗口。即使局部变量声明在函数sayAlert之后，局部变量仍然可以被访问到。

**例子5：每次函数调用的时候创建一个新的闭包**

[javascript]  

function newClosure(someNum, someRef) {  

// Local variables that end up within closure  

var num = someNum;  

var anArray = [1,2,3];  

var ref = someRef;  

return function(x) {  

num += x;  

anArray.push(num);  

alert(‘num: ‘ + num +  

‘\nanArray ‘ + anArray.toString() +  

‘\nref.someVar ‘ + ref.someVar);  

}  

}  

closure1=newClosure(40,{someVar:’closure 1′});  

closure2=newClosure(1000,{someVar:’closure 2′});

closure1(5); // num:45 anArray[1,2,3,45] ref:’someVar closure1′  

closure2(-10);// num:990 anArray[1,2,3,990] ref:’someVar closure2’  

[/javascript]

#### **闭包的应用**

**Singleton 单件：**

[javascript]  

var singleton = function () {  

var privateVariable;  

function privateFunction(x) {  

…privateVariable…  

}

return {  

firstMethod: function (a, b) {  

…privateVariable…  

},  

secondMethod: function (c) {  

…privateFunction()…  

}  

};  

}();  

[/javascript]

这个单件通过闭包来实现。通过闭包完成了私有的成员和方法的封装。匿名主函数返回一个对象。对象包含了两个方法，方法1可以方法私有变量，方法2访问内部私有函数。需要注意的地方是匿名主函数结束的地方的'()’，如果没有这个'()’就不能产生单件。因为匿名函数只能返回了唯一的对象，而且不能被其他地方调用。这个就是利用闭包产生单件的方法。

## **参考：**

[JavaScript Closures For Dummies(](http://blog.morrisjohns.com/javascript_closures_for_dummies.html)[镜像](http://web.archive.org/web/20080209105120/http://blog.morrisjohns.com/javascript_closures_for_dummies)[)](http://blog.morrisjohns.com/javascript_closures_for_dummies.html) 可惜都被墙了。  

[Advance Javascript](http://yuiblog.com/blog/2006/11/27/video-crockford-advjs/) （Douglas Crockford 大神的视频，一定要看啊）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
