# 关于C++构造函数的FAQ

下面是一些关于C++构造函数的FAQ。你能回答得出来吗？你可以点链接查看答案，不过是英文版的。他们来自于[*C++ FAQ Lite*](http://www.parashift.com/c++-faq-lite/index.html "C++ FAQ Lite")。当然，也有中文版的，只可惜中文版的太老了，只更新到了2001年。在[*C++ FAQ Lite*](http://www.parashift.com/c++-faq-lite/index.html "C++ FAQ Lite")上还有很多关于其它部分的FAQ，大家可以去看看。

[[1] 构造函数是用来干什么的？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.1 "[1] What's the deal with constructors?")

[[2] List x; 和 List x();有什么不同?](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.2 "[2] Is there any difference between List x; and List x();?")

[[3] 是否一个类的构造函数可以调用另一个构造函数来初始化自己？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.3 "[3] Can one constructor of a class call another constructor of the same class to initialize the this object?")

[[4] 是否Fred类的默认的函数函数就一定是Fred::Fred()？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.4 "[4] Is the default constructor for Fred always Fred::Fred()?")

[[5] 如果要创建一个Fred 对像数组，什么样的构数函数会被调用?](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.5 "[5] Which constructor gets called when I create an array of Fred objects?")

[[6] 构造函数初始化成员变量时，用 “初始化列表” 还是 “赋值”？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.6 "[6] Should my constructors use \"initialization lists\" or \"assignment\"?")

[[7] 在构造函数中用this 指针是否有问题？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.7 "[7] Should you use the this pointer in the constructor?")

[[8]什么是“名字构造函数”（Named Constructor Idiom）？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.8 "[8] What is the \"Named Constructor Idiom\"?")

[[9] “值返回”意味着额外的拷贝吗？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.9 "[9] Does return-by-value mean extra copies and extra overhead?")

[[10] 为什么我们不能在构造函数初始化列表中初始化一个 static 成员变量？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.10 "[10] Why can't I initialize my static member data in my constructor's initialization list?")

[[11] 为什么一个有 static 成员变量的类会有链接错误？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.11 "[11] Why are classes with static data members getting linker errors?")

[[12] 什么是“static initialization order fiasco”？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.12 "[12] What's the \"static initialization order fiasco\"?")

[[13] 我该如果避免 “static initialization order fiasco”?](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.13 "[13] How do I prevent the \"static initialization order fiasco\"?")

[[14] 为什么 construct-on-first-use 什么静态变量而不是指针？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.14 "[14] Why doesn't the construct-on-first-use idiom use a static object instead of a static pointer?")

[[15] 怎么才能避免静态成员中的“static initialization order fiasco” ？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.15 "[15] How do I prevent the \"static initialization order fiasco\" for my static data members?")

[[16] 我是否要为内建类型的“static initialization order fiasco”而担心？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.16 "[16] Do I need to worry about the \"static initialization order fiasco\" for variables of built-in/intrinsic types?")

[[17] 如果构造函数出错了怎么办？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.17 "[17] How can I handle a constructor that fails?")

[[18] 什么是“命名参数惯用法”（Named Parameter Idiom）？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.18 "[18] What is the \"Named Parameter Idiom\"?")

[[19] 为什么我通过Foo x(Bar())声明一个Foo 对象会得到一个错误？](http://www.parashift.com/c++-faq-lite/ctors.html#faq-10.19 "[19] Why am I getting an error after declaring a Foo object via Foo x(Bar())?")

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
