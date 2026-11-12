# 实例分析Java Class的文件结构

**【感谢网友 @[Krq\_Tiger](http://weibo.com/xmuzyq "Krq_Tiger") 投稿】**

今天把之前在Evernote中的笔记重新整理了一下，发上来供对java class 文件结构的有兴趣的同学参考一下。

学习Java的朋友应该都知道Java从刚开始的时候就打着平台无关性的旗号，说“一次编写，到处运行”，其实说到无关性，Java平台还有另外一个无关 性那就是语言无关性，要实现语言无关性，那么Java体系中的class的文件结构或者说是字节码就显得相当重要了，其实Java从刚开始的时候就有两套 规范，一个是Java语言规范，另外一个是Java虚拟机规范，Java语言规范只是规定了Java语言相关的约束以及规则，而虚拟机规范则才是真正从跨 平台的角度去设计的。今天我们就以一个实际的例子来看看，到底Java中一个Class文件对应的字节码应该是什么样子。 这篇文章将首先总体上阐述一下Class到底由哪些内容构成，然后再用一个实际的Java类入手去分析class的文件结构。

在继续之前，我们首先需要明确如下几点：

1）Class文件是有8个字节为基础的字节流构成的，这些字节流之间都严格按照规定的顺序排列，并且字节之间不存在任何空隙，对于超过8个字节的数据，将按 照Big-Endian的顺序存储的，也就是说高位字节存储在低的地址上面，而低位字节存储到高地址上面，其实这也是class文件要跨平台的关键，因为 PowerPC架构的处理采用Big-Endian的存储顺序，而x86系列的处理器则采用Little-Endian的存储顺序，因此为了Class文 件在各中处理器架构下保持统一的存储顺序，虚拟机规范必须对起进行统一。

2） Class文件结构采用类似C语言的结构体来存储数据的，主要有两类数据项，无符号数和表，无符号数用来表述数字，索引引用以及字符串等，比如 u1,u2,u4,u8分别代表1个字节，2个字节，4个字节，8个字节的无符号数，而表是有多个无符号数以及其它的表组成的复合结构。可能大家看到这里 对无符号数和表到底是上面也不是很清楚，不过不要紧，等下面实例的时候，我会再以实例来解释。

明确了上面的两点以后，我们接下来后来看看Class文件中按照严格的顺序排列的字节流都具体包含些什么数据：

![](https://coolshell.cn/uploads/2013/03/1.png "点击查看原始大小图片")

（上图来自The Java Virtual Machine Specification Java SE 7 Edition)

在看上图的时候，有一点我们需要注意，比如cp\_info，cp\_info表示常量池，上图中用 constant\_pool[constant\_pool\_count-1]的方式来表示常量池有constant\_pool\_count-1个常量，它 这里是采用数组的表现形式，但是大家不要误以为所有的常量池的常量长度都是一样的，其实这个地方只是为了方便描述采用了数组的方式，但是这里并不像编程语 言那里，一个int型的数组，每个int长度都一样。明确了这一点以后，我们在回过头来看看上图中每一项都具体代表了什么含义。

1）u4 magic 表示魔数，并且魔数占用了4个字节，魔数到底是做什么的呢？它其实就是表示一下这个文件的类型是一个Class文件，而不是一张JPG图片，或者AVI的电影。而Class文件对应的魔数是0xCAFEBABE.

2）u2 minor\_version 表示Class文件的次版本号，并且此版本号是u2类型的无符号数表示。

3） u2 major\_version 表示Class文件的主版本号，并且主版本号是u2类型的无符号数表示。major\_version和minor\_version主要用来表示当前的虚拟 机是否接受当前这种版本的Class文件。不同版本的Java编译器编译的Class文件对应的版本是不一样的。高版本的虚拟机支持低版本的编译器编译的 Class文件结构。比如Java SE 6.0对应的虚拟机支持Java SE 5.0的编译器编译的Class文件结构，反之则不行。

4） u2 constant\_pool\_count 表示常量池的数量。这里我们需要重点来说一下常量池是什么东西，请大家不要与Jvm内存模型中的运行时常量池混淆了，Class文件中常量池主要存储了字 面量以及符号引用，其中字面量主要包括字符串，final常量的值或者某个属性的初始值等等，而符号引用主要存储类和接口的全限定名称，字段的名称以及描 述符，方法的名称以及描述符，这里名称可能大家都容易理解，至于描述符的概念，放到下面说字段表以及方法表的时候再说。另外大家都知道Jvm的内存模型中 有堆，栈，方法区，程序计数器构成，而方法区中又存在一块区域叫运行时常量池，运行时常量池中存放的东西其实也就是编译器长生的各种字面量以及符号引用， 只不过运行时常量池具有动态性，它可以在运行的时候向其中增加其它的常量进去，最具代表性的就是String的intern方法。

5）cp\_info 表示常量池，这里面就存在了上面说的各种各样的字面量和符号引用。放到常量池的中数据项在The Java Virtual Machine Specification Java SE 7 Edition 中一共有14个常量，每一种常量都是一个表，并且每种常量都用一个公共的部分tag来表示是哪种类型的常量。

下面分别简单描述一下具体细节等到后面的实例 中我们再细化。

- CONSTANT\_Utf8\_info      tag标志位为1,   UTF-8编码的字符串
- CONSTANT\_Integer\_info  tag标志位为3， 整形字面量
- CONSTANT\_Float\_info     tag标志位为4， 浮点型字面量
- CONSTANT\_Long\_info     tag标志位为5， 长整形字面量
- CONSTANT\_Double\_info  tag标志位为6， 双精度字面量
- CONSTANT\_Class\_info    tag标志位为7， 类或接口的符号引用
- CONSTANT\_String\_info    tag标志位为8，字符串类型的字面量
- CONSTANT\_Fieldref\_info  tag标志位为9,  字段的符号引用
- CONSTANT\_Methodref\_info  tag标志位为10，类中方法的符号引用
- CONSTANT\_InterfaceMethodref\_info tag标志位为11, 接口中方法的符号引用
- CONSTANT\_NameAndType\_info tag 标志位为12，字段和方法的名称以及类型的符号引用

6） u2 access\_flags 表示类或者接口的访问信息，具体如下图所示：  

![](https://coolshell.cn/uploads/2013/03/2.png "点击查看原始大小图片")

7）u2 this\_class 表示类的常量池索引，指向常量池中CONSTANT\_Class\_info的常量

8）u2 super\_class 表示超类的索引，指向常量池中CONSTANT\_Class\_info的常量

9）u2 interface\_counts 表示接口的数量

10）u2 interface[interface\_counts]表示接口表，它里面每一项都指向常量池中CONSTANT\_Class\_info常量

11）u2 fields\_count 表示类的实例变量和类变量的数量

12） field\_info fields[fields\_count]表示字段表的信息，其中字段表的结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/3.png)

上图中access\_flags表示字段的访问表示，比如字段是public,private，protect 等，name\_index表示字段名 称，指向常量池中类型是CONSTANT\_UTF8\_info的常量，descriptor\_index表示字段的描述符，它也指向常量池中类型为 CONSTANT\_UTF8\_info的常量，attributes\_count表示字段表中的属性表的数量，而属性表是则是一种用与描述字段，方法以及 类的属性的可扩展的结构，不同版本的Java虚拟机所支持的属性表的数量是不同的。

13） u2 methods\_count表示方法表的数量

14）method\_info 表示方法表，方法表的具体结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/4.png)  

