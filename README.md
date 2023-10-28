## 免责声明
### 1.本项目的目的是学习和练习爬虫技术，本项目提供的爬虫源代码仅用学习，请勿用于商业盈利。
### 2.用户使用本系统从事任何违法违规的事情，一切后果由用户自行承担，作者不承担任何责任。
### 3.如有侵犯权利，请联系作者删除，邮箱：daluan2000@qq.com。
### 4.下载本站源码则代表你同意上述的免责声明协议

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
6. www.trxs.cc 同人小说网，二次元比较多，使用样例如下：

```shell
.\novel_crawler.exe -f 我的后桌居然是珈百璃 -u http://www.trxs.cc/tongren/3650.html
```

6. www.00txt.com 科幻小说网，出版科幻小说

```shell
go run .\main.go -u http://www.00txt.com/santi/ -f 三体
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

3. youyouxs.com 友友小说网

ps：这个网站限制ip访问频次，我没钱买ip池，所以只能限制程序的并发量，爬取速度会比较慢一些
ps：这个网站会封ip，建议连接手机热点

```shell
.\novel_crawler.exe -u https://youyouxs.com/xs_350417/zjml_1 -f 超能力者不想受欢迎
```