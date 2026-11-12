# JavaMail使用

（**本文由网友jjzhx\_1211投递，感谢!**）

使用JavaMail需要两个包：activation-1.1.jar和mail-1.4.2.jar（当然现在最新的版本已经不止了），也可以直接包含Java SE 6的j2ee.jar，自带了前面的两个包。我把邮件功能写成了一个服务，发送邮件的数据都通过Map<String, String>类型的参数封装了起来。**代码见文章最后**。

目录


- [Session](#Session "Session")
- [Message](#Message "Message")
- [Address](#Address "Address")
- [Authenticator](#Authenticator "Authenticator")
- [Transport](#Transport "Transport")
- [Store and Folder](#Store_and_Folder "Store and Folder")
- [代码](#%E4%BB%A3%E7%A0%81 "代码")
#### Session

Session 定义了一个基本的邮件会话，任何工作都是基于这个Session的。Session 对象需要一个 java.util.Properties 对象来得到类似 邮件服务器，用户名，密码这样的信息。Session 的构造函数是私有的，可以通过 getDefaultInstance() 方法来取得一个单一的可以被共享的默认session 如：

```java
Properties props = new Properties();
Session session = Session.getDefaultInstance(props,null);
```

或者，可以使用 getInstance() 方法来创建一个唯一的 session如：

```java
Properties props = new Properties();
Session session = Session.getInstance(props,null);
```

在这两种方法中 其中的 null 参数是一个 Authenticator 对象，在这里没有被使用的，所以就是null。在大多数案例中，使用一个共享session 已经做够了。

#### Message

一旦你创建了Session对象，那么下面要做的就是创建message来发送。Message是一个抽象类，在大部分应用中你可以使用它的子类javax.mail.internet.MimeMessage 。MimeMessage 是一个理解在不同RFCs中定义的MIME类型以及headers的e-mail message。Message headers 必须使用 US-ASCII 字符集。可以用如下的方法创建一个Message

`MimeMessage message = new MimeMessage(session);`

我们注意到，这里需要用session对象作为构造函数的参数。当然，还有其它的构造函数，比如从用RFC822格式化过的输入流来创建message。

一旦你得到了 message ,你就可以来设置它的各个部分（parts）。设置内容（content）的基本的机制是使用setContent() 方法。

`message.setContent("Email Content. ","text/plain");`

如果，你能够明确你的使用MimeMessage来创建message 并且只是使用普通的文本（plain text） 那么你也可以使用 setText() 方法，setTest()方法只需要设置具体的内容，它默认的MIME类型是 text/plain

`message.setText("Email Content. ");`

对于普通文本类型的邮件，有一种机制是首选（ message.setText(“Email Content. “)）的设置内容的方法。如果要创建其它类型的message ，比如　HTML类型的message   那么还是需要使用前者　（　message.setContent(“Email Content. “,”text/html”);　）  

设置主题（subject ），使用setSubject() 方法

`message.setSubject(" Subject ");`

#### Address

当你已经创建Session 以及 Message，并且已经为message 填充了内容，那么接下来要做的就是给你的邮件添加一个地址（Address）。　就像Message一样，Address也是一个抽象类，我们可以使用它的一个子

javax.mail.internet.InternetAddress

创建一个地址非常简单

`Address address = new InternetAddress("<a href="mailto:su[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)">[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)</a>");`

如果，你希望在出现邮件地址的地方出现一个名称，那么你只需要再多传递一个参数。

`Address address = new InternetAddress("<a href="mailto:[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)&quot;,&quot;Steve">[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)","Steve</a>");`

你需要为 message 的from以及 to 字段创建address对象。为了识别发送者，你需要使用setFrom() 和 setReplyTo() 方法。

`messge.setFrom(address);`

如果你的message 需要显示多个 from 地址，可以使用 addFrom() 方法

```java
Address address[] = {....};
message.addFrom(address);
```

为了辨识message 的收件人，你需要使用 setRecipient() 方法。这个方法除了address参数之外，还需要一

Message.RecipientType 。  

message.addRecipient(type,address);  

Message.RecipientType有几个预先定义好的类型  

Message.RecipientType.TO　　收件人  

Message.RecipientType.CC　　抄送  

Message.RecipientType.BCC　 暗送

如果你的一封邮件，需要发送给你的老师，并还要给你的几个同学，那么你可以这样

```java
Address toAddress = new InternetAddress("<a href="mailto:[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)">[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)</a>");
Address[] ccAddress = {new InternetAddress("<a href="mailto:[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)&quot;),new">[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)"),new</a> InternetAddress("<a href="mailto:[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)">[[email protected]](https://coolshell.cn/cdn-cgi/l/email-protection)</a>")};
message.addRecipient(Message.RecipientType.To, toAddress);
message.addRecipient(Message.RecipientType.CC, ccAddress);
```

JavaMail 没有提供电子邮件地址有效性的检测。这些超越了JavaMail API的范围。

#### Authenticator

通过Authenticator设置用户名、密码，来访问受保护的资源，这里的资源一般指的是邮件服务器。

Authenticator也是一个抽象类，你需要自己编写子类已备应用。你需要实现getPasswordAuthentication()方法，并返回一个PasswordAuthentication实例。你必须在 session被创建时， 注册你的 Authenticator。这样，当需要进行认证是，你的Authenticator就可以被得到。

```java
Properties props = new Properties();
//设置属性
Authenticator auth = new YourAuthenticator();
Session session = Session.getDefaultInstance(props, auth);
```
#### Transport

发送消息最后的一步就是使用Transport类，你可以通过两种方法来进行发送。  

Transport 是一个抽象类，你可以调用它静态的send() 方法来发送

`Transport.send(message);`

或者，你可以为你使用的协议从session中取得一个指定的实例，

```java
Transport transport = session.getTransport("smtp");
transport.sendMessage(message, message.getAllRecipients());
transport.close();
```
#### Store and Folder

这两个类重要用于取得信息。在创建了Session之后，需要连接到一个 Store ，你需要告诉Store你使用的是什么协议。

```java
// Store store = session.getStore("imap");
Store store = session.getStore("pop3");
store.connect(host, username, password);
```

在连接到一个 Store 后，你可以得到一个 Folder，当然，这个Floder必须是打开的。

```java
Folder folder = store.getFolder("INBOX");
folder.open(Folder.READ_ONLY);
Message message[] = folder.getMessages();
```

如果使用POP3那么，INDEX是唯一可用的文件夹。如果使用的是IMAP，你就可以使用其它的文件夹。

#### 代码

```java
public boolean sendEmail(Map<String, String> data) {
    // 创建Properties 对象
    Properties props = System.getProperties();
    props.put("mail.smtp.host", Constants.HOST); // 全局变量
    props.put("mail.smtp.auth", "true");

    // 创建邮件会话
    Session session = Session.getDefaultInstance(props,
    new Authenticator() { // 验账账户
        @Override
        public PasswordAuthentication getPasswordAuthentication() {
            return new PasswordAuthentication(Constants.USERNAME,
                                              Constants.PASSWORD);
        }
    });

    try {
        // 定义邮件信息
        MimeMessage message = new MimeMessage(session);
        message.setFrom(new InternetAddress(Constants.FROM));
        message.addRecipient(
            Message.RecipientType.TO,
            new InternetAddress(
                // 这里可以添加多个目的用户
                data.get(Constants.EMAIL_TO)
            )
        );
        // 添加邮件发送时间（不知道体现在哪儿）
        message.setSentDate(new Date());
        // 要编码，否则中文会出乱码，貌似这个方法是对数据进行了
        //("=?GB2312?B?"+enc.encode(subject.getBytes())+"?=")形势的包装
        message.setSubject(MimeUtility.encodeText(data.get(Constants.EMAIL_SUBJECT), "gbk", "B"));

        MimeMultipart mmp = new MimeMultipart();
        MimeBodyPart mbp_text = new MimeBodyPart();
        // "text/plain"是文本型，没有样式，
        //"text/html"是html样式，可以解析html标签
        mbp_text.setContent(data.get(Constants.EMAIL_TEXT),
                            "text/html;charset=gbk");
        mmp.addBodyPart(mbp_text); // 加入邮件正文

        // 处理附件，可以添加多个附件
        if (data.get(Constants.EMAIL_ATTACHMENT) != null) {
            String[] files = data.get(Constants.EMAIL_ATTACHMENT).split(",");
            if (files.length != 0) {
                for (String file : files) {
                    MimeBodyPart mbp_file = new MimeBodyPart();
                    FileDataSource fds = new FileDataSource(file);
                    mbp_file.setDataHandler(new DataHandler(fds));
                    mbp_file.setFileName(MimeUtility.encodeText(fds.getName(), "gbk", "B"));
                    mmp.addBodyPart(mbp_file);
                }
            }
        }
        message.setContent(mmp);
        // message.setText(data.get(Constants.EMAIL_TEXT));

        // 发送消息
        // session.getTransport("smtp").send(message); //也可以这样创建Transport对象
        Transport.send(message);
        return true;
    } catch (Exception e) {
        e.printStackTrace();
        return false;
    }
}

```

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