其中access\_flags表示方法的访问表示，name\_index表示名称的索引，descriptor\_index表示方法的描述 符，attributes\_count以及attribute\_info类似字段表中的属性表，只不过字段表和方法表中属性表中的属性是不同的，比如方法 表中就Code属性，表示方法的代码，而字段表中就没有Code属性。其中具体Class中到底有多少种属性，等到Class文件结构中的属性表的时候再 说说。

15） attribute\_count表示属性表的数量，说到属性表，我们需要明确以下几点：

- 属性表存在于Class文件结构的最后，字段表，方法表以及Code属性中，也就是说属性表中也可以存在属性表
- 属性表的长度是不固定的，不同的属性，属性表的长度是不同的

上面说完了Class文件结构中每一项的构成以后，我们以一个实际的例子来解释以下上面所说的内容。

```java
package com.ejushang.TestClass;

public class TestClass implements Super{

private static final int staticVar = 0;

private int instanceVar=0;

public int instanceMethod(int param){
 return param+1;
 }

}

interface Super{ }
```

通过jdk1.6.0\_37的javac 编译后的TestClass.java对应的TestClass.class的二进制结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/5.png)

下面我们就根据前面所说的Class的文件结构来解析以下上图中字节流。

**1）魔数**  

从Class的文件结构我们知道，刚开始的4个字节是魔数，上图中从地址00000000h-00000003h的内容就是魔数，从上图可知Class的文件的魔数是0xCAFEBABE。

 **2）主次版本号**  

