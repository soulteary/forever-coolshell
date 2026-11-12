# Python脚本如何对文件通配符匹配

有时候，我们可能会写一些轻量级的脚本去处理很多符合某种pattern的文件，例如“某目录下的 \*logfile.csv” 但是，我们大多数脚本的参数都是 sys.argv, 如何解析 wildcard 匹配呢？

#### test.py

```python
 from glob import glob
...
if __name__ == "__main__":
    file_names = glob(sys.argv[1])
    for file_name in file_names:
        do_something(file) 
```

这样就可以像使用其他终端命令一样使用脚本test.py 进行wildcard匹配了

#### >> test.py ./\*logfile.csv

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
