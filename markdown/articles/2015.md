# google的免费dns服务器

google推出了自己的免费dns服务器，以供公众使用。服务器地址是：

dns1: 8.8.8.8

dns2: 8.8.4.4

我在我的机器上测试了一下：

$ host -a g.cn 8.8.8.8
Trying “g.cn”
Using domain server:
Name: 8.8.8.8
Address: 8.8.8.8#53
Aliases:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 33253
;; flags: qr rd ra; QUERY: 1, ANSWER: 11, AUTHORITY: 0, ADDITIONAL: 0
;; QUESTION SECTION:
;g.cn. IN ANY
;; ANSWER SECTION:
g.cn. 300 IN A 72.14.203.160
g.cn. 259200 IN NS ns3.google.com.
g.cn. 10800 IN MX 10 google.com.s9b2.psmtp.com.
g.cn. 259200 IN NS ns1.google.cn.
g.cn. 259200 IN NS ns2.google.com.
g.cn. 86400 IN SOA ns1.google.com. dns-admin.google.com. 1402219 21600 3600 1209600 300
g.cn. 10800 IN MX 10 google.com.s9b1.psmtp.com.
g.cn. 10800 IN MX 10 google.com.s9a2.psmtp.com.
g.cn. 10800 IN MX 10 google.com.s9a1.psmtp.com.
g.cn. 259200 IN NS ns1.google.com.
g.cn. 259200 IN NS ns4.google.com.

Received 325 bytes from **8.8.8.8#53 in 217 ms**


$ host -a g.cn 8.8.4.4
Trying “g.cn”
Using domain server:
Name: 8.8.4.4
Address: 8.8.4.4#53
Aliases:

;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 40871
;; flags: qr rd ra; QUERY: 1, ANSWER: 11, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;g.cn. IN ANY

;; ANSWER SECTION:
g.cn. 227 IN A 72.14.203.160
g.cn. 259127 IN NS ns3.google.com.
g.cn. 10727 IN MX 10 google.com.s9b2.psmtp.com.
g.cn. 259127 IN NS ns1.google.cn.
g.cn. 259127 IN NS ns2.google.com.
g.cn. 86327 IN SOA ns1.google.com. dns-admin.google.com. 1402219 21600 3600 1209600 300
g.cn. 10727 IN MX 10 google.com.s9b1.psmtp.com.
g.cn. 10727 IN MX 10 google.com.s9a2.psmtp.com.
g.cn. 10727 IN MX 10 google.com.s9a1.psmtp.com.
g.cn. 259127 IN NS ns1.google.com.
g.cn. 259127 IN NS ns4.google.com.

Received 325 bytes from **8.8.4.4#53 in 196 ms**



好记又免费，爽哉！！ :)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
