# Eclipse开发Android应用程序入门

By [Chris Blunt](http://www.smashingmagazine.com/author/chris-blunt/ "Posts by Chris Blunt")

**翻译：赵锟**  

原文出处：<http://www.smashingmagazine.com/2010/10/25/get-started-developing-for-android-with-eclipse/>

如今的移动设备应用程序开发充满着让人振奋的东西。功能强大的硬件支持，平板电脑，多样的软件平台（塞班 OS，iOS，WebOS，Windows Phone 7…)，移动设备开发者前景充满了机会和挑战。

当你想要开始开发你的移动设备程序时，如此多的选择可能让你产生困扰。究竟应该选择神马平台？我应该学习神马语言？为你计划的项目选择神马工具？在本教程中，你将学会如何在Google公司的开源移动设备操作系统Android下开发应用程序。

### 为神马选Android

Android是一个基于Linux内核的开源平台， 并且被安装在来自于不同厂商的上千种设备中。Android将各种移动设备的硬件如 电子罗盘，摄像头，GPS，方向感应，等等暴露给你的应用程序。  

  

Android的免费开发工具可以让你以0成本开始编写你的软件。当你想向世界展示你的应用程序的时候，你可以将你的软件发布到Google的 Android 市场。向Andriod Market 发布程序只一次性的收取注册费用（25元），并且不像苹果的App Store ，对每一次的提交都要做检查，除非你的程序明显地违法，在经过一个快速检查的流程后，才能让你的程序提供给客户下载和购买。

下面是Android对于开发者的优点：

