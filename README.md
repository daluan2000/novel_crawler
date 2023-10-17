## 声明
编写本爬虫程序的初衷是学习和练习爬虫技术，本程序仅作学习交流使用，禁止用于任何商业用途。

如果本程序侵犯了您的权益，请通过邮箱daluan2000@qq.com联系我，我将尽快删除侵权部分。

## 功能

爬取网站上小说章节的内容，并以txt形式保存在本地。

程序运行需要的参数如下：

| 参数名 | 参数值          | 样例                              |
|-----|--------------|---------------------------------|
| -f  | 保存在本地的文件名    | 斗破苍穹                            |
| -u  | 小说章节列表的url链接 | https://www.52bqg.org/book_361/ |

使用样例如下：
```shell
novel_crawler.exe -f 斗破苍穹 -u https://www.52bqg.org/book_361/
```



## 已支持网站

本爬虫的适应性非常强，只需经过简易的拓展，便基本能够爬取所有的公开小说网站（指那些不需要登陆就能阅读小说的网站）。

我这里只随便弄了几个网站作为样例，如果需要爬更多的网站那么告诉我网址就好，我这边稍微修改下程序就可以了。


### 第一类网站

1. www.2biqu.com 笔趣阁，使用样例如下：

```shell
.\novel_crawler.exe -f 择日飞升 -u https://www.2biqu.com/biqu5396/
```

2. www.bige3.cc 笔趣阁，使用样例如下：

```shell
.\novel_crawler.exe -f 神秘复苏 -u https://www.bige3.cc/book/66/
```

3. www.52bqg.org 笔趣阁，使用样例如下：

```shell
.\novel_crawler.exe -f 深空彼岸 -u https://www.52bqg.org/book_99524/
```

4. www.ujxsw.net 悠久小说网，使用样例如下：

```shell
.\novel_crawler.exe -f 我的26岁女房客 -u http://www.ujxsw.net/read/15871/ 
```
5. www.tianyabook.com 天涯书库，使用样例如下：

```shell
.\novel_crawler.exe -f 终极斗罗 -u https://www.tianyabook.com/shu/40027.html
```

### 第二类网站

1. www.xbiqugeo.com 新笔趣阁，使用样例如下：

```shell
.\novel_crawler.exe -f 少年歌行 -u https://www.xbiqugeo.com/shu/6420/  
```

2. www.zrfsxs.com 择日小说网，使用样例如下：

```shell
.\novel_crawler.exe -f 深空彼岸 -u https://www.zrfsxs.com/xiaoshuo/42/
```