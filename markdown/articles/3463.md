# 图解SQL的Join

对于SQL的Join，在学习起来可能是比较乱的。我们知道，[SQL的Join语法](https://en.wikipedia.org/wiki/Join_(SQL))有很多inner的，有outer的，有left的，有时候，对于Select出来的结果集是什么样子有点不是很清楚。Coding Horror上有[一篇文章](http://www.codinghorror.com/blog/2007/10/a-visual-explanation-of-sql-joins.html)（实在不清楚为什么Coding Horror也被墙）通过 文氏图 [Venn diagrams](https://en.wikipedia.org/wiki/Venn_diagram) 解释了SQL的Join。我觉得清楚易懂，转过来。

假设我们有两张表。

- **Table A** 是左边的表。
- **Table B** 是右边的表。

其各有四条记录，其中有两条记录是相同的，如下所示：

```
id name       id  name
-- ----       --  ----
1  Pirate     1   Rutabaga
2  Monkey     2   Pirate
3  Ninja      3   Darth Vader
4  Spaghetti  4   Ninja
```

下面让我们来看看不同的Join会产生什么样的结果。

| ``` SELECT * FROM TableA INNER JOIN TableB ON TableA.name = TableB.name  id  name       id   name --  ----       --   ---- 1   Pirate     2    Pirate 3   Ninja      4    Ninja ```  **Inner join** 产生的结果集中，是A和B的交集。 | Venn diagram of SQL inner join |
| --- | --- |
| ``` SELECT * FROM TableA FULL OUTER JOIN TableB ON TableA.name = TableB.name  id    name       id    name --    ----       --    ---- 1     Pirate     2     Pirate 2     Monkey     null  null 3     Ninja      4     Ninja 4     Spaghetti  null  null null  null       1     Rutabaga null  null       3     Darth Vader ```  **Full outer join** 产生A和B的并集。但是需要注意的是，对于没有匹配的记录，则会以null做为值。 | Venn diagram of SQL cartesian join |
| ``` SELECT * FROM TableA LEFT OUTER JOIN TableB ON TableA.name = TableB.name  id  name       id    name --  ----       --    ---- 1   Pirate     2     Pirate 2   Monkey     null  null 3   Ninja      4     Ninja 4   Spaghetti  null  null ```  **Left outer join** 产生表A的完全集，而B表中匹配的则有值，没有匹配的则以null值取代。 | Venn diagram of SQL left join |
| ``` SELECT * FROM TableA LEFT OUTER JOIN TableB ON TableA.name = TableB.name WHERE TableB.id IS null   id  name       id     name --  ----       --     ---- 2   Monkey     null   null 4   Spaghetti  null   null ```  产生在A表中有而在B表中没有的集合。 | join-left-outer.png |
| ``` SELECT * FROM TableA FULL OUTER JOIN TableB ON TableA.name = TableB.name WHERE TableA.id IS null OR TableB.id IS null  id    name       id    name --    ----       --    ---- 2     Monkey     null  null 4     Spaghetti  null  null null  null       1     Rutabaga null  null       3     Darth Vader ```  产生A表和B表都没有出现的数据集。 | join-outer.png |

还需要注册的是我们还有一个是“交差集” **cross join**, 这种Join没有办法用文式图表示，因为其就是把表A和表B的数据进行一个N\*M的组合，即笛卡尔积。表达式如下：

```
SELECT * FROM TableA
CROSS JOIN TableB
```

这个笛卡尔乘积会产生 4 x 4 = 16 条记录，一般来说，我们很少用到这个语法。但是我们得小心，如果不是使用嵌套的select语句，一般系统都会产生笛卡尔乘积然再做过滤。这是对于性能来说是非常危险的，尤其是表很大的时候。

***更新:2014年3月30日***

![](https://coolshell.cn/uploads/2011/01/SQL-Join.jpg)

(全文完)

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
