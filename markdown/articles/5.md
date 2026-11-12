# Java EE6 初探

在[tss](http://www.theserverside.com/ "TSS")上，Reza Rahman发表了一篇关于JAVA EE6[《Java EE6 Overview》](http://www.theserverside.com/tt/articles/article.tss?l=JavaEE6Overview "《Java EE6 Overview》")的文章，在文章里面他谈及一些关于JavaEE6草案的一些修改内容。

想JAVA企业级应用一路走来，从J2ee到Java EE5 在到即将要推出的Java EE6经历了一个由复杂到简单，由繁到简的过程。

Reza Rahman 文章说，Java EE6将会更简单，更轻量级，更易部署，Java EE6将会裁剪到Java EE5中不实用的部分。并且Java EE6将会为不同的用户群提供不同的规范子集。

回顾上一版本Java EE 5中，主要引入了以下改变：  

1）引入了EJB3.0  

2）引入了JSF作Tier framework.  

3）使用JAX-WS2.0取代了JAX-RPC作新一代的SOAP的Web service API  

4）使用POJO编程，零配置系统和自由的XML减轻了系统的复杂性。

而新版本的Java EE 6中，提供了一个更为简单，新型和完美整全的平台，并提供了非常丰富的技术，其包含WebBeans 1.0和JAX-RS 1.1，以及更为成熟的Servlet 3.0。

他废除了如下的API:  

1)JAX-RPC，(被JAX-WS API 取代)  

2)EJB2.x 实体Beans CMP,（被EJB3.0取代）  

3)JAXR, Java EE Appcliation Deployment(JSF-88)  

4)Java EE Management (JSR-77)

他具体包含了如下组件：  

1）WebBeans 1.0：它基本上来说整合了 JSF, JPA 和 EJB 3的编程模型。  

2）JSF 2.0：主要增加了Facelets，Facelets是一个开源的技术，并让JSF支持AJAX，等。  

3）EJB 3.1：主要加入了单实例Bean的根念，支持cron-style调试，异步调用Session Bean的方法，等。  

4）JPA 2.0：主要加了一系列的ORM映射扩展（如maps或lists等），增强了EntityManager和Query的APIs，JPQL支持SQL-like CASE, NULLIF, COALESCE，以及加入了Criteria API。  

5）Servlet 3.0：更为成熟的Servlet引入了Java EE 5模型，更容易配置的web.xml，以及可以通过ServletContext以程序的方式添加Servlets, Filters 和Listeners，等等。  

6）JAX-RS 1.1：REST 开始逐渐成为WEB开发的范例，JAX-RS主要设计目的是简化REST开发，POJO编程和Annotation配置。

有关规范的草案可以<http://jcp.org/en/jsr/detail?id=316>.这里找到

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
