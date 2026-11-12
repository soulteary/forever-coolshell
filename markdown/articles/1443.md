# C++的std::string的“读时也拷贝”技术！

C++的std::string的读时也拷贝技术！

嘿嘿，你没有看错，我也没有写错，是读时也拷贝技术。什么?我的错，你之前听说写过时才拷贝，嗯，不错的确有这门技术，英文是Copy On Write，简写就是COW,非常’牛’！那么我们就来看看这个’牛’技术的效果吧。

我们先编写一段程序  



```c
#include <string>
#include <iostream>
#include <sys/time.h>

static long getcurrenttick()
{
    long tick ;
    struct timeval time_val;
    gettimeofday(&time_val , NULL);
    tick = time_val.tv_sec * 1000 + time_val.tv_usec / 1000 ;
    return tick;
}


int main( )
{
    string the_base(1024 * 1024 * 10, 'x');
    long begin =  getcurrenttick();
    for( int i = 0 ;i< 100 ;++i ) {
       string the_copy = the_base ;
    }
    fprintf(stdout,"耗时[%d] \n",getcurrenttick() - begin );
}

```

嗯，一个非常大的字符串，有10M字节的x，并且执行了100此拷贝。编译执行它，非常快，在我的虚拟机甚至不要1个毫秒。

现在我们来对这个string加点料！

```
int main(void) {
    string the_base(1024 * 1024 * 10, 'x');
    long begin =  getcurrenttick();
    for (int i = 0; i < 100; i++) {
        string the_copy = the_base;
        the_copy[0] = 'y';
    }
    fprintf(stdout,"耗时[%d] \n",getcurrenttick() - begin );
}

```

现在我们再编译并执行这断程序，居然需要4~5秒！哇！非常美妙的写时才拷贝技术，性能和功能的完美统一。

我们再来看看另外一种情况！

```
string original = "hello";
char & ref = original[0];
string clone = original;
ref = 'y';

```

我们生成了一个string，并保留了它首字符的引用，然后复制这个string，修改string中的首字符。因为写操作只是直接的修改了内存中的指定位置，这个string就根本不能感知到有写发生，如果写时才拷贝是不成熟的，那么我们将同时会修改original和clone两个string。那岂不是灾难性的结果？幸好上述问题不会发生。clone的值肯定是没有被修改的。看来COW就是非常的牛！

以上都证明了我们的COW技术非常牛！

有太阳就有黑暗，这句说是不是有点耳熟？

```
int main(void) {
    string the_base(1024 * 1024 * 10, 'x');
    fprintf(stdout,"the_base's first char is [%c]\n",the_base[0] );
    long begin =  getcurrenttick();
    for (int i = 0; i < 100; i++) {
        string the_copy = the_base;
    }
    fprintf(stdout,"耗时[%d] \n",getcurrenttick() - begin );
}

```

啊，居然也是4~5秒！你可能在想，我只是做了一个读，没有写嘛，这到底是怎么回事？难道还有读时也拷贝的技术！。

不错，为了避免了你通过[]操作符获取string内部指针而直接修改字符串的内容，在你使用了the\_base[0]后，这个字符串的写时才拷贝技术就失效了。

C++标准的确就是这样的，C++标准认为，当你通过迭代器或[]获取到string的内部地址的时候，string并不知道你将是要读还是要写。这是它无法确定，为此，当你获取到内部引用后，为了避免不能捕获你的写操作，它在此时废止了写时才拷贝技术！

这样看来我们在使用COW的时候，一定要注意，如果你不需要对string的内部进行修改，那你就千万不要使用通过[]操作符和迭代器去获取字符串的内部地址引用，如果你一定要这么做，那么你就必须要付出代价。当然，string还提供了一些使迭代器和引用失效的方法。比如说push\_back，等， 你在使用[]之后再使用迭代器之后，引用就有可能失效了。那么你又回到了COW的世界！比如下面的一个例子

```
int main( )
{
    struct timeval time_val;
    string the_base(1024 * 1024 * 10, 'x');
    long begin = 0 ;
    fprintf(stdout,"the_base's first char is [%c]\n",the_base[0] );
    the_base.push_back('y');
    begin = getcurrenttick();
    for( int i = 0 ;i< 100 ;++i ) {
        string the_copy = the_base ;
    }
    fprintf(stdout,"耗时[%d] \n",getcurrenttick() - begin );
}

```

一切又恢复了正常！如果对[]返回引用进行了操作又会发生情况呢，有兴趣的朋友可以试试！结果非常令人惊讶。

另外：上述例子是在linux环境下编译的，使用STL是GNU的STL。windows上我用的是vs2003，但是非常明显vs2003一点都不支持COW。

这篇文章出自<http://ridiculousfish.com/blog/archives/2009/09/17/i-didnt-order-that-so-why-is-it-on-my-bill-episode-2/> 这里，我使用了它的例子。但是我重新自己组织了内容。

编写这篇文章的同时，我还参考了耗子的[《标准C＋＋类string的Copy-On-Write技术》](http://blog.csdn.net/haoel/archive/2004/06/23/24058.aspx)一文

**（转载本站文章请注明作者和出处 [酷 壳 - CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
