# erlang打包独立环境

最近公司代码需要在非erlang的系统上执行，需要能在独立的环境里运行erlang。研究甚久，于是写下这篇博文。国内用erlang的朋友不多，希望这篇blog能对有需要的朋友起到参考作用。

> Application-Vsn/ebin  
> 
> /include  
> 
> /priv  
> 
> /src  
> 
> /Application-Vsn.rel

以上是代码的目录表.

> {release, {“nextim”, “2.0”},  
> 
> {erts, “5.7.5”},  
> 
> [{kernel, “2.12.3”},  
> 
> {stdlib, “1.15.3”},  
> 
> {sasl, “2.1.5.3”}]  
> 
> }.

以上是Application-Vsn.rel的内容,[]中是代码本身需要的lib。

1.执行erl -pa ./ebin . 这一步会生成nextim-2.boot文件

> 1> systools:make\_script(nextim-2″, [local]).  
> 
> ok

2.erl -boot nextim-2 . 这一步会生成nextim-2.tar.gz

> systools:make\_tar(“nextim-2”).

3.现在建议把tar.gz文件放到独立的路径里 这样不会影响Application-Vsn文件夹 ，然后解压 并进入目录， 复制erlang系统目录里的 erts-5.7.5 到当前目录

4.建立bin文件夹 复制  `erts-5.7.5/bin/start` 到 `bin/ 编辑 bin/start 改 ROOTDIR为当前目录的路径`

5.复制`erts-5.7.5/bin/run_erl` `和` `erts-5.7.2/bin/erl` `到 bin 并且如同上一步一样修改ROOTDIR.`

6.复制 `$ERLDIR/bin/start_sasl.boot` 到  `bin/start.boot`.

7. `echo` `"5.7.5` `2.0"` `>` `releases/start_erl.data`.

6.执行bin文件里的erl

> release\_handler:create\_RELEASES(“$ROOTDIR”, “$ROOTDIR/releases/”, “$ROOTDIR/releases/nextim-2.rel”, []).

7.再把自己的项目文件复制到lib中  然后启动时 -pa参数是 lib文件夹. 完成这一步，就能独立出erlang环境了。

以上内容 参考自

http://spawnlink.com/articles/an-introduction-to-releases-with-erlybank/

http://streamhacker.com/2009/07/02/how-to-create-an-erlang-first-target-system/

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
