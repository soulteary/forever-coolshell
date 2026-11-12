# 关于Amazon云宕机的网贴收集

最近，互联网上最大的事可能是Amazon的AWS宕机了，而且好几天都没有完全恢复。整个Internet都在讨论这个事，Internet很不高兴，后果可能很严重。可能是因为这个事件对中国没有影响，所以中文这边相关的文章不多，大家可以参考一下和讯网的这篇《[伤不起！亚马逊史前最大宕机事件的启示](http://tech.hexun.com/2011-04-24/128998619.html)》。

国外有人把所有和这个事件相关的贴子都收集了起来，都是一些相当不错的贴子和文章，尤其是一些经验教训的贴子，很受教，转给大家看看。这个贴子的[来源在这里](http://highscalability.com/blog/2011/4/25/the-big-list-of-articles-on-the-amazon-outage.html)。

目录


- [个别公司的经历，有好有坏](#%E4%B8%AA%E5%88%AB%E5%85%AC%E5%8F%B8%E7%9A%84%E7%BB%8F%E5%8E%86%EF%BC%8C%E6%9C%89%E5%A5%BD%E6%9C%89%E5%9D%8F "个别公司的经历，有好有坏")
- [Amazon Web Services 讨论区](#Amazon_Web_Services_%E8%AE%A8%E8%AE%BA%E5%8C%BA "Amazon Web Services 讨论区")
- [总结](#%E6%80%BB%E7%BB%93 "总结")
- [立场：这是用户的错](#%E7%AB%8B%E5%9C%BA%EF%BC%9A%E8%BF%99%E6%98%AF%E7%94%A8%E6%88%B7%E7%9A%84%E9%94%99 "立场：这是用户的错")
- [立场：这是Amazon的错](#%E7%AB%8B%E5%9C%BA%EF%BC%9A%E8%BF%99%E6%98%AFAmazon%E7%9A%84%E9%94%99 "立场：这是Amazon的错")
- [教训和启示](#%E6%95%99%E8%AE%AD%E5%92%8C%E5%90%AF%E7%A4%BA "教训和启示")
- [Vendor很生气](#Vendor%E5%BE%88%E7%94%9F%E6%B0%94 "Vendor很生气")
#### 个别公司的经历，有好有坏

- [How Heroku Survived the Amazon Outage](https://status.heroku.com/incident/151) on the Heroku status page
- [How SimpleGeo Stayed Up During the AWS Downtime](http://developers.simplegeo.com/blog/2011/04/26/how-simplegeo-stayed-up/) by Mike Malone
- [How SmugMug survived the Amazonpocalypse](http://don.blogs.smugmug.com/2011/04/24/how-smugmug-survived-the-amazonpocalypse) by Don MacAskill  ([Hacker News](http://news.ycombinator.com/item?id=2480763) discussion)
- [How Bizo survived the Great AWS Outage of 2011 relatively unscathed…](http://dev.bizo.com/2011/04/how-bizo-survived-great-aws-outage-of.html) by Someone at Bizo
- [Joe Stump’s explanation](http://www.focus.com/questions/information-technology/amazon-ec2-has-gone-down--what-would-prefered-hosting-be/#comment43192) of how SimpleGeo survived
- [How Netflix Survived the Outage](http://www.slideshare.net/adrianco/netflix-in-the-cloud-2011)
- [Why Twilio Wasn’t Affected by Today’s AWS Issues](http://www.twilio.com/engineering/2011/04/22/why-twilio-wasnt-affected-by-todays-aws-issues/) on Twilio Engineering’s Blog ([Hacker News](http://news.ycombinator.com/item?id=2472999) thread)
- [On reddit’s outage](http://www.reddit.com/r/announcements/comments/gva4t/on_reddits_outage/#)
- [What caused the Quora problems/outage in April 2011?](https://www.quora.com/Quora-Outage-April-21-22-2011/What-caused-the-Quora-problems-outage-in-April-2011)
- [Recovering from Amazon cloud outage](http://tomatohater.com/2011/04/21/recovering-amazon-cloud-outage/) by Drew Engelson of PBS.
  - PBS was affected for a while primarily because we do use EBS-backed RDS databases. Despite being spread across multiple availability-zones, we weren’t easily able to launch new resources ANYWHERE in the East region since everyone else was trying to do the same. I ended up pushing the RDS stuff out West for the time being.  [From Comment](http://don.blogs.smugmug.com/2011/04/24/how-smugmug-survived-the-amazonpocalypse/#comment-4737)

#### Amazon Web Services 讨论区

有一些有经验的人共享了很多相当不错的宕机的经历。

- [Amazon Web Services Discussion Forum](https://forums.aws.amazon.com/forum.jspa?forumID=30&start=0)
- [Cost-effective backup plan from now on?](https://forums.aws.amazon.com/thread.jspa?threadID=65860&tstart=0)
- [Life of our patients is at stake – I am desperately asking you to contact](https://forums.aws.amazon.com/thread.jspa?threadID=65649&tstart=0)
- [Why did the EBS, RDS, Cloudformation, Cloudwatch and Beanstalk all fail?](https://forums.aws.amazon.com/thread.jspa?threadID=65897&tstart=0)
- [Moved all resources off of AWS](https://forums.aws.amazon.com/thread.jspa?threadID=65896&tstart=0)
- [Any success stories?](https://forums.aws.amazon.com/forum.jspa?forumID=30&start=300)
- [Is the mass exodus from East going to cause demand problems in the West?](https://forums.aws.amazon.com/thread.jspa?threadID=65784&tstart=25)
- [Finally back online after about 71 hours](https://forums.aws.amazon.com/thread.jspa?threadID=65828&tstart=25)
- [Amazon EC2 features vs windows azure](https://forums.aws.amazon.com/thread.jspa?threadID=65834&tstart=25)
- [Aren’t Availability Zones supposed to be “insulated from failures”?](https://forums.aws.amazon.com/thread.jspa?threadID=65221&tstart=25)
- [What a lot of people aren’t realizing about the downtime:](https://forums.aws.amazon.com/thread.jspa?threadID=65850&tstart=0)
- [ELB CNAME](https://forums.aws.amazon.com/thread.jspa?threadID=32044&tstart=50&start=150)
- [Availability Zones were used in a misleading manner](https://forums.aws.amazon.com/thread.jspa?threadID=65457&tstart=425)
- [Tip: How to recover your instance](https://forums.aws.amazon.com/thread.jspa?threadID=65371&tstart=325)
- [Crying in Forum Gets Results, Silver-level AWS Premium Support Doesn’t](https://forums.aws.amazon.com/thread.jspa?threadID=65617&tstart=325)
- [Well-worth reading: “design for failure” cloud deployment strategy](https://forums.aws.amazon.com/thread.jspa?threadID=65780&tstart=25)
- [New best practice](https://forums.aws.amazon.com/thread.jspa?threadID=65749&tstart=25)
- [Don’t bother with Premium Support](https://forums.aws.amazon.com/thread.jspa?threadID=65136&tstart=475)
- [Best practices for multi-region redundancy](https://forums.aws.amazon.com/thread.jspa?threadID=65185&tstart=450)
- “[Postmortum](https://forums.aws.amazon.com/thread.jspa?threadID=65450&tstart=175)“
- [Learning from this case](https://forums.aws.amazon.com/thread.jspa?threadID=65513&tstart=125)
- [Amazon, still no instructions what to do?](https://forums.aws.amazon.com/thread.jspa?threadID=65388&tstart=525)
- [Anyone else prepared for an all-nighter?](https://forums.aws.amazon.com/thread.jspa?threadID=65338&tstart=550)
- [Is Jeff Bezos going to give a public statement?](https://forums.aws.amazon.com/thread.jspa?threadID=65811&tstart=100)
- [Rackspace, GoGrid, StormonDemand and Others](https://forums.aws.amazon.com/thread.jspa?threadID=65857&tstart=100)
- [Jeff Barr, Werner Vogels and other AWS persons – where have you been???](https://forums.aws.amazon.com/thread.jspa?threadID=65815&tstart=150)
- [After you guys fix EBS do I have do anything on my side?](https://forums.aws.amazon.com/thread.jspa?threadID=65168&tstart=175)
- [Need Help!!! Lives of people and billions in revenue are at risk now!!!](https://forums.aws.amazon.com/thread.jspa?threadID=65765&tstart=225)
- [I’ve Got A Suspicion](https://forums.aws.amazon.com/thread.jspa?threadID=65678&tstart=275)
- [Farewell EC2, Farewell](https://forums.aws.amazon.com/thread.jspa?threadID=65585&tstart=325)

There were also many many instances of support and help in the log.

#### 总结

- [Amazon EC2 outage: summary and lessons learned](http://blog.rightscale.com/2011/04/25/amazon-ec2-outage-summary-and-lessons-learned/) by RightScale
- [AWS outage timeline & downtimes by recovery strategy](http://www.randomhacks.net/articles/2011/04/25/aws-outage-timeline-and-recovery-strategy-downtimes) by Eric Kidd
- [The Aftermath of Amazon’s Cloud Outage](http://www.datacenterknowledge.com/archives/2011/04/25/the-aftermath-of-amazons-cloud-outage) by Rich Miller

#### 立场：这是用户的错

- [So Your AWS-based Application is Down? Don’t Blame Amazon](http://www.thestoragearchitect.com/2011/04/22/so-your-aws-based-application-is-down-dont-blame-amazon/) by The Storage Architect
- [The Cloud is not a Silver Bullet](http://stu.mp/2011/04/the-cloud-is-not-a-silver-bullet.html) by Joe Stump ([Hacker News](http://news.ycombinator.com/item?id=2482581) thread)
- [The AWS Outage: The Cloud’s Shining Moment](http://broadcast.oreilly.com/2011/04/the-aws-outage-the-clouds-shining-moment.html) by George Reese ([Hacker News](http://news.ycombinator.com/item?id=2477540) discussion)
- [Failing to Plan is Planning to Fail](http://blog.acrowire.com/cloud-computing/failing-to-plan-is-planning-to-fail) by Ted Theodoropoulos
- [Get a life and build redundancy/resiliency in your apps](https://groups.google.com/group/cloud-computing/browse_thread/thread/e8079a54e6a8c4b9/72756bf9e587869d?show_docid=72756bf9e587869d) on the Cloud Computing group

#### 立场：这是Amazon的错

- [Stop Blaming the Customers – the Fault is on Amazon Web Services](http://www.readwriteweb.com/cloud/2011/04/almost-as-galling-as-the.php) by Klint Finley
- [AWS is down: Why the sky is falling](http://justinsb.posterous.com/aws-down-why-the-sky-is-falling) by Justin Santa Barbara  ([Hacker News](http://news.ycombinator.com/item?id=2471899) thread)
- [Amazon Web Services are down](http://news.ycombinator.com/item?id=2469838) – Huge Hacker News thread

#### 教训和启示

- [People Using Amazon Cloud: Get Some Cheap Insurance At Least](http://smoothspan.wordpress.com/2011/04/23/people-using-amazon-cloud-get-some-cheap-insurance-at-least/) by Bob Warfield
- [Basic scalability principles to avert downtime](http://ronaldbradford.com/blog/basic-scalability-principles-to-avert-downtime-2011-04-23) by Ronald Bradford
- [Amazon crash reveals ‘cloud’ computing actually based on data centers](http://www.itworld.com/cloud-computing/158517/amazon-crash-reveals-cloud-computing-actually-based-data-centers) by Kevin Fogarty
- [Seven lessons to learn from Amazon’s outage](https://www.zdnet.com/blog/saas/seven-lessons-to-learn-from-amazons-outage/1296) By Phil Wainewright
- [The Cloud and Outages : Five Key Lessons](http://www.cloudsigma.com/en/blog/2011/04/23/21-cloud-outages-lessons-learned) by Patrick Baillie ([Cloud Computing Group](https://groups.google.com/group/cloud-computing/browse_thread/thread/6e9549afbff6386f/05919d8527c69a09?show_docid=05919d8527c69a09#) discussion)
- [Some thoughts on outages](http://till.klampaeckel.de/blog/archives/151-Some-thoughts-on-outtages.html) by Till Klampaeckel
- [Amazon.com’s real problem isn’t the outage, it’s the communication](http://www.geekwire.com/2011/amazoncoms-real-problem-outage-communication) by Keith Smith
- [How to work around Amazon EC2 outages](http://webmonkeyuk.wordpress.com/2011/04/21/how-to-work-around-amazon-ec2-outages/) by James Cohen ([Hacker News](http://news.ycombinator.com/item?id=2471258) thread)
- [Today’s EC2 / EBS Outage: Lessons learned](http://agilesysadmin.net/ec2-outage-lessons) on Agile Sysadmin
- [Amazon EC2 has gone down -what would a prefered hosting platform be?](http://www.focus.com/questions/information-technology/amazon-ec2-has-gone-down--what-would-prefered-hosting-be/) on Focus
- [Single Points of Failure](http://cloudability.com/single-points-of-failure) by Mat
- [Coping with Cloud Downtime with Puppet](http://www.reddit.com/r/programming/comments/gvac7/coping_with_cloud_downtime_with_puppet/)
- [Amazon Outage Concerns Are Overblown](http://timcrawford.org/2011/04/21/amazon-outage-concerns-are-overblown/) by Tim Crawford
- [Where There Are Clouds, It Sometimes Rains](http://claylo.com/post/4817029650/where-there-are-clouds-it-sometimes-rains) by Clay Loveless
- [Availability, redundancy, failover and data backups at LearnBoost](http://blog.learnboost.com/blog/availability-redundancy-and-failover-at-learnboost/)  by Guillermo Rauch
- [Cloud hosting vs colocation](http://chrischandler.name/the-real-cost-of-cloud-hosting) by Chris Chandler ([Hacker News](http://news.ycombinator.com/item?id=2482123) thread)
- [Amazon’s EC2 & EBS outage](http://arnon.me/2011/04/amazons-ec2-ebs-outage/) by Arnon Rotem-Gal-Oz

#### Vendor很生气

- [Amazon Outage Proves Value of Riak’s Vision](http://www.productionscale.com/home/2011/4/22/on-clouds-and-spofs-or-the-great-aws-outage-of-april-2011.html#axzz1KZPTwX4z) by Basho
- [Magical Block Store: When Abstractions Fail Us](http://joyeur.com/2011/04/24/magical-block-store-when-abstractions-fail-us/) by Mark Joyent ([Hacker News](http://news.ycombinator.com/item?id=2479613) discussion)
- [On Cascading Failures and Amazon’s Elastic Block Store](http://joyeur.com/2011/04/22/on-cascading-failures-and-amazons-elastic-block-store/) by Jason
- [An unofficial EC2 outage postmortem – the sky is not falling](http://cloudharmony.com/b/2011/04/unofficial-ec2-outage-postmortem-sky-is.html) from CloudHarmony

（全文完）

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
