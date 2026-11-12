# Facebook 的系统架构

**来源**：[http://www.quora.com/What-is-Facebooks-architecture](https://www.quora.com/What-is-Facebooks-architecture "What is Facebook's Architecture?") （由[Micha?l Figuière](https://www.quora.com/Micha%C3%ABl-Figui%C3%A8re)回答）

根据我现有的阅读和谈话，我所理解的今天Facebook的架构如下：

- Web 前端是由 PHP 写的。Facebook 的 [HipHop]( http://developers.facebook.com/blog/post/358) [1] 会把PHP转成 C++ 并用 g++编译，这样就可以为模板和Web逻贺业务层提供高的性能。

- 业务逻辑以Service的形式存在，其使用[Thrift](http://thrift.apache.org/) [2]。这些Service根据需求的不同由PHP，C++或Java实现（也可以用到了其它的一些语言……）

- 用Java写的Services没有用到任何一个企业级的应用服务器，但用到了Facebook自己的定制的应用服务器。看上去好像是重新发明轮子，但是这些Services只被暴露给Thrift使用（绝大所数是这样），Tomcat太重量级了，即使是Jetty也可能太过了点，其附加值对Facebook所需要的没有意义。

- 持久化由MySQL, [Memcached](http://memcached.org/) [3], Facebook 的 [Cassandra](http://cassandra.apache.org/) [4], Hadoop 的 [HBase](http://hbase.apache.org/) [5] 完成。Memcached 使用了MySQL的内存Cache。Facebook 工程师承认他们的Cassandra 使用正在减少，因为他们更喜欢HBase，因为它的更简单的一致性模型，以到其MapReduce能力。

- 离线处理使用Hadoop 和 Hive。

- 日志，点击，feeds数据使用[Scribe](https://github.com/facebook/scribe) [6]，把其聚合并存在 HDFS，其使用[Scribe-HDFS](https://hadoopblog.blogspot.com/2009/06/hdfs-scribe-integration.html) [7]，因而允许使用MapReduce进行扩展分析。

- [BigPipe](https://www.facebook.com/notes/facebook-engineering/bigpipe-pipelining-web-pages-for-high-performance/389414033919) [8] 是他们的定制技术，用来加速页面显示。

- [Varnish Cache](http://www.varnish-cache.org/) [9]用作HTTP代理。他们用这个的原因是[高速和有效率](http://www.varnish-software.com/customers/facebook)。 [10].

- 用来搞定用户[上传的十亿张照片的存储](https://www.facebook.com/note.php?note_id=76191543919)，其由Haystack处理，Facebook自己开发了一个Ad-Hoc存储方案，其主要做了一些低层优化和“仅追加”写技术 [11].

- Facebook Messages 使用了自己的架构，其明显地构建在了一个动态集群的基础架构上。业务逻辑和持久化被封装在一个所谓的’Cell’。每个‘Cell’都处理一部分用户，新的‘Cell’可以因为访问热度被添加[12]。 持久化归档使用HBase [13]。

- Facebook Messages 的搜索引擎由存储在HBase中的一个倒置索引的构建。 [14]

- Facebook 搜索引擎实现细节据我所知目前是未知状态。

- Typeahead 搜索使用了一个定制的存储和检索逻辑。 [15]

- Chat 基于一个Epoll 服务器，这个服务器由Erlang 开发，由Thrift存取 [16]

关于那些供给给上述组件的资源，下面是一些信息和数量，但是有一些是未知的：

- Facebook估计有超过60,000 台服务器[16]。他们最新的数据中心在俄勒冈州的Prineville，其基于完全自定设计的硬件[17] 那是最近才公开的 [Open Compute 项目](http://opencompute.org)[18]。

- 300 TB 的数据存在 Memcached 中处理 [19]

- 他们的Hadoop 和 Hive 集群由3000 服务器组成，每台服务器有8个核，32GB的内存，12TB的硬盘，全部有2万4千个CPU的核，96TB内存和36PB的硬盘。 [20]

- 每天有1000亿的点击量，500亿张照片， 3 万亿个对象被 Cache，每天130TB的日志（[2010年7月的数据](https://www.facebook.com/note.php?note_id=409881258919)） [21]

**参考引用**

[1] *HipHop for PHP*: [http://developers.facebook.com/blog/post/358](https://developers.facebook.com/blog/post/358)  
[2] *Thrift*: <http://thrift.apache.org/>  
[3] *Memcached*: <http://memcached.org/>  
[4] *Cassandra*: <http://cassandra.apache.org/>  
[5] *HBase*: <http://hbase.apache.org/>  
[6] *Scribe*: <https://github.com/facebook/scribe>  
[7] *Scribe-HDFS*: [http://hadoopblog.blogspot.com/2009/06/hdfs-scribe-integration.html](https://hadoopblog.blogspot.com/2009/06/hdfs-scribe-integration.html)  
[8] *BigPipe*: [http://www.facebook.com/notes/facebook-engineering/bigpipe-pipelining-web-pages-for-high-performance/389414033919](https://www.facebook.com/notes/facebook-engineering/bigpipe-pipelining-web-pages-for-high-performance/389414033919)  
[9] *Varnish Cache*: <http://www.varnish-cache.org/>  
[10] *Facebook goes for Varnish*: <http://www.varnish-software.com/customers/facebook>  
[11] *Needle in a haystack*: efficient storage of billions of photos: [http://www.facebook.com/note.php?note\_id=76191543919](https://www.facebook.com/note.php?note_id=76191543919)  
[12] *Scaling the Messages Application Back End*: [http://www.facebook.com/note.php?note\_id=10150148835363920](https://www.facebook.com/note.php?note_id=10150148835363920)  
[13] *The Underlying Technology of Messages*: <https://www.facebook.com/note.php?note_id=454991608919>  
[14] *The Underlying Technology of Messages Tech Talk*: [http://www.facebook.com/video/video.php?v=690851516105](https://www.facebook.com/video/video.php?v=690851516105)  
[15] *Facebook’s typeahead search architecture*: [http://www.facebook.com/video/video.php?v=432864835468](https://www.facebook.com/video/video.php?v=432864835468)  
[16] *Facebook Chat*: [http://www.facebook.com/note.php?note\_id=14218138919](https://www.facebook.com/note.php?note_id=14218138919)  
[17] *Who has the most Web Servers?*: <http://www.datacenterknowledge.com/archives/2009/05/14/whos-got-the-most-web-servers/>  
[18] B*uilding Efficient Data Centers with the Open Compute Project*: [http://www.facebook.com/note.php?note\_id=10150144039563920](https://www.facebook.com/note.php?note_id=10150144039563920)  
[19] *Open Compute Project*: <http://opencompute.org/>  
[20] *Facebook’s architecture presentation at Devoxx 2010*: <http://www.devoxx.com>  
[21] *Scaling Facebook to 500 millions users and beyond*: [http://www.facebook.com/note.php?note\_id=409881258919](https://www.facebook.com/note.php?note_id=409881258919)

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
