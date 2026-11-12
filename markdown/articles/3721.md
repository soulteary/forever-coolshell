# Stack Exchange 的架构

近日，Stack Exchange系统管理员blog上发布了一篇关于[Stack Exchange的架构一瞥](http://blog.serverfault.com/post/stack-exchanges-architecture-in-bullet-points/)，其包括了Stack Overflow, Server Fault 和 Super User的 Stack Exchange 网络。注意最后一个关于人员的配置。希望能给大家一些相关的参考。

目录


- [网络流量](#%E7%BD%91%E7%BB%9C%E6%B5%81%E9%87%8F "网络流量")
- [数据中心](#%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83 "数据中心")
- [生产服务器](#%E7%94%9F%E4%BA%A7%E6%9C%8D%E5%8A%A1%E5%99%A8 "生产服务器")
- [使用了的相关的软件和技术](#%E4%BD%BF%E7%94%A8%E4%BA%86%E7%9A%84%E7%9B%B8%E5%85%B3%E7%9A%84%E8%BD%AF%E4%BB%B6%E5%92%8C%E6%8A%80%E6%9C%AF "使用了的相关的软件和技术")
- [程序员和系统管理员](#%E7%A8%8B%E5%BA%8F%E5%91%98%E5%92%8C%E7%B3%BB%E7%BB%9F%E7%AE%A1%E7%90%86%E5%91%98 "程序员和系统管理员")
#### 网络流量

- 每月9千5百万个PV
- 每秒800 HTTP 请求
- 每秒180 DNS 请求
- 每秒55Mb 的带宽

#### 数据中心

- 1 机柜 位于俄勒冈的 [Peak Internet](http://www.peakinternet.com/) (用于[chat](http://chat.stackexchange.com/) 和[Data Explorer](http://data.stackexchange.com/))
- 2 机框 位于 纽约的 [Peer 1](http://www.peer1.com/) ( 用于其它的 Stack Exchange Network)

#### 生产服务器

- 12 Web Servers (Windows Server 2008 R2)
- 2 Database Servers (Windows Server 2008 R2 and SQL Server 2008 R2)
- 2 Load Balancers (Ubuntu Server and HAProxy)
- 2 Caching Servers (Redis on CentOS)
- 1 Router / Firewall (Ubuntu Server)
- 3 DNS Servers (Bind on CentOS)

(生产服务器不含故障备份和管理服务器)

#### 使用了的相关的软件和技术

- [C# / .NET](http://www.microsoft.com/net/)
- [Windows Server 2008 R2](http://www.microsoft.com/windowsserver2008/en/us/default.aspx)
- [SQL Server 2008 R2](http://www.microsoft.com/sqlserver/en/us/default.aspx)
- [Ubuntu Server](http://www.ubuntu.com/server)
- [CentOS](http://www.centos.org/)
- [HAProxy](http://haproxy.1wt.eu/) 用于负载均衡
- [Redis](http://redis.io/) 用于缓存
- [CruiseControl.NET](http://sourceforge.net/projects/ccnet/) 用于做builds
- [Lucene.NET](http://lucene.apache.org/lucene.net/) 用于搜索
- [Bacula](http://www.bacula.org/en/) 用于做备份
- [Nagios](http://www.nagios.org/) (with n2rrd and drraw plugins) 用于系统监控
- [Splunk](http://www.splunk.com/) 用于日志
- [SQL Monitor from Red Gate](http://www.red-gate.com/products/dba/sql-monitor/) 用于监控SQL Server
- [Mercurial](http://mercurial.selenic.com/) / [Kiln](http://www.fogcreek.com/kiln/) 用于源码管理
- [Bind](http://www.isc.org/software/bind) 用于 DNS

#### 程序员和系统管理员

- 14 程序员
- 2 系统管理员

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