接下来的4个字节是主次版本号，有上图可知从00000004h-00000005h对应的是0x0000,因此Class的minor\_version 为0x0000,从00000006h-00000007h对应的内容为0x0032,因此Class文件的major\_version版本为 0x0032,这正好就是jdk1.6.0不带target参数编译后的Class对应的主次版本。

 **3）常量池的数量**  

接下来的2个字节从00000008h-00000009h表示常量池的数量，由上图可以知道其值为0x0018，十进制为24个,但是对于常量池的数量 需要明确一点，常量池的数量是constant\_pool\_count-1，为什么减一，是因为索引0表示class中的数据项不引用任何常量池中的常 量。

 **4）常量池**  

我们上面说了常量池中有不同类型的常量，下面就来看看TestClass.class的第一个常量，我们知道每个常量都有一个u1类型的tag标识来表示 常量的类型，上图中0000000ah处的内容为0x0A，转换成二级制是10，有上面的关于常量类型的描述可知tag为10的常量是Constant\_Methodref\_info,而Constant\_Methodref\_info的结够如下图所示：

![](https://coolshell.cn/uploads/2013/03/6.png)

其中class\_index指向常量池中类型为CONSTANT\_Class\_info的常量，从TestClass的二进制文件结构中可以看出 class\_index的值为0x0004（地址为0000000bh-0000000ch)，也就是说指向第四个常量。

name\_and\_type\_index指向常量池中类型为CONSTANT\_NameAndType\_info常量。从上图可以看出name\_and\_type\_index的值为0x0013，表示指向常量池中的第19个常量。

接下来又可以通过同样的方法来找到常量池中的所有常量。不过JDK提供了一个方便的工具可以让我们查看常量池中所包含的常量。通过javap -verbose TestClass 即可得到所有常量池中的常量，截图如下：

