# Python处理中文的时候的一些小技巧

相信第一次处理中文的朋友们可能都会对中文的encoding 和程序的报错很头疼。

如果你像我一样希望能够把事情尽快做好而不去深究，你可能会写一些异常处理的代码把 UnicodeEncodingError糊弄过去先，但当你开始怀疑有多少encoding出错的信息被你丢弃的时候，可能你会很惊奇。于是，你还是会想坐下来，（洗把脸）然后面对自己必须弄懂什么是utf-8，什么是 ‘gb2312’， 什么是 ‘gbk’ 和其中的猫腻。正如有时候猛撕小伤口上邦迪胶布的快感一样，有时候当你认真面对一些你平时一直回避的问题的时候（其实有时候需要的不是勇气）， 你反而会觉得“不过如此”，并且能够一劳永逸的解决问题。

关于Python处理Unicode，我所能找到的最言简意赅的入门教程是：

### [**Unicode* In *Python*, Completely Demystified* （揭秘Python Unicode）](http://farmdev.com/talks/unicode/)

简要罗列一下最重要最实用的点：

Solution

1. **Decode early** （尽早decode, 将文件中的内容转化成 unicode 再进行下一步处理)
2. **Unicode everywhere** (程序内部处理都用unicode)
3. **Encode late** (最后encode回所需的encoding, 例如把最终结果写进结果文件)

**1. Decode early**

> Decode to <type ‘unicode’> ASAP
> 
> >>> def to\_unicode\_or\_bust(
> 
> …         obj, encoding=’utf-8′):
> 
> …     if isinstance(obj, basestring):
> 
> …         if not isinstance(obj, unicode):
> 
> …             obj = unicode(obj, encoding)
> 
> …     return obj
> 
> …
> 
> >>>
> 
> detects if object is a string and if so converts to unicode, if not already.

**2. Unicode everywhere**

> >>> to\_unicode\_or\_bust(ivan\_uni)
> 
> u’Ivan Krsti\u0107′
> 
> >>> to\_unicode\_or\_bust(ivan\_utf8)
> 
> u’Ivan Krsti\u0107′
> 
> >>> to\_unicode\_or\_bust(1234)
> 
> 1234

**3. Encode late**

> Encode to <type ‘str’> when you write to disk or print
> 
> >>> f = open(‘/tmp/ivan\_out.txt’,’w’)
> 
> >>> f.write(ivan\_uni.encode(‘utf-8’))
> 
> >>> f.close()

我以前一直觉得unicode相关的处理都是很 dirty 的工作，一般都会一边尝试，一边用异常处理去补丁，看完以上这个教程以后豁然开朗。

祝大家也能早日理清处理中文的时候的头绪，坦然直面“神秘”的unicode

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
