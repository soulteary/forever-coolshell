# 早期XML Schema中的open content模型

**摘要**：在看SDO的一些规范文档，可能会出现open content这样的词组，上网查了相关资料，发现这是一种XML Schema的模型，本文就描述了XML Schema的Open Content模型的含义，在最新的XML Schema规范中，好像已经没有Open模型，它的等价物是any模型。

早期发布的XML Schema规范中支持一种新的element定义，在这个定义中，你可以将XML的Element的内容定义为开放的。下面我们将会介绍一下XML的Open Content 模型。

在Open Content模型中，如果一个XML的元素在XML Schema中被声明为开放的，那么这个Schema对应的XML文档的实例就可以包含一个没有在Schema中罗列的子元素。例如，一个包含着如下的XML Schema的Schema文件

```xml
      <element name=&quot;Book&quot;>
           <type>
               <element name=&quot;Title&quot; type=&quot;string&quot;/>
               <element name=&quot;Author&quot; type=&quot;string&quot;/>
               <element name=&quot;Date&quot; type=&quot;string&quot;/>
               <element name=&quot;ISBN&quot; type=&quot;string&quot;/>
               <element name=&quot;Publisher&quot; type=&quot;string&quot;/>
           </type>
      </element>
```

这个book element的声明意味着这个Schema的实例XML文件必须包含5个元素 – Title,Author，Date，ISBN，Pbulish。例如：

```xml
     <Book>
         <Title>Illusions The Adventures of a Reluctant Messiah</Title>
         <Author>Richard Bach</Author>
         <Date>1977</Date>
         <ISBN>0-440-34319-4</ISBN>
         <Publisher>Dell Publishing Co.</Publisher>
     </Book>


```

假设，在实例XML文件，你希望增加book的另外一个子元素，比如，你希望增加一个到某一个网页的连接：

```xml
     <Book>
         <Title>Illusions The Adventures of a Reluctant Messiah</Title>
         <Author>Richard Bach</Author>
         <Date>1977</Date>
         <ISBN>0-440-34319-4</ISBN>
         <Publisher>Dell Publishing Co.</Publisher>
         <AuthorsWebPage xlink:href=&quot;<a href=&quot;http://www.rbach.com&quot;/&quot;>http://www.rbach.com&quot;/</a>>
    </Book>


```

对于上面这个XML文件，XML Schema分析器将会认为这个XML文件是无效的XML，因为上面的文件的包含了Scheme中没有定义的元素。但是在我们的应用场景中，我们可能会希望XML Schema分析器不要报错，因为，应用程序自己知道如何处理<AuthorsWebPage>这个元素。为了达到这个目的，我们就可以将Book声明为开放的。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
