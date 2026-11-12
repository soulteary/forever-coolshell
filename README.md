# 酷壳 - CoolShell 电子存档

谨纪念和保存陈皓先生的公开的技术分享内容，感恩皓叔对中文互联网，尤其是技术领域无私的分享。

![](.github/farewell.png)

## 使用方法

使用方法有两种：使用直接编译好的可执行文件，或使用 Docker 运行。

方法一：在[Release 页面](https://github.com/soulteary/forever-coolshell/releases/)下载适合你的操作系统的可执行文件后，使用下面的命令运行即可。

```bash
./coolshell
```

方法二：使用 Docker 下载镜像之后，指定合适的端口号，一键启动容器即可。

```bash
# 下载镜像
docker pull soulteary/forever-coolshell:v1.0.0
# 启动镜像
docker run --rm -it -p 8080:8080 soulteary/forever-coolshell:v1.0.0
```

网站电子存档默认使用 `8080` 端口，如需改变端口，请指定环境变量 `PORT`，如：

```bash
# 直接运行
PORT=10240 ./coolshell
# 使用容器运行
docker run --rm -it -e PORT=10240 -p 10240:10240 soulteary/forever-coolshell:v1.0.0
```

## 最后

[@haoel](https://github.com/haoel)，再也听不到你吐槽 WordPress 数据库挂掉、CDN 等维护问题了。

你的博客内容会一直留存下去的。

--EOF 晓白

## 导出 Markdown

批量提取 `content/articles/*.html` 的文章标题和正文，排除导航、评论和评分控件，保留代码块、表格、图片及链接。

```bash
python3 -m venv .venv
.venv/bin/pip install -r scripts/requirements-markdown.txt
.venv/bin/python scripts/html_to_markdown.py
```

默认输出到 `markdown/articles/<文章ID>.md`，不会修改原始 HTML。已有文件默认跳过，添加 `--overwrite` 可重新生成；也可用 `--input`、`--output` 指定目录。

站内文章链接会改为同目录的 `.md` 链接；图片等资源使用原站绝对地址，需联网访问（脚本不下载资源）。