![](https://coolshell.cn/uploads/2013/03/7.png)

从上图我们可以清楚的看到，TestClass中常量池有24个常量，不要忘记了第0个常量，因为第0个常量被用来表示 Class中的数据项不引用任何常量池中的常量。从上面的分析中我们得知TestClass的第一个常量表示方法，其中class\_index指向的第四 个常量为java/lang/Object，name\_and\_type\_index指向的第19个常量值为<init>:()V,从这里可 以看出第一个表示方法的常量表示的是java编译器生成的实例构造器方法。通过同样的方法可以分析常量池的其它常量。OK，分析完常量池，我们接下来再分 析下access\_flags。  

**5）u2 access\_flags** 表示类或者接口方面的访问信息，比如Class表示的是类还是接口，是否为public,static，final等。具体访问标示的含义之前已经说过 了，下面我们就来看看TestClass的访问标示。Class的访问标示是从0000010dh-0000010e，期值为0x0021，根据前面说的 各种访问标示的标志位，我们可以知道：0x0021=0x0001|0x0020 也即ACC\_PUBLIC 和 ACC\_SUPER为真，其中ACC\_PUBLIC大家好理解，ACC\_SUPER是jdk1.2之后编译的类都会带有的标志。

**6）u2 this\_class** 表示类的索引值，用来表示类的全限定名称，类的索引值如下图所示：

![](https://coolshell.cn/uploads/2013/03/8.png)

从上图可以清楚到看到，类索引值为0x0003，对应常量池的第三个常量，通过javap的结果，我们知道第三个常量为 CONSTANT\_Class\_info类型的常量，通过它可以知道类的全限定名称为：com/ejushang/TestClass /TestClass

 **7）u2 super\_class** 表示当前类的父类的索引值，索引值所指向的常量池中类型为CONSTANT\_Class\_info的常量，父类的索引值如下图所示，其值为0x0004, 查看常量池的第四个常量，可知TestClass的父类的全限定名称为：java/lang/Object

![](https://coolshell.cn/uploads/2013/03/9.png)

**8）interfaces\_count和  interfaces[interfaces\_count]**表示接口数量以及具体的每一个接口，TestClass的接口数量以及接口如下图所示，其中 0x0001表示接口数量为1，而0x0005表示接口在常量池的索引值，找到常量池的第五个常量，其类型为CONSTANT\_Class\_info，其 值为：com/ejushang/TestClass/Super

![](https://coolshell.cn/uploads/2013/03/10.png)

**9）fields\_count 和 field\_info**, fields\_count表示类中field\_info表的数量，而field\_info表示类的实例变量和类变量，这里需要注意的是 field\_info不包含从父类继承过来的字段，field\_info的结构如下图所示：  

![](https://coolshell.cn/uploads/2013/03/11.png)

其中access\_flags表示字段的访问标示，比如public,private,protected，static,final等，access\_flags的取值如下图所示：  

![](https://coolshell.cn/uploads/2013/03/12.png "点击查看原始大小图片")

其中name\_index 和 descriptor\_index都是常量池的索引值，分别表示字段的名称和字段的描述符，字段的名称容易理解，但是字段的描述符如何理解呢？其实在JVM 规范中，对于字段的描述符规定如下图所示：  

![](https://coolshell.cn/uploads/2013/03/13.png "点击查看原始大小图片")  

其中大家需要关注一下上图最后一行，它表示的是对一维数组的描述符，对于String[][]的描述符将是[[ Ljava/lang/String,而对于int[][]的描述符为[[I。接下来的attributes\_count以及 attribute\_info分别表示属性表的数量以及属性表。下面我们还是以上面的TestClass为例，来看看TestClass的字段表吧。

首先我们来看一下字段的数量，TestClass的字段的数量如下图所示：

![](https://coolshell.cn/uploads/2013/03/14.png)

从上图中可以看出TestClass有两个字段，查看TestClass的源代码可知，确实也只有两个字段，接下来我们看看第一个字段，我们知道第一个字段应该为private int staticVar,它在Class文件中的二进制表示如下图所示：

![](https://coolshell.cn/uploads/2013/03/15.png)  

其中0x001A表示访问标示，通过查看access\_flags表可知，其为ACC\_PRIVATE,ACC\_STATIC,ACC\_FINAL,接下 来0x0006和0x0007分别表示常量池中第6和第7个常量，通过查看常量池可知，其值分别为：staticVar和I，其中staticVar为字 段名称，而I为字段的描述符，通过上面对描述符的解释，I所描述的是int类型的变量，接下来0x0001表示staticVar这个字段表中的属性表的 数量，从上图可以staticVar字段对应的属性表有1个，0x0008表示常量池中的第8个常量，查看常量池可以得知此属性为 ConstantValue属性，而ConstantValue属性的格式如下图所示：  

![](https://coolshell.cn/uploads/2013/03/16.png)

其中attribute\_name\_index表述属性名的常量池索引，本例中为ConstantValue，而ConstantValue的 attribute\_length固定长度为2，而constantValue\_index表示常量池中的引用，本例中，其中为0x0009，查看第9个 常量可以知道，它表示一个类型为CONSTANT\_Integer\_info的常量，其值为0。

上面说完了private static final int staticVar=0，下面我们接着说一下TestClass的private int instanceVar=0,在本例中对instanceVar的二进制表示如下图所示：

![](https://coolshell.cn/uploads/2013/03/17.png)  

其中0x0002表示访问标示为ACC\_PRIVATE,0x000A表示字段的名称，它指向常量池中的第10个常量，查看常量池可以知道字段名称为 instanceVar，而0x0007表示字段的描述符，它指向常量池中的第7个常量，查看常量池可以知道第7个常量为I，表示类型为 instanceVar的类型为I，最后0x0000表示属性表的数量为0.

 **10）methods\_count 和 method\_info** ，其中methods\_count表示方法的数量，而method\_info表示的方法表，其中方法表的结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/18.png)

从上图可以看出method\_info和field\_info的结构是很类似的，方法表的access\_flag的所有标志位以及取值如下图所示：

![](https://coolshell.cn/uploads/2013/03/19.png "点击查看原始大小图片")

其中name\_index和descriptor\_index表示的是方法的名称和描述符，他们分别是指向常量池的索引。这里需要结解释一下方法的描述 符，方法的描述符的结构为：（参数列表）返回值，比如public int instanceMethod(int param)的描述符为：（I）I，表示带有一个int类型参数且返回值也为int类型的方法，接下来就是属性数量以及属性表了，方法表和字段表虽然都有 属性数量和属性表，但是他们里面所包含的属性是不同。接下来我们就以TestClass来看一下方法表的二进制表示。首先来看一下方法表数量，截图如下：

![](https://coolshell.cn/uploads/2013/03/20.png)  

从上图可以看出方法表的数量为0x0002表示有两个方法，接下来我们来分析第一个方法，我们首先来看一下TestClass的第一个方法的access\_flag，name\_index,descriptor\_index，截图如下：

![](https://coolshell.cn/uploads/2013/03/21.png)  

从上图可以知道access\_flags为0x0001，从上面对access\_flags标志位的描述，可知方法的access\_flags的取值为 ACC\_PUBLIC,name\_index为0x000B，查看常量池中的第11个常量，知道方法的名称为<init>，0x000C表示 descriptor\_index表示常量池中的第12常量，其值为()V,表示<init>方法没有参数和返回值，其实这是编译器自动生成 的实例构造器方法。接下来的0x0001表示<init>方法的方法表有1个属性，属性截图如下：  

![](https://coolshell.cn/uploads/2013/03/22.png)  

从上图可以看出0x000D对应的常量池中的常量为Code,表示的方法的Code属性，所以到这里大家应该明白方法的那些代码是存储在Class文件方法表中的属性表中的Code属性中。接下来我们在分析一下Code属性，Code属性的结构如下图所示：  

![](https://coolshell.cn/uploads/2013/03/23.png)

其中attribute\_name\_index指向常量池中值为Code的常量，attribute\_length的长度表示Code属性表的长度（这里 需要注意的时候长度不包括attribute\_name\_index和attribute\_length的6个字节的长度）。

max\_stack表示最大栈深度，虚拟机在运行时根据这个值来分配栈帧中操作数的深度，而max\_locals代表了局部变量表的存储空间。

max\_locals的单位为slot，slot是虚拟机为局部变量分配内存的最小单元，在运行时，对于不超过32位类型的数据类型，比如 byte,char,int等占用1个slot，而double和Long这种64位的数据类型则需要分配2个slot，另外max\_locals的值并 不是所有局部变量所需要的内存数量之和，因为slot是可以重用的，当局部变量超过了它的作用域以后，局部变量所占用的slot就会被重用。

code\_length代表了字节码指令的数量，而code表示的时候字节码指令，从上图可以知道code的类型为u1,一个u1类型的取值为0x00-0xFF,对应的十进制为0-255，目前虚拟机规范已经定义了200多条指令。

exception\_table\_length以及exception\_table分别代表方法对应的异常信息。

attributes\_count和attribute\_info分别表示了Code属性中的属性数量和属性表，从这里可以看出Class的文件结构中，属性表是很灵活的，它可以存在于Class文件，方法表，字段表以及Code属性中。

接下来我们继续以上面的例子来分析一下，从上面init方法的Code属性的截图中可以看出，属性表的长度为0x00000026,max\_stack的 值为0x0002,max\_locals的取值为0x0001,code\_length的长度为0x0000000A，那么00000149h- 00000152h为字节码，接下来exception\_table\_length的长度为0x0000，而attribute\_count的值为 0x0001，00000157h-00000158h的值为0x000E,它表示常量池中属性的名称，查看常量池得知第14个常量的值为 LineNumberTable，LineNumberTable用于描述java源代码的行号和字节码行号的对应关系，它不是运行时必需的属性，如果通 过-g:none的编译器参数来取消生成这项信息的话，最大的影响就是异常发生的时候，堆栈中不能显示出出错的行号，调试的时候也不能按照源代码来设置断 点，接下来我们再看一下LineNumberTable的结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/24.png)

其中attribute\_name\_index上面已经提到过，表示常量池的索引，attribute\_length表示属性长度，而start\_pc和 line\_number分表表示字节码的行号和源代码的行号。本例中LineNumberTable属性的字节流如下图所示：

![](https://coolshell.cn/uploads/2013/03/25.png)

上面分析完了TestClass的第一个方法，通过同样的方式我们可以分析出TestClass的第二个方法，截图如下：

![](https://coolshell.cn/uploads/2013/03/26.png)

其中access\_flags为0x0001,name\_index为0x000F,descriptor\_index为0x0010，通过查看常量池可 以知道此方法为public int instanceMethod(int param)方法。通过和上面类似的方法我们可以知道instanceMethod的Code属性为下图所示：

![](https://coolshell.cn/uploads/2013/03/27.png)

最后我们来分析一下，Class文件的属性，从00000191h-00000199h为Class文件中的属性表，其中0x0011表示属性的名称，查看常量池可以知道属性名称为SourceFile，我们再来看看SourceFile的结构如下图所示：

![](https://coolshell.cn/uploads/2013/03/28.png)

其中attribute\_length为属性的长度，sourcefile\_index指向常量池中值为源代码文件名称的常量，在本例中SourceFile属性截图如下：

![](https://coolshell.cn/uploads/2013/03/29.png)  

其中attribute\_length为0x00000002表示长度为2个字节，而soucefile\_index的值为0x0012,查看常量池的第18个常量可以知道源代码文件的名称为TestClass.java

最后，希望对技术感兴趣的朋友多交流。个人微博：（<http://weibo.com/xmuzyq>)

(全文完)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
