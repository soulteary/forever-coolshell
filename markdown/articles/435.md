# Python中实现多属性排序

我们有一组记录:

```python
list_records =
(
 (department, name, salary),
 (department, name, salary),
 ...
 (department, name, salary)
)
```

然后我们想进行类似 MS – Excel 里的 “then sort by” 中的功能一样先基于department排序，然后再在部门内按照salary排序。

其他编程语言可能相对复杂，我这里写出一个用Python实现的最简方法（也许有比这个还短的，来挑战吧）

```python
list_records.sort(
 key = lambda l: (l[0], l[2])
)
```

这个就是函数是编程的好处，可以无中生有的构造出一个没有名字的inline函数。假设我们有另外一个dictionary\_age 是保存的 { name: ages }， 我们还可以简单的实现基于外部属性进行排序。例如，如果我们想先按照部门排序，然后在部门里按照年龄排序，我们可以写：

```python
list_record.sort(
 key = lambda l:( l[0], dictionary_age(l[1]) )
)
```

如果需要降序排列，可以设置 revserse = True; 如果想基于两个属性，一个升序，一个降序，可以试试将其中一个构造一个外部规则，然后如同上例子中的dictionary\_age一样传递进去。

Done!

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