- Android的SDK可以在Windows,Mac和Linux上运行，因此你不需要为了开发环境支付额外的新硬件投入。（译者注：我曾近在Win7 64x + VMWare上成功的安装Mac Snow leopard + XCode的开发环境，对于爱用盗版的人来说，这点MS优势不是很大啊）
- 构建于JAVA上的SDK。如果你熟悉JAVA语言，你就是事半功倍了。（译者注：这个酷壳有篇文章讨论过，大家可以参看：<https://coolshell.cn>）
- 你只要在Android Market上发布应用程序，你将有潜在的成千上万的用户。而且你不一定非要把程序发布在Android Market上，你还可以在你的博客上发布。而且有传言，Amazon已近在最近准备搭建他们自己的Android 应用程序商店了。
- 除了了技术性的[SDK 文档](https://developer.android.com/sdk/index.html)外,还可以找到其他更多的使用者和开发者的资源。

闲话少说——下面让我们进入正题，开始开发我们的Android应用程序。

### 安装Eclipse和Android SDK

Android应用程序的推荐开发环境是带有Android开发包插件(Android Devlopment Toolkit (ADT))的Eclipse。我在这里简要说明一下安装流程。如果你需要更多的细节，Google的[开发人员网页](https://developer.android.com/sdk/)中详尽地解释了具体的安装配置过程

- 为你的平台下载[Android SDK](https://developer.android.com/)（Windows ， Mac OS X 或者 Linux）。
- 在你的硬盘上解压下载文件 (在Linux, 我使用 /opt/local/).
- 如果你没有安装Eclipse，下载并安装[Eclipse JAVA 集成开发环境](http://eclipse.org/downloads/packages/eclipse-ide-java-developers/galileosr2)包。 用于编程的话, Google推荐使用Eclipse 3.5 (Galileo).
- 运行Eclipse 并选择*Help->Install New Software*.
- 在Available Software窗口中点击Add按钮。
- 进入 Android Development Tools 的*Name*输入框, 在Location 输入框输入https://dl-ssl.google.com/android/eclipse/
- 检查可用软件中有Developer Tools并点击OK按钮。这将安装Android Development Tools 和DDMS, Android的调试工具。

![](https://coolshell.cn/uploads/2011/04/install.gif "install")

- 点击Next和Finish按钮以完成安装，安装完成后，你需要重启你的Eclipse一次。
- 在Eclipse重启后，选择Window->Preference 后你可以在分类列表中看到Android这一项了。
- 现在需要告诉Eclipse，你的Android SDK安装在什么地方。点击Android项后浏览选择你解压后的Android SDK所在的路径。例如/opt/local/android-sdk。

![](https://coolshell.cn/uploads/2011/04/eclipse_android_preferences.jpg "eclipse_android_preferences")

- 点击OK按钮，保存信息。

### 选择Android 平台

在你开始编写Android应用程序之前，你需要为你需要开发应用程序的Android设备下载SDK平台。每个平台都有可以安装在用户设备上的不同版本的SDK。对于Android1.5或以上版本，有两个可用的平台： *Android Open Source Project* 和 *Google*.

*Android Open Source Project* 平台是开源的，但是不包括Google公司的私有化扩展，比如Google Map。如果不选择使用Google的API，Google的地图功能就不会在你的应用程序中生效。除非你有特别的原因，否则我们推荐你选择Google平台，因为这样你可享受到Google的扩展类库提供的便利。

- 选择*Window Android SDK and AVD Manager*.
- 点击左栏中的*Available Packages* 并选择选择Respository中有效的Android SDK平台。
- 你可以选择列表中所需要的平台，或全选下载所有有效的平台。当你选择完毕，单击*Install Selected* 并完成安装。

![](https://coolshell.cn/uploads/2011/04/sdk.jpg "sdk")  

一旦成功的下载所有的平台后，你就可以准备开始开发Android应用程序了。

### 创建一个新的Android项目

Eclipse的新建项目向导能为你创建一个新的Android项目，并生成可以开始运行的文件和代码。通过向导生成代码，可以让你马上得到一个Android程序运行的直观映像并为你提供了一个帮助你快速入门的方法：

- 选择 *File->New->Project…*
- 选择*Android Project*
- 在*New Project* 对话框, 键入如下的设置:

[code]  

Project Name: BrewClock  

Build Target: Google Inc. 1.6 (Api Level 4)  

Application Name: BrewClock  

Package Name: com.example.brewclock  

Create Activity: BrewClockActivity  

Min SDK Version: 4  

[/code]

![](https://coolshell.cn/uploads/2011/04/eclipse_new_project_settings.jpg "eclipse_new_project_settings")

在点击了完成按钮之后，Eclipse将为你创建一个新的可以运行的Android项目。注意，你通知了Eclipse生成了一个叫做BrewClockActivity的Activity。这个Activity的代码用于运行你的应用程序。生成的代码将在程序运行时非常简单地显示一条“Hello World”消息。

目录


- [包](#%E5%8C%85 "包")
- [SDK 版本](#SDK_%E7%89%88%E6%9C%AC "SDK 版本")
- [Android 虚拟设备](#Android_%E8%99%9A%E6%8B%9F%E8%AE%BE%E5%A4%87 "Android 虚拟设备")
- [运行代码](#%E8%BF%90%E8%A1%8C%E4%BB%A3%E7%A0%81 "运行代码")
- [设计用户界面](#%E8%AE%BE%E8%AE%A1%E7%94%A8%E6%88%B7%E7%95%8C%E9%9D%A2 "设计用户界面")
- [开发用户界面](#%E5%BC%80%E5%8F%91%E7%94%A8%E6%88%B7%E7%95%8C%E9%9D%A2 "开发用户界面")
- [关联你的布局Layout与代码](#%E5%85%B3%E8%81%94%E4%BD%A0%E7%9A%84%E5%B8%83%E5%B1%80Layout%E4%B8%8E%E4%BB%A3%E7%A0%81 "关联你的布局Layout与代码")
- [资源对象](#%E8%B5%84%E6%BA%90%E5%AF%B9%E8%B1%A1 "资源对象")
- [监听事件](#%E7%9B%91%E5%90%AC%E4%BA%8B%E4%BB%B6 "监听事件")
- [避免在你的代码中硬编码](#%E9%81%BF%E5%85%8D%E5%9C%A8%E4%BD%A0%E7%9A%84%E4%BB%A3%E7%A0%81%E4%B8%AD%E7%A1%AC%E7%BC%96%E7%A0%81 "避免在你的代码中硬编码")
- [运行Brew Clock](#%E8%BF%90%E8%A1%8CBrew_Clock "运行Brew Clock")
#### 包

包名是你的应用程序标示。当你开始准备在Android Market上发布你的应用程序的时候，Android用这个标识符精确地记录你的应用程序的更新过程，因此让包名唯一是非常重要的。尽管我们在这里使用了com.example.brewclock这样的名字空间，对于真实的应用程序，你应该选择类似于com.你的公司名.你的应用程序名 这样的包名。

#### SDK 版本

Min SDK Version 是你的Android程序所能运行得最早版本号。对于每个新发布的Android，SDK会增加并修改一些方法。通过选择一个版本号，Android（Android Market）会知道你的应用程序能运行在等于或晚于指定版本的设备之上。

### 运行你的应用程序

现在让我们开始在Eclipse中运行我们的应用程序。由于是第一次运行，Eclipse将会询问你的项目类型：

- 选择*Run->Run* 或 按下 *Ctrl+F11*.
- 选择*Android Application* 并点击 *OK* 按钮.

Eclipse 将会在一个Android设备上运行一个应用程序。在这个时候，由于你没有任何Android设备，因此在运行时一定会返回一个失败，并且询问你是否要新建一个Android的虚拟设备。（AVD）  

![](https://coolshell.cn/uploads/2011/04/eclipse_no_avd.jpg "eclipse_no_avd")

#### Android 虚拟设备

Android 虚拟设备 (AVD) 是一个模拟真实世界中Android设备的模拟器，例如移动电话或平板电脑。你可以在不买任何真实Android设备情况下，使用AVD测试你的应用。

你可以创建任意多个你喜欢的AVD，每个可以建立在不同版本的Android平台之上。对于你创建的每个Android设备，你可以配置不同的硬件属性，比如是否具有物理键盘，是否支持GPS，摄像头的像素，等等。

在你开始运行你的应用程序之前，你需要创建你的AVD，来运行指定的SDK平台（Google APIs 1.6）。

现在让我开始:

- 如果还没有开始运行你的应用程序，点击run（或按下 *Ctrl+F11*）。
- 当目标设备弹出警告，点击*Yes* 以创建新的AVD。
- 单击*Android SDK and AVD Manager* 对话框内的*New* 按钮.
- 为你的AVD键入如下的设置：

[code]  

Name: Android\_1.6  

Target: Google APIs (Google Inc.) – API Level 4  

SD Card Size: 16 MiB  

Skin Built In: Default (HVGA)  

[/code]

- 单击 *Create AVD* 让Android为你创建一个新虚拟设备。
- 关闭the *Android SDK and AVD Manager* 对话框.

![](https://coolshell.cn/uploads/2011/04/sdk_manager_new_avd.jpg "sdk_manager_new_avd")

#### 运行代码

再次运行你的应用程序（*Ctrl+F11*）。 Eclipse 将build 你的项目并运行一个新的AVD。记住，AVD模拟了一个完全的Android系统，因此你需要有耐心来等待这个缓慢的启动过程，就如同你重启真实的Android设备一样。一个好的做法是不要关闭你的AVD，直到你完成了你一天的工作。  

当你的模拟器启动后，Eclipse自动地安装并运行你的应用程序。

![](https://coolshell.cn/uploads/2011/04/app_running-550-e1287474474253.jpg "app_running-550-e1287474474253")

### 开发你第一个Android应用

生成的代码能良好的运行，但是你真正想要的是开发一个真实的应用程序。为此，我们首先果一个咸蛋的设计流程，并开始创建一个可以让你部署在Android设备上的应用。

大部分的开发者（包括我自己）都喜欢每天一杯咖啡或茶。在下一节中，你将开发一个简单的泡茶计数器应用程序来记录用户泡了多少杯茶，并为泡每杯茶做一个定时器。

你可以从[GitHub](https://github.com/cblunt/brewclock)下载整个教程的源代码.

#### 设计用户界面

在开发任何Android应用程序之前的第一步就是设计和开发用户界面。下面是一个我们这个应用程序的用户界面的一个概览。

![](https://coolshell.cn/uploads/2011/04/design_sketch.jpg "design_sketch")

用户将能通过+和-按钮设置一个泡茶的定时器。当单击开始按钮，定时器将开始按指定的时间递减。除非用户再次点击按钮以取消计时，否则当定时器为0的时候，累计的泡茶计数brew将增加1。

#### 开发用户界面

Android 用户界面或布局*layouts*, 是通过XML文档来描述的，可以在项目的res/layouts目录下找到。在之前运行在模拟器上代码中，我们可以看到由eclipse自动生成的布局代码在res/layouts/main.xml 中。

Eclipse有一个图形化的布局设计器，通过在屏幕上的拖拽控制来完成布局的设计，然而，我却发现直接写XML并使用图形布局来预览是更容易的方式。

现在让我们对main.xml做一些工作以达到上图的效果：

- 在Eclipse中通过双击PackageExplorer的res/layouts/main.xml 来打开xml。
- 点击屏幕下方main.xml 来切换为xml视图。

将main.xml中内容改为如下的内容：

[code]  

# /res/layouts/main.xml  

<?xml version="1.0" encoding="utf-8"?>  

<LinearLayout  

xmlns:android="http://schemas.android.com/apk/res/android"  

android:orientation="vertical"  

android:layout\_width="fill\_parent"  

android:layout\_height="fill\_parent">  

<LinearLayout  

android:orientation="horizontal"  

android:layout\_width="fill\_parent"  

android:layout\_height="wrap\_content"  

android:padding="10dip">  

<TextView  

android:layout\_width="wrap\_content"  

android:layout\_height="wrap\_content"  

android:textSize="20dip"  

android:text="Brews: " />  

<TextView  

android:layout\_width="fill\_parent"  

android:layout\_height="wrap\_content"  

android:text="None"  

android:gravity="right"  

android:textSize="20dip"  

android:id="@+id/brew\_count\_label" />  

</LinearLayout>  

<LinearLayout  

android:orientation="horizontal"  

android:layout\_width="fill\_parent"  

android:layout\_height="wrap\_content"  

android:layout\_weight="1"  

android:gravity="center"  

android:padding="10dip">  

<Button  

android:id="@+id/brew\_time\_down"  

android:layout\_width="wrap\_content"  

android:layout\_height="wrap\_content"  

android:text="-"  

android:textSize="40dip" />  

<TextView  

android:id="@+id/brew\_time"  

android:layout\_width="wrap\_content"  

android:layout\_height="wrap\_content"  

android:text="0:00"  

android:textSize="40dip"  

android:padding="10dip" />  

<Button  

android:id="@+id/brew\_time\_up"  

android:layout\_width="wrap\_content"  

android:layout\_height="wrap\_content"  

android:text="+"  

android:textSize="40dip" />  

</LinearLayout>  

<Button  

android:id="@+id/brew\_start"  

android:layout\_width="fill\_parent"  

android:layout\_height="wrap\_content"  

android:layout\_gravity="bottom"  

android:text="Start" />  

</LinearLayout>

[/code]

正如你所见的，Android的XML布局文件是繁琐的，但却能让你控制到屏幕的各个元素。

在Android中最重要的接口元素是布局Layout容器，例如例子中使用的LinearLayout 。这些元素对于用户是不可见的,但是却扮演者例如Buttons 和TextViews这些元素的布局容器。

Android中有几种不同类型的布局视图layout view，每一种都用于开发不同的布局。如同LinearLayout 和AbsoluteLayout ，TableLayout 可以让你使用更为复杂的基于表格结构的布局。你可以在SDK的API文档的[通用布局对象](https://developer.android.com/guide/topics/ui/layout-objects.html)中查找到更多的布局。

#### 关联你的布局Layout与代码

保存你的布局，在Eclipse中点击*Run*图标或按下*Ctrl+F11*重新在模拟器中运行你的程序。你现看到不是之前出现的Hello World消息了，你将看到Android显示了一个新的界面。

如果点击界面上的任何按钮，他们将期望的显示为高亮，但是不会执行任何操作。现在让我们在布局修改后改进一下我们的源码：

# /src/com/example/brewclock/BrewClockActivity.java

```java
...
import android.widget.Button;
import android.widget.TextView;

public class BrewClockActivity extends Activity {
  /** Properties **/
  protected Button brewAddTime;
  protected Button brewDecreaseTime;
  protected Button startBrew;
  protected TextView brewCountLabel;
  protected TextView brewTimeLabel;

  ...
 }

```

下一步,我们将修改调用onCreate。当Android启动你的应用程序的时候，Android会首先调用这个方法。 在Eclipse生成的代码中，onCreate把activity的视图设置成R.layout.main。这行代码告诉Android解释我们的布局配置XML文件，并显示它。

#### 资源对象

在Android中，R是一个自动生成的对象，这是一个特殊的对象，你可以在代码中通过这个对象访问项目中的资源（布局，字符串，菜单，图标，…） 。每个资源都有一个给定的id。在上面的那个布局文件中，有一些@+id XML 属性。我们将通过这些值来关联布局中的Buttons 与TextViews和我们的代码和：

# /src/com/example/brewclock/BrewClockActivity.java

```java
...
public class BrewClockActivity extends Activity {
  ...
  public void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.main);

    // Connect interface elements to properties
    brewAddTime = (Button) findViewById(R.id.brew_time_up);
    brewDecreaseTime = (Button) findViewById(R.id.brew_time_down);
    startBrew = (Button) findViewById(R.id.brew_start);
    brewCountLabel = (TextView) findViewById(R.id.brew_count_label);
    brewTimeLabel = (TextView) findViewById(R.id.brew_time);
  }
}

```
#### 监听事件

为了检测到用户单击我们的按钮，我们需要实现一个监听器listener。你可能会从其他的事件驱动系统中熟悉监听器或回调函数*callbacks*。比如Javascript/JQuery事件或Rails的回调函数。

Android通过Listener接口提供相似的机制，例如OnClickListener，这个接口中定义了那些会被事件触发的方法。当用户点击屏幕的时候，实现OnClickListener 接口将会通知你的应用程序，并告诉他们所按得屏幕按钮。你当然也需要告诉每个button的ClickListener，以便Android知道具体通知到那个监听器：

# /src/com/example/brewclock/BrewClockActivity.java

```java
...
// Be sure not to import
// `android.content.dialoginterface.OnClickListener`.
import android.view.View.OnClickListener;

public class BrewClockActivity extends Activity
  implements OnClickListener {
  ...
  public void onCreate(Bundle savedInstanceState) {
    ...
    // Setup ClickListeners
    brewAddTime.setOnClickListener(this);
    brewDecreaseTime.setOnClickListener(this);
    startBrew.setOnClickListener(this);
  }
  ...
  public void onClick(View v) {
    // TODO: Add code to handle button taps
  }
}

```

下一步，我们将增加每个按钮按下的处理过程。我们将为Activity类增加4个属性，这些属性将用来让用户设置和记录我们泡茶时间，泡茶计数，计时器是否在运行的标志。

# /src/com/example/brewclock/BrewClockActivity.java

```java
...
public class BrewClockActivity extends Activity
  implements OnClickListener {
  ...
  protected int brewTime = 3;
  protected CountDownTimer brewCountDownTimer;
  protected int brewCount = 0;
  protected boolean isBrewing = false;
  ...
  public void onClick(View v) {
    if(v == brewAddTime)
      setBrewTime(brewTime + 1);
    else if(v == brewDecreaseTime)
      setBrewTime(brewTime -1);
    else if(v == startBrew) {
      if(isBrewing)
        stopBrew();
      else
        startBrew();
    }
  }
}

```

注意我们使用了Android提供的类CountDownTimer 。这让我们非常容易的创建和开始一个简单的递减计数，这个递减计数在递减运行的时候，每当执行一个递减就发出一个通知。你将在下面的startBrew 方法中使用到这个计数器。

在下面的方法是所有处理逻辑，这些处理逻辑用于处理设置泡茶时间，开始停止计数和维护计数器。我们同样地在onCreate方法中来初始化我们的 brewTime和 brewCount变量。

将这些代码放入到不同的类中是一种好做法。但是为了简洁，我把我们所有的代码都放到了BrewClockActivity中：

# /src/com/example/brewclock/BrewClockActivity.java

```java
...
public class BrewClockActivity extends Activity
  implements OnClickListener {
  ...
  public void onCreate(Bundle savedInstanceState) {
    ...
    // Set the initial brew values
    setBrewCount(0);
    setBrewTime(3);
  }

  /**
   * Set an absolute value for the number of minutes to brew.
   * Has no effect if a brew is currently running.
   * @param minutes The number of minutes to brew.
   */
  public void setBrewTime(int minutes) {
    if(isBrewing)
      return;

    brewTime = minutes;

    if(brewTime < 1)
      brewTime = 1;

    brewTimeLabel.setText(String.valueOf(brewTime) + "m");
  }

  /**
   * Set the number of brews that have been made, and update
   * the interface.
   * @param count The new number of brews
   */
  public void setBrewCount(int count) {
    brewCount = count;
    brewCountLabel.setText(String.valueOf(brewCount));
  }

  /**
   * Start the brew timer
   */
  public void startBrew() {
    // Create a new CountDownTimer to track the brew time
    brewCountDownTimer = new CountDownTimer(brewTime * 60 * 1000, 1000) {
      @Override
      public void onTick(long millisUntilFinished) {
        brewTimeLabel.setText(String.valueOf(millisUntilFinished / 1000) + "s");
      }

      @Override
      public void onFinish() {
        isBrewing = false;
        setBrewCount(brewCount + 1);

        brewTimeLabel.setText("Brew Up!");
        startBrew.setText("Start");
      }
    };

    brewCountDownTimer.start();
    startBrew.setText("Stop");
    isBrewing = true;
  }

  /**
   * Stop the brew timer
   */
  public void stopBrew() {
    if(brewCountDownTimer != null)
      brewCountDownTimer.cancel();

    isBrewing = false;
    startBrew.setText("Start");
  }
  ...
}

```

这段代码唯一和Android相关的就是使用setText方法来设置文本的显示文字。在startBrew方法中，我们创建，并开始了一个CountDownTimer来开每秒递减计数直到计数器为0。注意，我们定义了CountDownTimer以内联方式监听onTick 和 onFinish方法。 onTick 方法将每1000毫秒（1秒）执行一次，并递减, 当计数器为0的时候，onFinish方法被调用。

#### 避免在你的代码中硬编码

为了使教程代码简单，我故意地在程序中将控件的标号直接写到字串中（例如： “Brew Up!”, “Start”, “Stop”） 通常，这不是一个好的做法，因为如果在大型项目中，这样做会使得修改变得麻烦。

Android 提供了一种简洁的方法让你使用R对象来使字符串和代码分离。R 让你在xml文件（res/values/strings.xml）定义所有你程序中字符串，并让你可以在代码中应用到这些字符串。例如：

# /res/values/strings.xml

[code]  

<string name="brew\_up\_label">Brew Up!</string>  

…  

[/code]

# /res/com/example/brewclock/BrewClockActivity.java

[code]  

…  

brewLabel.setText(R.string.brew\_up\_label);  

…  

[/code]

现在，如果你想改变Brew Up! 字样，你只要一次性的修改strings.xml文件就行了。你的应用将生成一堆代码来保证你程序中所有使用到这些字符串的地方都能被生效！

#### 运行Brew Clock

代码完成之后，现在是试运行程序的时候了。单击*Run* 或 *Ctrl+F11* 在模拟器中启动我们的应用. 所有都运行良好，你将会看到你创建的用户界面在准备时间一到就可以喝你所泡的茶了！试着设置不同的时间，并点击*Start* 观看倒计时。

![](https://coolshell.cn/uploads/2011/04/app_finished-550-e1287474491689.jpg "app_finished-550-e1287474491689")

### 总结

在这个关于Android的简单介绍中，你已学会如何安装Android SDK和Eclipse的Android 开发工具插件（ADT）。你也学会如何创建一个模拟设备，并通过这个设备来测试你的应用程序。你还学会了如何开发Android应用程序。上面了那些作为标题的关键概念在以后你自己开发Android应用程序的时候将会经常用到。

我们希望，这个教程能激发你的开发移动应用程序的欲望，并步入这个令人激动的领域。Android为当前和即将到来的移动设备应用程序开发提供了一条宽广的道路。如果你已经开发你自己的移动应用，请在评论中告诉我们。

*(ik), (vf)*

*（全文完）*

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
