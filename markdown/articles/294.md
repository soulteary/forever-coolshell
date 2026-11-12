# OSGi和Java企业级运算的未来方向

**摘要**： OSGi也是译者最近才接触到的技术，但是在OSGi的发展中，它越来越收到了来自行业的关注。作为OSGi的动态部署，译者认为此项规范对于企业应用应该是非常有帮助的。特别在银行的信息化建设中将会起到很重要的作用，因为国内大多的银行业都在强调7\*24小时系统，但是其业务发展又非常迅速，常常有新需求，新变更。如果每一次上线变更都将重启系统的话，对银行的服务质量和形象将会造成较大的影响。 此文只是讲述了OSGi在Java企业运算中的新动向，并没有具体的介绍OSGi的规范。关于OSGi规范的文档可以从jcp上下载

原文出处：[这里](http://itknowledgeexchange.techtarget.com/soa-talk/osgi-and-future-directions-for-enterprise-java/)  



**OSGi和Java企业级运算的未来方向**

by Eric Newcomer

无论JCP是否完全的迷失了它的方向，它都不同程度受到来自外部活动的影响。Spring框架和Hibernate影响了EJB3，而且JPA也是一个好的例子。另外日渐感觉到的影响来自于对OSGi规范的采用和其实现，特别是实现了OSGi的开源的Eclipse Equinox，Apache Felix和Knoplerfish框架

OSGi规范为Java定义动态模组元信息系统和在其交互模组中的面向服务的编程模型。这个规范定义了一个为服务查找的注册表，还定义了一组通用功能集合，例如安全，生命周期管理，日志等。OSGi的框架如今已经被Eclipse基金采用，许多的主要Java厂商采用这个规范来开发中间件产品，同时OSGi也被很多开源项目组采用，包括用来开发应用服务器，企业服务总线，和集成开发环境。

作为在商业产品和开源项目中广泛被使用的的核心平台，OSGi联盟开始接收到来自更复杂的的对企业应用的支持需求。在1999年,OSGi规范最初是JSR-8，主要的目的是用于家庭自助网关(home automation gateways)。自从那时起，OSGi技术就被在各种个样自助，移动电话，和家庭娱乐的嵌入应用程序所使用。2006年的8月份，OSGi联盟，接收许多关注于OSGi企业版本的建议并举行一个关于讨论成立一个OSGi企业专家组(EEG）可能性的会议。

自从2007年1月第一次会议一来，OSGi企业专家组EEG用了两年时间编写了致力于使OSGi更好支持企业级Java应用的需求细节和设计细节。这个工作的成果是：在2009年年中，将会对OSGi规范有一个主要的更新(两个的草案版本已经发布)，这个修改主要包括扩展了核心框架服务和定义现有存在企业Java技术与OSGi框架的接口以满足业务应用需求的案例。主要的特性包括被称为蓝图服务(Blueprint Service)Spring框架组件模型到OSGi服务模型的映射和分布计算协议到OSGi服务模型的映射, JavaEE映射的关键部分是Web apps,JDBC,JPA,JMX,JTA,JNDI,和JAAS。

软件行业已经接受并支持OSGi带来的模组化的好处，下一个改进将会是通过适配已经用于企业运算的Java技术接口，进而对企业级Java应用的支撑。这个目标将帮助OSGi的开发人员更容易的以标准的方式创建企业服务务应用程序。

Eric Newcomer是分布计算的专家和独立咨询师，Newcomer是OSGi企业专家组的主席，之前他是IONA技术公司的CTO.他在 [blog on OSGi](https://modualrit.blogspot.com/)发布了很多的OSGi的文章

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
