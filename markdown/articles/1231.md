# Javascripts加密库

一般说来，使用HTTP协议是不加密的，所有的数据都是以纯文本方式提交的，就算是你提交数据时，也是使用纯文本的方式发送。只有HTTPS协议会有SSL加密数据，但一般来说，HTTPS需要服务器端进行SSL设置，并有些麻烦。而jCryption这个jQuery插件能够加密由Forms提交的POST/GET数据。jCryption使用RSA公钥密码算法加密，另外，该项目还提供一个PHP文件来处理数据的解密。

[![](http://www.webresourcesdepot.com/wp-content/uploads/image/javascript-encyrption.jpg "JCryption")](http://www.jcryption.org/)

这个库是一个开源库，也是一个同时使用MIT和GPL协议的项目。

你需要注意的是，这个库无法取代SSL，使用这个库，你依然可能受到[MITM攻击](https://en.wikipedia.org/wiki/Man-in-the-middle_attack)（中间人攻击 Man-in-the-middle-attacks）

主页：<http://www.jcryption.org/>  

下载：[http://code.google.com/p/jcryption/downloads/list](https://code.google.com/p/jcryption/downloads/list)  

示例：<http://www.jcryption.org/demo/>

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
