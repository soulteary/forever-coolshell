#!/usr/bin/env python3
"""Convert archived CoolShell article bodies to Markdown."""
import argparse
import re
import sys
from pathlib import Path
from urllib.parse import urlsplit, urlunsplit

from bs4 import BeautifulSoup
from markdownify import markdownify

ROOT = Path(__file__).resolve().parents[1]


def convert(source, known):
    soup = BeautifulSoup(source.read_text(encoding='utf-8'), 'html.parser')
    title = soup.select_one('.entry-title')
    body = soup.select_one('.entry-content')
    if title is None or body is None:
        raise ValueError('缺少文章标题或正文')
    for node in body.select('script, style, .post-ratings, .post-ratings-loading, .sharedaddy'):
        node.decompose()
    for node in body.select('[href], [src]'):
        attr = 'href' if node.has_attr('href') else 'src'
        url = urlsplit(node[attr])
        if url.netloc and url.netloc not in ('coolshell.cn', 'www.coolshell.cn'):
            continue
        match = re.fullmatch(r'/articles/(\d+)\.html', url.path)
        if attr == 'href' and match and match[1] in known:
            node[attr] = urlunsplit(('', '', match[1] + '.md', url.query, url.fragment))
        elif not url.netloc and url.path.startswith('/'):
            node[attr] = 'https://coolshell.cn' + node[attr]
    def language(pre):
        return pre.get('data-enlighter-language', '')
    text = markdownify(str(body), heading_style='ATX', bullets='-',
                       code_language_callback=language)
    return '# ' + title.get_text(' ', strip=True) + '\n\n' + text.strip() + '\n'


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--input', type=Path, default=ROOT / 'content/articles')
    parser.add_argument('--output', type=Path, default=ROOT / 'markdown/articles')
    parser.add_argument('--overwrite', action='store_true', help='覆盖已有 Markdown 文件')
    args = parser.parse_args()
    files = sorted(args.input.glob('*.html'))
    if not files:
        parser.error(f'没有找到 HTML 文章: {args.input}')
    args.output.mkdir(parents=True, exist_ok=True)
    known = {p.stem for p in files}
    written = skipped = failed = 0
    for source in files:
        target = args.output / (source.stem + '.md')
        if target.exists() and not args.overwrite:
            skipped += 1
            continue
        try:
            result = convert(source, known)
            target.write_text(result, encoding='utf-8')
            written += 1
        except Exception as error:
            failed += 1
            print(f'{source}: {error}', file=sys.stderr)
    print(f'已转换 {written} 篇，跳过 {skipped} 篇，失败 {failed} 篇；输出: {args.output}')
    return bool(failed)


if __name__ == '__main__':
    sys.exit(main())
