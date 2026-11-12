# 16个简单实用的.htaccess小贴示

.htaccess 文件 (Hypertext Access file) 是Apache Web服务器的一个非常强大的配置文件，对于这个文件，Apache有一堆参数可以让你配置出几乎随心所欲的功能。.htaccess 配置文件坚持了Unix的一个文化——使用一个ASCII 的纯文本文件来配置你的网站的访问策略。

这篇文章包括了16个非常有用的小技巧。另外，因为.htaccess 是一个相当强大的配置文件，所以，一个轻微的语法错误会造成你整个网站的故障，所以，在你修改或是替换原有的文件时，一定要备份旧的文件，以便出现问题的时候可以方便的恢复。

**1. 使用.htaccess 创建自定义的出错页面。**对于Linux Apache来说这是一项极其简单的事情。使用下面的.htaccess语法你可以轻松的完成这一功能。（把.htaccess放在你的网站根目录下）

ErrorDocument 401 /error/401.php  

ErrorDocument 403 /error/403.php  

ErrorDocument 404 /error/404.php  

ErrorDocument 500 /error/500.php

**2. 设置网站的时区**

SetEnv TZ America/Houston

**3. 阻止IP列表**  

有些时候，你需要以IP地址的方式阻止一些访问。无论是对于一个IP地址还是一个网段，这都是一件非常简单的事情，如下所示：

allow from all  

deny from 145.186.14.122  

deny from 124.15

Apache对于被拒绝的IP会返回403错误。

**4. 把一些老的链接转到新的链接上——搜索引擎优化SEO** 

Redirect 301 /d/file.html <http://www.htaccesselite.com/r/file.html>

**5. 为服务器管理员设置电子邮件。**

ServerSignature EMail  

SetEnv SERVER\_ADMIN [[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection#2b4f4e4d4a5e475f6b4f44464a424505484446)

**6. 使用.htaccess 访止盗链。**如果你网站上的一个图片被别的N多的网站引用了，那么，这很有可能会导致你服务器的性能下降，使用下面的代码可以保护某些热门的链接不被过多的引用。

Options +FollowSymlinks  

# Protect Hotlinking  

RewriteEngine On  

RewriteCond %{HTTP\_REFERER} !^$  

RewriteCond %{HTTP\_REFERER} !^http://([www.)?domainname.com/](http://www.%29/?domainname.com/) [nc]  

RewriteRule .\*.(gif|jpg|png)$ <http://domainname.com/img/hotlink_f_o.png> [nc]

**7. 阻止 User Agent 的所有请求**

## .htaccess Code :: BEGIN  

## Block Bad Bots by user-Agent  

SetEnvIfNoCase user-Agent ^FrontPage [NC,OR]  

SetEnvIfNoCase user-Agent ^Java.\* [NC,OR]  

SetEnvIfNoCase user-Agent ^Microsoft.URL [NC,OR]  

SetEnvIfNoCase user-Agent ^MSFrontPage [NC,OR]  

SetEnvIfNoCase user-Agent ^Offline.Explorer [NC,OR]  

SetEnvIfNoCase user-Agent ^[Ww]eb[Bb]andit [NC,OR]  

SetEnvIfNoCase user-Agent ^Zeus [NC]

Order Allow,Deny  

Allow from all  

Deny from env=bad\_bot

## .htaccess Code :: END

**8. 把某些特殊的IP地址的请求重定向到别的站点**

ErrorDocument 403 [http://www.youdomain.com](http://www.youdomain.com/)  

Order deny,allow  

Deny from all  

Allow from 124.34.48.165  

Allow from 102.54.68.123

**9. 直接找开文件而不是下载** – 通常，我们打开网上文件的时候总是会出现一个对话框问我们是下载还是直接打开，使用下面的设置就不会出现这个问题了，直接打开。

AddType application/octet-stream .pdf  

AddType application/octet-stream .zip  

AddType application/octet-stream .mov

**10. 修改文件类型** – 下面的示例可以让任何的文件都成为PHP那么被服务器解释。比如：myphp, cgi，phtml等。

ForceType application/x-httpd-php  

SetHandler application/x-httpd-php

**11. 阻止存取.htaccess 文件**

# secure htaccess file

 order allow,deny  

 deny from all  

**12. 保护服务器上的文件被存取**

# prevent access of a certain file order allow,deny  

 deny from all  

**13. 阻止目录浏览**

# disable directory browsing  

Options All -Indexes

**14. 设置默认主页**

# serve alternate default index page  

DirectoryIndex about.html

**15. 口令认证** – 你可以创建一个文件用于认证。下面是一个示例：

# to protect a file

AuthType Basic  

AuthName “Prompt”  

AuthUserFile /home/path/.htpasswd  

Require valid-user

# password-protect a directory  

resides  

AuthType basic  

AuthName “This directory is protected”  

AuthUserFile /home/path/.htpasswd  

AuthGroupFile /dev/null  

Require valid-user

**16. 把老的域名转像新的域名**

# redirect from old domain to new domain  

RewriteEngine On  

RewriteRule ^(.\*)$ <http://www.yourdomain.com/$1> [R=301,L]

文章：[来源](http://rafeekphp.wordpress.com/2009/06/06/16-great-htaccess-tricks-and-hacks/)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
