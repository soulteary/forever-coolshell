# 一个C的序列化库tpl

tpl(<http://tpl.sourceforge.net/>)是一个开源的小项目，其主要是提供一个可以序列化或反序列化C语言数据的一个API函数库。tpl号称是最有效率的也是最快的，它可以把你内存中的数据存放到文件中，并可以在你需要的时候用文件中反序例化到内存变量中。而且这个函数库没有依赖于别的函数库。

下面是一个简单的示例（来源于其主页）

把一个数组（“序号”和“人名”）序例化到文件中。

```c
#include "tpl.h"

int main(int argc, char *argv[]) {
    tpl_node *tn;
    int id=0;
    char *name, *names[] = { "joe", "bob", "cary" };

    tn = tpl_map("A(is)", &id, &name);

    for(name=names[0]; id < 3; name=names[++id]) {
        tpl_pack(tn,1);
    }

    tpl_dump(tn, TPL_FILE, "users.tpl");
    tpl_free(tn);
}

```

把上面那个序列化到文件的“序号”和“人名”反序列化回来。

```c
#include "tpl.h"

int main(int argc, char *argv[]) {
    tpl_node *tn;
    int id;
    char *name;

    tn = tpl_map("A(is)", &id, &name);
    tpl_load(tn, TPL_FILE, "users.tpl");

    while ( tpl_unpack(tn,1) > 0 ) {
        printf("id %d, user %s\n", id, name);
        free(name);
    }
    tpl_free(tn);
}

```

更详细的使用说明请看其文档：  

<http://tpl.sourceforge.net/userguide.html>

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
