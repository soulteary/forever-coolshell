# Sony PS3 Root Key 被破解

著名的黑客George “GeoHot” Hotz（其也帮助破解了iPhone）宣称破解了Sony P3的root key（也称front door key），并将这个key公布于 <http://www.geohot.com/> （墙）。不但发布了root key，还做了一个hello world。Youtube上也有一个相关的视频：[http://www.youtube.com/watch?v=UkLSXsCKDkg](https://www.youtube.com/watch?v=UkLSXsCKDkg)

```c
erk: C0 CE FE 84 C2 27 F7 5B D0 7A 7E B8 46 50 9F 93 B2 38 E7 70 DA CB 9F F4 A3 88 F8 12 48 2B E2 1B
riv: 47 EE 74 54 E4 77 4C C9 B8 96 0C 7B 59 F4 C1 4D
pub: C2 D4 AA F3 19 35 50 19 AF 99 D4 4E 2B 58 CA 29 25 2C 89 12 3D 11 D6 21 8F 40 B1 38 CA B2 9B 71 01 F3 AE B7 2A 97 50 19
 R: 80 6E 07 8F A1 52 97 90 CE 1A AE 02 BA DD 6F AA A6 AF 74 17
 n: E1 3A 7E BC 3A CC EB 1C B5 6C C8 60 FC AB DB 6A 04 8C 55 E1
 K: BA 90 55 91 68 61 B9 77 ED CB ED 92 00 50 92 F6 6C 7A 3D 8D
 Da: C5 B2 BF A1 A4 13 DD 16 F2 6D 31 C0 F2 ED 47 20 DC FB 06 70
```

之所以叫“front door key”，其是相对于“back door” 而言，传统的破解一般是通过软件的某个 bug或是后门来破解。而这次的PS3走的是前门，这就是说——这已经不是破解了，这是完全意义上的PS3正版了。

为什么呢。这和PS3的开发有关。其很像Symbian 的Sign，也就是说，游戏开发商要想让他们的游戏在PS3上发布，其需要把游戏通过法律流程交给Sony，然后被Sign上一个key，就可以成为正式的发行版并可在所有用户的PS3上运行了。所以，这个key是PS3到今天没有盗版游戏的关键。不过随着这个key被找到，这意味着任何人都可以在PS3上发布软件了。

最要命的是，这个Key和PS3的硬件绑定，也就是说，**如果Sony要阻止这个事的话，无法通过升级firmware完成，必需更换硬件！！**

目前，SONY正式对PS3的Root Key被公布导致可以进行自制系统开发的问题[进行回应](http://www.next-gen.biz/news/sony-responds-to-ps3-hacks)。SONY表示目前正在进行相关调查，问题会通过网络更新进行解决，具体情况涉及信息安全问题不便透露。

不过之前黑客集团表示除非Sony出新硬件否则无法修正这一情况，Sony应该会接受这一事实。

而最新的[PS3 Custom Firmware Creator](http://www.ps3-hacks.com/2011/01/04/ps3-custom-firmware-creator-released-permanently-add-install-pkgs-to-the-xmb/)应该是把PS3送上断头台了。而且已经证实，3.55的玩友可以安装3.55的CFW自制系统，安裝之后可以正常运行正版游戏，可以通过选项菜单中多出的pkg安裝功能，安裝u盘里通过电脑下载的游戏更新补丁或是游戏的试玩版。**现在的PS3就像一个PC机，等待着各种不受Sony控制的软件的到来……**

（另：Freebsd[宣布](http://lists.freebsd.org/pipermail/freebsd-current/2011-January/022104.html)支持索尼的游戏机PS3，支持的型号是索尼Playstation 3 Fat版，固件版本号< 3.21 （最新的固件版本是[3.55版](https://us.playstation.com/support/systemupdates/ps3/index.htm)），必须能网络启动。不过，因为黑客已经破译了root key，并允许创作自制固件，因此未来Freebsd或能支持所有版本的PS3。）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
