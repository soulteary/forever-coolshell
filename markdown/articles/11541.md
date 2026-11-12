# 面向GC的Java编程

**（感谢网友 [@Hesey小纯纯](http://weibo.com/tbmujian) 投稿  [博客](http://blog.hesey.net/) |　[原文链接](http://blog.hesey.net/2014/05/gc-oriented-java-programming.html)）**

Java程序员在编码过程中通常不需要考虑内存问题，JVM经过高度优化的GC机制大部分情况下都能够很好地处理堆(Heap)的清理问题。以至于许多Java程序员认为，我只需要关心何时创建对象，而回收对象，就交给GC来做吧！甚至有人说，如果在编程过程中频繁考虑内存问题，是一种退化，这些事情应该交给编译器，交给虚拟机来解决。

这话其实也没有太大问题，的确，大部分场景下关心内存、GC的问题，显得有点“杞人忧天”了，高老爷说过：

过早优化是万恶之源。

但另一方面，**什么才是“过早优化”？**

If we could do things right for the first time, why not?

事实上**JVM的内存模型**( [JMM](http://www.cs.umd.edu/~pugh/java/memoryModel/) )理应是Java程序员的基础知识，处理过几次JVM线上内存问题之后就会很明显感受到，很多系统问题，都是内存问题。

对JVM内存结构感兴趣的同学可以看下 [浅析Java虚拟机结构与机制](http://blog.hesey.net/2011/04/introduction-to-java-virtual-machine.html) 这篇文章，本文就不再赘述了，本文也并不关注具体的GC算法，相关的文章汗牛充栋，随时可查。

另外，不要指望GC优化的这些技巧，可以对应用性能有成倍的提高，特别是对I/O密集型的应用，或是实际落在YoungGC上的优化，可能效果只是帮你减少那么一点YoungGC的频率。

但我认为，**优秀程序员的价值，不在于其所掌握的几招屠龙之术，而是在细节中见真著**，就像前面说的，**如果我们可以一次把事情做对，并且做好，在允许的范围内尽可能追求卓越，为什么不去做呢？**

目录


- [一、GC分代的基本假设](#%E4%B8%80%E3%80%81GC%E5%88%86%E4%BB%A3%E7%9A%84%E5%9F%BA%E6%9C%AC%E5%81%87%E8%AE%BE "一、GC分代的基本假设")
- [二、对象分配的优化](#%E4%BA%8C%E3%80%81%E5%AF%B9%E8%B1%A1%E5%88%86%E9%85%8D%E7%9A%84%E4%BC%98%E5%8C%96 "二、对象分配的优化")
- [三、不可变对象的好处](#%E4%B8%89%E3%80%81%E4%B8%8D%E5%8F%AF%E5%8F%98%E5%AF%B9%E8%B1%A1%E7%9A%84%E5%A5%BD%E5%A4%84 "三、不可变对象的好处")
- [四、引用置为null的传说](#%E5%9B%9B%E3%80%81%E5%BC%95%E7%94%A8%E7%BD%AE%E4%B8%BAnull%E7%9A%84%E4%BC%A0%E8%AF%B4 "四、引用置为null的传说")
- [五、手动档的GC](#%E4%BA%94%E3%80%81%E6%89%8B%E5%8A%A8%E6%A1%A3%E7%9A%84GC "五、手动档的GC")
- [六、指定容器初始化大小](#%E5%85%AD%E3%80%81%E6%8C%87%E5%AE%9A%E5%AE%B9%E5%99%A8%E5%88%9D%E5%A7%8B%E5%8C%96%E5%A4%A7%E5%B0%8F "六、指定容器初始化大小")
- [七、对象池](#%E4%B8%83%E3%80%81%E5%AF%B9%E8%B1%A1%E6%B1%A0 "七、对象池")
- [八、对象作用域](#%E5%85%AB%E3%80%81%E5%AF%B9%E8%B1%A1%E4%BD%9C%E7%94%A8%E5%9F%9F "八、对象作用域")
- [九、各类引用](#%E4%B9%9D%E3%80%81%E5%90%84%E7%B1%BB%E5%BC%95%E7%94%A8 "九、各类引用")
#### 一、GC分代的基本假设

大部分GC算法，都将堆内存做分代(Generation)处理，但是为什么要分代呢，又为什么不叫内存分区、分段，而要用面向时间、年龄的“代”来表示不同的内存区域？

GC分代的**基本假设**是：

**绝大部分对象的生命周期都非常短暂，存活时间短。**

而这些短命的对象，恰恰是GC算法需要首先关注的。所以在大部分的GC中，YoungGC（也称作MinorGC）占了绝大部分，对于负载不高的应用，可能跑了数个月都不会发生FullGC。

基于这个前提，在编码过程中，我们应该**尽可能地缩短对象的生命周期**。在过去，分配对象是一个比较重的操作，所以有些程序员会尽可能地减少new对象的次数，尝试减小堆的分配开销，减少内存碎片。

但是，短命对象的创建在JVM中比我们想象的性能更好，所以，不要吝啬new关键字，大胆地去new吧。

当然前提是不做无谓的创建，对象创建的速率越高，那么GC也会越快被触发。

结论：

- 分配小对象的开销分享小，不要吝啬去创建。
- GC最喜欢这种小而短命的对象。
- 让对象的生命周期尽可能短，例如在方法体内创建，使其能尽快地在YoungGC中被回收，不会晋升(romote)到年老代(Old Generation)。

#### 二、对象分配的优化

基于大部分对象都是小而短命，并且不存在多线程的数据竞争。这些小对象的分配，会优先在线程私有的 **TLAB** 中分配，TLAB中创建的对象，不存在锁甚至是CAS的开销。

TLAB占用的空间在Eden Generation。

当对象比较大，TLAB的空间不足以放下，而JVM又认为当前线程占用的TLAB剩余空间还足够时，就会直接在Eden Generation上分配，此时是存在并发竞争的，所以会有CAS的开销，但也还好。

当对象大到Eden Generation放不下时，JVM只能尝试去Old Generation分配，这种情况需要尽可能避免，因为一旦在Old Generation分配，这个对象就只能被Old Generation的GC或是FullGC回收了。

#### 三、不可变对象的好处

GC算法在扫描存活对象时通常需要从ROOT节点开始，扫描所有存活对象的引用，构建出对象图。

不可变对象对GC的优化，主要体现在Old Generation中。

可以想象一下，如果存在Old Generation的对象引用了Young Generation的对象，那么在每次YoungGC的过程中，就必须考虑到这种情况。

Hotspot JVM为了提高YoungGC的性能，避免每次YoungGC都扫描Old Generation中的对象引用，采用了 **卡表(Card Table)** 的方式。

简单来说，当Old Generation中的对象发生对Young Generation中的对象产生新的引用关系或释放引用时，都会在卡表中响应的标记上标记为脏(dirty)，而YoungGC时，只需要扫描这些dirty的项就可以了。

可变对象对其它对象的引用关系可能会频繁变化，并且有可能在运行过程中持有越来越多的引用，特别是容器。这些都会导致对应的卡表项被频繁标记为dirty。

而不可变对象的引用关系非常稳定，在扫描卡表时就不会扫到它们对应的项了。

注意，这里的不可变对象，不是指仅仅自身引用不可变的final对象，而是真正的**Immutable Objects**。

#### 四、引用置为null的传说

早期的很多Java资料中都会提到在方法体中将一个变量置为null能够优化GC的性能，类似下面的代码：

```java
List<String> list = new ArrayList<String>();
// some code
list = null; // help GC

```

事实上这种做法对GC的帮助微乎其微，有时候反而会导致代码混乱。

我记得几年前 @rednaxelafx 在HLL VM小组中详细论述过这个问题，原帖我没找到，结论基本就是：

- 在一个非常大的方法体内，对一个较大的对象，将其引用置为null，某种程度上可以帮助GC。
- 大部分情况下，这种行为都没有任何好处。

所以，还是早点放弃这种“优化”方式吧。

GC比我们想象的更聪明。

#### 五、手动档的GC

在很多Java资料上都有下面两个奇技淫巧：

- 通过**Thread.yield()**让出CPU资源给其它线程。
- 通过**System.gc()**触发GC。

事实上JVM从不保证这两件事，而System.gc()在JVM启动参数中如果允许显式GC，则会**触发FullGC**，对于响应敏感的应用来说，几乎等同于自杀。

So，让我们牢记两点：

- Never use Thread.yield()。
- Never use System.gc()。除非你真的需要回收Native Memory。

第二点有个Native Memory的例外，如果你在以下场景：

- 使用了NIO或者NIO框架（Mina/Netty）
- 使用了DirectByteBuffer分配字节缓冲区
- 使用了MappedByteBuffer做内存映射

由于**Native Memory只能通过FullGC（或是CMS GC）回收**，所以除非你非常清楚这时真的有必要，否则不要轻易调用System.gc()，且行且珍惜。

另外为了防止某些框架中的System.gc调用（例如NIO框架、Java RMI），建议在启动参数中加上-XX:+DisableExplicitGC来禁用显式GC。

这个参数有个巨大的坑，如果你禁用了System.gc()，那么上面的3种场景下的内存就无法回收，可能造成OOM，如果你使用了CMS GC，那么可以用这个参数替代：-XX:+ExplicitGCInvokesConcurrent。

关于System.gc()，可以参考 @bluedavy 的几篇文章：

- [CMS GC会不会回收Direct ByteBuffer的内存](http://hellojava.info/?p=56)
- [说说在Java启动参数上我犯的错](http://hellojava.info/?p=323)
- [java.lang.OutOfMemoryError:Map failed](http://hellojava.info/?p=319)

#### 六、指定容器初始化大小

Java容器的一个特点就是可以动态扩展，所以通常我们都不会去考虑初始大小的设置，不够了反正会自动扩容呗。

但是扩容不意味着没有代价，甚至是很高的代价。

例如一些基于数组的数据结构，例如StringBuilder、StringBuffer、ArrayList、HashMap等等，在扩容的时候都需要做ArrayCopy，对于不断增长的结构来说，经过若干次扩容，会存在大量无用的老数组，而回收这些数组的压力，全都会加在GC身上。

这些容器的构造函数中通常都有一个可以指定大小的参数，如果对于某些大小可以预估的容器，建议加上这个参数。

可是因为容器的扩容并不是等到容器满了才扩容，而是有一定的比例，例如HashMap的扩容阈值和负载因子(loadFactor)相关。

Google Guava框架对于容器的初始容量提供了非常便捷的工具方法，例如：

[code lang=”java”]Lists.newArrayListWithCapacity(initialArraySize);

Lists.newArrayListWithExpectedSize(estimatedSize);

Sets.newHashSetWithExpectedSize(expectedSize);

Maps.newHashMapWithExpectedSize(expectedSize);  

[/code]

这样我们只要传入预估的大小即可，容量的计算就交给Guava来做吧。

**反例**：如果采用默认无参构造函数，创建一个ArrayList，不断增加元素直到OOM，那么在此过程中会导致：

- 多次数组扩容，重新分配更大空间的数组
- 多次数组拷贝
- 内存碎片

#### 七、对象池

为了减少对象分配开销，提高性能，可能有人会采取对象池的方式来缓存对象集合，作为复用的手段。

但是对象池中的对象由于在运行期长期存活，大部分会晋升到Old Generation，因此无法通过YoungGC回收。

并且通常……没有什么效果。

对于对象本身：

- 如果对象很小，那么分配的开销本来就小，对象池只会增加代码复杂度。
- 如果对象比较大，那么晋升到Old Generation后，对GC的压力就更大了。

从线程安全的角度考虑，通常池都是会被并发访问的，那么你就需要处理好同步的问题，这又是一个大坑，并且**同步带来的开销，未必比你重新创建一个对象小**。

对于对象池，唯一合适的场景就是**当池中的每个对象的创建开销很大**时，缓存复用才有意义，例如每次new都会创建一个连接，或是依赖一次RPC。

比如说：

- 线程池
- 数据库连接池
- TCP连接池

即使你真的需要实现一个对象池，也请使用成熟的开源框架，例如Apache Commons Pool。

另外，使用JDK的ThreadPoolExecutor作为线程池，不要重复造轮子，除非当你看过AQS的源码后认为你可以写得比Doug Lea更好。

#### 八、对象作用域

尽可能缩小对象的作用域，即生命周期。

- 如果可以在方法内声明的局部变量，就不要声明为实例变量。
- 除非你的对象是单例的或不变的，否则尽可能少地声明static变量。

#### 九、各类引用

java.lang.ref.Reference有几个子类，用于处理和GC相关的引用。JVM的引用类型简单来说有几种：

- Strong Reference，最常见的引用
- Weak Reference，当没有指向它的强引用时会被GC回收
- Soft Reference，只当临近OOM时才会被GC回收
- Phantom Reference，主要用于识别对象被GC的时机，通常用于做一些清理工作

当你需要实现一个缓存时，可以考虑优先使用WeakHashMap，而不是HashMap，当然，更好的选择是使用框架，例如Guava Cache。

最后，再次提醒，以上的这些未必可以对代码有多少性能上的提升，但是熟悉这些方法，是为了帮助我们写出更卓越的代码，和GC更好地合作。

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
