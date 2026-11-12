# 并发框架Disruptor译文

**（感谢同事[方腾飞](http://ifeve.com)投递本文）**

![](https://coolshell.cn/uploads/2013/02/Disruptor-300x144.png)Martin Fowler在自己网站上写了一篇[LMAX架构](http://ifeve.com/lmax)的文章，在文章中他介绍了LMAX是一种新型零售金融交易平台，它能够以很低的延迟产生大量交易。这个系统是建立在JVM平台上，其核心是一个业务逻辑处理器，它能够在一个线程里每秒处理6百万订单。业务逻辑处理器完全是运行在内存中，使用事件源驱动方式。业务逻辑处理器的核心是Disruptor。

Disruptor它是一个开源的并发框架，并获得[2011 Duke’s](http://www.java.net/dukeschoice)程序框架创新奖，能够在无锁的情况下实现网络的Queue并发操作。本文是[Disruptor官网](https://code.google.com/p/disruptor/wiki/BlogsAndArticles)中发布的文章的译文（[现在被移到了GitHub](https://lmax-exchange.github.com/disruptor/)）。

#### **剖析Disruptor:为什么会这么快**

- [剖析Disruptor:为什么会这么快？(一)锁的缺点](http://ifeve.com/locks-are-bad/)

- [剖析Disruptor:为什么会这么快？(二)神奇的缓存行填充](http://ifeve.com/disruptor-cacheline-padding/ "剖析Disruptor:为什么会这么快？（二）神奇的缓存行填充")

- [剖析Disruptor:为什么会这么快？(三)伪共享](http://ifeve.com/falsesharing/ "伪共享(False Sharing)")

- [剖析Disruptor:为什么会这么快？(四)揭秘内存屏障](http://ifeve.com/disruptor-memory-barrier/ "剖析Disruptor:为什么会这么快？(四)揭秘内存屏障")

#### Disruptor如何工作和使用

- [如何使用Disruptor（一）Ringbuffer的特别之处](http://ifeve.com/dissecting-disruptor-whats-so-special/ "剖析Disruptor:为什么会这么快？（一）Ringbuffer的特别之处")

- [如何使用Disruptor（二）如何从Ringbuffer读取](http://ifeve.com/dissecting_the_disruptor_how_doi_read_from_the_ring_buffer/ "如何使用Disruptor（二）如何从Ringbuffer读取")

- [如何使用Disruptor（三）写入Ringbuffer](http://ifeve.com/disruptor-writing-ringbuffer/ "如何使用 Disruptor（三）写入 Ringbuffer")

- [Disruptor(无锁并发框架)-发布](http://ifeve.com/the-disruptor-lock-free-publishing/ "Disruptor(无锁并发框架)-发布")

- [LMAX Disruptor——一个高性能、低延迟且简单的框架](http://ifeve.com/disruptor-dsl/ "LMAX Disruptor——一个高性能、低延迟且简单的框架")

- [Disruptor Wizard已死，Disruptor Wizard永存！](http://ifeve.com/disruptor-wizard/ "Disruptor Wizard已死，Disruptor Wizard永存！")

- [Disruptor 2.0更新摘要](http://ifeve.com/disruptor-2-change/ "Disruptor 2.0更新摘要")

- [线程间共享数据不需要竞争](http://ifeve.com/sharing-data-among-threads-without-contention/ "线程间共享数据无需竞争")

#### Disruptor的应用

- [LMAX的架构](http://ifeve.com/lmax/ "LMAX架构")

- [通过Axon和Disruptor处理1M tps](http://ifeve.com/axon/ "通过Axon和Disruptor处理1M tps")

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
