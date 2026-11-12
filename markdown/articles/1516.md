# mochiweb参数化模型Req相关功能

本文的笔记讲述如何从client请求中获取各种参数,如method, request path, headers, cookie等。

Mochiweb是Erlang实现的一个开源Web服务器，它设计的一个亮点就是他本身的Http请求的参数化模型。因此我们可以用OO的方式来理解它的相关用法。  

它的实现在mochiweb\_request模块.在mochiweb中,每个client请求其构造一个 Req 对象(注：这个“对象“只是便于理解的提法), Req 可以理解成 mochiweb\_request 的一个参数化或实例化.  



1.**Req:get(method)**-> ‘OPTIONS’ | ‘GET’ | ‘HEAD’ | ‘POST’ | ‘PUT’ | ‘DELETE’ | ‘TRACE’.  

获取Http请求的方式.

2.**Req:get(raw\_path)** -> String().  

获取raw\_path.比如 http://www.nextim.cn/session/login?username=test#p,那/session/login?username=test#p就是这个raw\_path.

3.**Req:get(path)**-> String().  

获取path.比如 http://www.nextim.cn/session/login?username=test#p,那/session/login就是这个raw\_path.

4.**Req:parse\_qs()** -> [{strng(), string()}].  

获取get参数.比如 http://www.nextim.cn/session/login?username=test#p,则返回[{“username”,”test”}].

5.**Req:parse\_post()** -> [{strng(), string()}].  

确保post数据类型为: application/x-www-form-urlencoded, 否则不要调用(其内部会调用Req:recv\_body),返回值类似Req:parse\_qs().

6.**Req:get(peer)**-> string().  

返回值为client的ip

7.**Req:get\_header\_value(Key)** -> undefined | string().  

获取某个header,比如Key为”User-Agent”时，返回”Mozila…….”

8.**Req:get\_primary\_header\_value(Key)** -> undefined | string().  

获取http headers中某个key对应的主值(value以分号分割).  

举例: 假设 Content-Type 为 application/x-www-form-urlencoded; charset=utf8,则  

Req:get\_header\_value(“content-type”) 返回 application/x-www-form-urlencoded

9.**Req:get(headers)** -> dict().  

获取所有headers  

说明: 返回结果为stdlib/dict 数据结构,可以通过mochiweb\_headers模块进行操作.  

举例: 下面代码显示请求中所有headers:  

Headers = Req:get(headers),  

lists:foreach(fun(Key, Value) ->  

io:format(“~p : ~p ~n”, [Key, Value])  

end,  

mochiweb\_headers:to\_list(Headers)).

10.**Req:parse\_cookie()** -> [{string(), string()}].  

解析Cookie

11.**R****eq:get\_cookie\_value(Key)**-> string().  

类似Req:get\_header\_value(Key)

最近搜了下，发现用mochiweb的挺多的。但自己用的时候发现来不少疑难。以上文档皆由litaocheng总结提供。感谢所带来的帮助。希望这个对国内使用mochiweb的朋友们带来帮助。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
