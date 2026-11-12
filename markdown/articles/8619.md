# 你可能不知道的Shell

Shell也叫做命令行界面，它是\*nix操作系统下用户和计算机的交互界面。Shell这个词是指操作系统中提供访问内核服务的程序。

这篇文章向大家介绍Shell一些非广为人知、但却实用有趣的知识，权当品尝shell主食后的甜点吧。

#### 科普

先科普几个你可能不知道的事实：

- Shell几乎是和Unix操作系统一起诞生，第一个Unix Shell是肯·汤普逊（Ken Thompson）以Multics上的Shell为模范在1971年改写而成，并命名Thompson sh。即便是后来流行的bash（shell的一种变体），它的年龄实际上比当前流行的所有的Linux kernel都大，可谓在Linux系统上是先有Shell再有Kernel。

- 当前绝大部分\*nix和MacOS操作系统里的默认的Shell都是bash，bash由Brian Fox在1987年创造，全称Bourne Again shell ( bash)。

- 你或许听说除了bash之外，还有Bourne shell ( sh)，Korn shell ( ksh)，C shell （包括 csh and tcsh），但是你知道这个星球上一共存在着大约50多种不同的shell么？想了解他们，请参考 <http://www.freebsd.org/ports/shells.html>。

- 每个月[tiobe](http://www.tiobe.com/index.php/content/paperinfo/tpci/index.html)上都会给一个编程语言的排名，来显示各种语言的流行度。排名指数综合了全球范围内使用该语言的工程师人数、教学的课程数和第三方供应商数。截止至2012年11月份，tiobe公布的编程语言排行榜里，bash的指数是0.56%排名22位。如果算上它旗下的awk 0.21%和tcl 0.146%，大概就能排到14名。注意这里还不包括bash的同源的兄弟姐妹csh、ksh等，算上它们，shell家族有望接近前十。值得一提的是一直以来shell的排名就很稳定，不像某些“暴发户”语言，比如objective-c，这些语言的流行完全是因为当前Apple系的崛起，但这种热潮极有可能来得快去得更快。

![](https://coolshell.cn/uploads/2012/11/shell.01.png)

![](https://coolshell.cn/uploads/2012/11/shell.02.png)

全球最大的源代码仓库Github里，shell相关的项目数占到了8%，跻身前5和Java相当，可见在实战工程里，shell可谓宝刀不老。图片来源，[参见这里](https://github.com/languages)

![](https://coolshell.cn/uploads/2012/11/shell.03.png)

#### 一些强大的命令

再分享一些可能你不知道的shell用法和脚本，简单&强大！

*在阅读以下部分前，强烈建议读者打开一个shell实验，这些都不是shell教科书里的大路货哦：）*

- **`!$`**  
  
  `!$`是一个特殊的环境变量，它代表了上一个命令的最后一个字符串。如：你可能会这样：
  ```shell
  $mkdir mydir
  $mv mydir yourdir
  $cd yourdir
  ```
  
  可以改成：
  
  ```shell
  $mkdir mydir
  $mv !$ yourdir
  $cd !$
  ```

- **`sudo !!`**  
  
  以root的身份执行上一条命令 。  
  
  场景举例：比如Ubuntu里用`apt-get`安装软件包的时候是需要root身份的，我们经常会忘记在`apt-get`前加`sudo`。每次不得不加上`sudo`再重新键入这行命令，这时可以很方便的用`sudo !!`完事。  
  
  （陈皓注：在shell下，有时候你会输入很长的命令，你可以使用!xxx来重复最近的一次命令，比如，你以前输入过，vi /where/the/file/is, 下次你可以使用 !vi 重得上次最近一次的vi命令。）

- **`cd –`**  
  
  回到上一次的目录 。  
  
  场景举例：当前目录为`/home/a`，用`cd ../b`切换到`/home/b`。这时可以通过反复执行`cd –`命令在`/home/a`和`/home/b`之间来回方便的切换。  
  
  （陈皓注：cd ~ 是回到自己的Home目录，cd ~user，是进入某个用户的Home目录）

- **`'ALT+.' or '<ESC> .'`**  
  
  热建alt+. 或 esc+. 可以把上次命令行的参数给重复出来。

- **`^old^new`**  
  
  替换前一条命令里的部分字符串。  
  
  场景：`echo "wanderful"`，其实是想输出`echo "wonderful"`。只需要`^a^o`就行了，对很长的命令的错误拼写有很大的帮助。（陈皓注：也可以使用 **!!:gs/old/new**）

- **`du -s * | sort -n | tail`**  
  
  列出当前目录里最大的10个文件。

- **`:w !sudo tee %`**  
  
  在vi中保存一个只有root可以写的文件

- **`date [[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)`**  
  
  时间截转时间

- **`> file.txt`**  
  
  创建一个空文件，比touch短。

- **`mtr coolshell.cn`**  
  
  mtr命令比traceroute要好。

- 在命令行前加空格，该命令不会进入history里。

- **`echo "ls -l" | at midnight`**  
  
  在某个时间运行某个命令。

- **`curl -u user:pass -d status="Tweeting from the shell" http://twitter.com/statuses/update.xml`**  
  
  命令行的方式更新twitter。

- **`curl -u username --silent "https://mail.google.com/mail/feed/atom" | perl -ne 'print "\t" if /<name>/; print "$2\n" if /<(title|name)>(.*)<\/\1>/;'`**  
  
  检查你的gmail未读邮件

- **`ps aux | sort -nk +4 | tail`**  
  
  列出头十个最耗内存的进程

- **`man ascii`**  
  
  显示ascii码表。  
  
  场景：忘记ascii码表的时候还需要google么?尤其在天朝网络如此“顺畅”的情况下，就更麻烦在GWF多应用一次规则了，直接用本地的`man ascii`吧。

- **`ctrl-x e`**  
  
  快速启动你的默认编辑器（由变量$EDITOR设置）。

- **`netstat –tlnp`**  
  
  列出本机进程监听的端口号。（陈皓注：netstat -anop 可以显示侦听在这个端口号的进程）

- **`tail -f /path/to/file.log | sed '/^Finished: SUCCESS$/ q'`**  
  
  当file.log里出现Finished: SUCCESS时候就退出tail，这个命令用于实时监控并过滤log是否出现了某条记录。

- **`ssh [[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) bash < /path/to/local/script.sh`**  
  
  在远程机器上运行一段脚本。这条命令最大的好处就是不用把脚本拷到远程机器上。

- **`ssh [[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection) cat /path/to/remotefile | diff /path/to/localfile -`**  
  
  比较一个远程文件和一个本地文件

- **`net rpc shutdown -I ipAddressOfWindowsPC -U username%password`**  
  
  远程关闭一台Windows的机器

- **`screen -d -m -S some_name ping my_router`**  
  
  后台运行一段不终止的程序，并可以随时查看它的状态。`-d -m`参数启动“分离”模式，`-S`指定了一个session的标识。可以通过`-R`命令来重新“挂载”一个标识的session。更多细节请参考screen用法 `man screen`。

- **`wget --random-wait -r -p -e robots=off -U mozilla http://www.example.com`**  
  
  下载整个www.example.com网站。（注：别太过分，大部分网站都有防爬功能了：））

- **`curl ifconfig.me`**  
  
  当你的机器在内网的时候，可以通过这个命令查看外网的IP。

- **`convert input.png -gravity NorthWest -background transparent -extent 720x200  output.png`**  
  
  改一下图片的大小尺寸

- **`lsof –i`**  
  
  实时查看本机网络服务的活动状态。

- **`vim scp://[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)//path/to/somefile`**  
  
  vim一个远程文件

- **`python -m SimpleHTTPServer`**  
  
  一句话实现一个HTTP服务，把当前目录设为HTTP服务目录，可以通过`http://localhost:8000`访问 这也许是这个星球上最简单的HTTP服务器的实现了。

- **`history | awk '{CMD[$2]++;count++;} END { for (a in CMD )print CMD[a] " " CMD[a]/count*100 "% " a }' | grep -v "./" | column -c3 -s " " -t | sort -nr | nl | head -n10`**  
  
  (陈皓注：有点复杂了，history|awk ‘{print $2}’|awk ‘BEGIN {FS=”|”} {print $1}’|sort|uniq -c|sort -rn|head -10)  
  
  这行脚本能输出你最常用的十条命令，由此甚至可以洞察你是一个什么类型的程序员。

- **`tr -c "[:digit:]" " " < /dev/urandom | dd cbs=$COLUMNS conv=unblock | GREP_COLOR="1;32" grep --color "[^ ]"`**  
  
  想看看Marix的屏幕效果吗？（不是很像，但也很Cool!）

看不懂行代码？没关系，系统的学习一下\*nix shell脚本吧，力荐[《Linux命令行与Shell脚本编程大全》](http://www.ituring.com.cn/book/980)。

#### 参考文献：

- [Unix Shell Wiki](https://en.wikipedia.org/wiki/Unix_shell#Shell_categories)
- [Github language ranking](https://github.com)
- [An introduction of Unix Shell history](http://www.softpanorama.org/People/Shell_giants/introduction.shtml)
- [Tiobe Software](http://www.tiobe.com/index.php/content/paperinfo/tpci/index.html)
- <http://www.commandlinefu.com/>

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
