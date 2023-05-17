package main_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title   string
	Path    string
	Date    string
	Author  string
	ViewNum string
}

func Test(t *testing.T) {
	var articles []Article
	files, _ := ioutil.ReadDir("content/list")
	filenames := []int{}
	for _, v := range files {
		if v.Name() == "featured.html" {
			continue
		}
		name := strings.ReplaceAll(v.Name(), ".html", "")
		num, _ := strconv.Atoi(name)
		filenames = append(filenames, num)
	}
	sort.Ints(filenames)
	for _, v := range filenames {
		data, err := ioutil.ReadFile("content/list/" + strconv.Itoa(v) + ".html")
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		head := doc.Find("header.entry-header")
		head.Each(func(i int, s *goquery.Selection) {
			title := s.Find("span.screen-reader-text").First()
			time := s.Find("time")
			author := s.Find("a[rel=\"author\"]")
			viewNum := s.Find("h5").Text()
			bookmark := s.Find("a[rel=\"bookmark\"]").First()
			v := strings.Split(viewNum, "条评论")
			viewCount := ""
			if len(v) > 1 {
				viewCount = v[1]
			}
			articles = append(articles, Article{Title: title.Text(), Path: bookmark.AttrOr("href", ""), Date: time.Text(), Author: author.Text(), ViewNum: viewCount})
		})

		t := template.New("test")
		t = template.Must(t.Parse(htmlTxt))
		f, _ := os.OpenFile("content/list/featured.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
		t.Execute(f, articles)
	}
}

var htmlTxt = `
<!DOCTYPE html><!--[if IE 8]>
<html id="ie8" dir="ltr" lang="zh-CN"
	prefix="og: https://ogp.me/ns#" >
<![endif]--><!--[if !(IE 8) ]><!--><html dir="ltr" lang="zh-CN" prefix="og: https://ogp.me/ns#"><!--<![endif]--><head><meta charset="UTF-8"/><title>酷壳 – CoolShell.cn</title><link rel="stylesheet" type="text/css" href="/assets/all.min.css"/><link rel="shortcut icon" href="/assets/favicon.png"/></head>
<body class="home blog group-blog">

<div id="page" class="hfeed site">
<header id="masthead" role="banner">

<div id="cc_spacer"></div>
<div class="site-header">

<div class="site-branding">
<a class="home-link" href="/" title="酷 壳 – CoolShell" rel="home">
<h1 class="site-title">酷 壳 – CoolShell</h1>
<h2 class="site-description">享受编程和技术所带来的快乐 – Coding Your Ambition</h2>
</a>
</div>
</div>
</header>
<div class="container">
<div class="row">

</div>
</div>
<div id="content" class="site-content">
<div class="container">
<div class="row">
<div id="primary" class="content-area  col-md-12" style="margin-top: 40px;">
<main id="main" class="site-main" role="main">
<article id="post-22422" class="post-content post-22422 post type-post status-publish format-standard hentry category-itnews category-misc category-602 tag-aws tag-lambda tag-microservice tag-serverless tag-step-function">
<header class="entry-header">
<span class="screen-reader-text">文章列表</span>
<h1 class="entry-title">推荐文章</h1>
<div class="entry-meta">
</div>
</header>
<div class="entry-content">
	<ul class="featured-post">
{{ range . }}
		<li>
		  <span>{{ .Date }}</span>
			<a href="{{ .Path }}" title="{{ .Title }}">{{ .Title }}</a>
			<span>{{ .ViewNum }} - {{ .Author }}</span>
		</li>
{{- end }}
	</ul>
</div>
<footer class="entry-footer">
</footer>
</article>
<nav class="navigation posts-navigation" role="navigation">
<h2 class="screen-reader-text">Posts navigation</h2>
</nav>
</main>
</div>

</div>
</div>

</div>

</div>

</body></html>
`
