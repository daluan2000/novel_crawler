## 免责声明
### 1.本项目的目的是学习和练习爬虫技术，本项目提供的爬虫源代码仅用学习，请勿用于商业盈利。
### 2.用户使用本系统从事任何违法违规的事情，一切后果由用户自行承担，作者不承担任何责任。
### 3.如有侵犯权利，请联系作者删除，邮箱：daluan2000@qq.com。
### 4.下载本站源码则代表你同意上述的免责声明协议

## 功能

爬取网站上小说章节的内容，并以txt形式保存在本地。

程序运行需要的参数如下，带默认值default的参数可以不输入，其余参数必须输入：

| 参数名  | 参数值                                          | 样例                              |
|------|----------------------------------------------|---------------------------------|
| -f   | 保存在本地的文件名                                    | 斗破苍穹                            |
| -u   | 小说章节列表的url链接                                 | https://www.52bqg.org/book_361/ |
| -ft  | 不改变标题为1，填充标题编号为2，不输入该参数默认为1 (default 1)      | 1                               |
| -log | 默认为1，打印详细log为2 (default 1)                   | 1                               |
| -rc  | 重新尝试的次数，默认为10 (default 10)                   | 10                              |
| -rs  | retry时的休眠时间，默认250ms (default 250ms)          | 250ms                           |
| -st  | 保存tittle为1，不保存title为2，不输入该参数默认为1 (default 1) | 1                               |

使用样例如下：
```shell
novel_crawler.exe -f 斗破苍穹 -u https://www.52bqg.org/book_361/
```



## 已支持网站

**注意，这里文档没有更新，下面内容只展示部分已支持的网站，**

**所有已支持网站的信息位于[/crawler/info/info.go](./crawler/info/info.go)，由于时间久远，可能有些网站已经g了**

对于未支持的网站，可使用[自定义配置文件](#自定义配置文件)功能自行添加


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

7. www.00txt.com 科幻小说网，出版科幻小说

```shell
.\novel_crawler.exe -u http://www.00txt.com/santi/ -f 三体
```

8. www.1688by.com 好笔阁

```shell
.\novel_crawler.exe -f 我真不是邪神走狗 -u https://www.1688by.com/book/203501
```

9. www.wbsz.org 完本神站

```shell
.\novel_crawler.exe -f 优等生不需要超能力.txt -u https://www.wbsz.org/18874/
```

10. www.xbiqugeo.com 新笔趣阁，使用样例如下：

sb网站禁止搜索

```shell
.\novel_crawler.exe -f 少年歌行 -u https://www.xbiqugeo.com/shu/6420/  
```

11. www.zrfsxs.com 择日小说网，使用样例如下：

```shell
.\novel_crawler.exe -f 深空彼岸 -u https://www.zrfsxs.com/xiaoshuo/42/
```

12. youyouxs.com 友友小说网

```shell
.\novel_crawler.exe -u https://youyouxs.com/xs_350417/zjml_1 -f 超能力者不想受欢迎
```

13. www.biqge.org 笔奇阁

```shell
.\novel_crawler.exe -u https://www.biqge.org/book/17130/ -f 修仙就是这样的
```

14. m.wfxs.tw 微風小說網

这个网站有妖，有时爬的成功，有时爬不成功，报错`read tcp 192.168.16.86:58143->104.26.2.82:443: wsarecv: An existing connection was forcibly closed by the remote host.`。太底层了我看不懂，有待学习。

```shell
.\novel_crawler.exe -u https://m.wfxs.tw/xs-2217283/ -f 下山就無敵，總裁倒追我 
```

## 自定义配置文件

### 数据结构

在novel_crawler.exe所在目录下创建info.yml文件，在文件中写入`map[string]Info`类型的信息，key值设置为info，在源代码中`Info`结构体如下：

```go
// Info结构体
type Info struct {
	// 目录页面，各章节标题a标签的选择器
	ASelector       string
	// 章节内容页面，小说内容的选择器
	ContentSelector string
	// html文本替换字符串
	StrReplace map[string]string
    // html文本替换字符串，正则表达式形式
	RegReplace map[string]string
	// html文本中要删除的标签对应的选择器
	RemoveSelector []string
	
	FrequencyLimit
	NextChapterList
	NextContent
}

// 并发限制，有默认值
type FrequencyLimit struct {
	// 并发数量限制
	Concurrent int
	// 每次请求后线程的休眠时间
	Gap time.Duration
}

// 目录为分页展示时，需要加上此部分信息
type NextChapterList struct {
	// 如果分页展示，设置为true
	MultiPageChapterList    bool
	// 目录页面中，下一页a标签的选择器
	ChapterListNextSelector string
	// 目录页面中，下一页a标签应包含的文本
	ChapterListNextStr      string
}

// 章节内容分页展示时，需要加上此部分信息
type NextContent struct {
    // 如果分页展示，设置为true
	MultiPageContent    bool
	// 章节内容页面中，下一页a标签的选择器
	ContentNextSelector string
	// 章节内容页面中，下一页a标签应包含的文本
	ContentNextStr      string
}

```

### 配置文件格式

info.yml文件的具体格式如下面所示，可以同时配置多个网站。

其中ASelector和ContentSelector是必需字段，其余字段如果没有特殊需求可以省略。

程序内部已经适配了一些网站，如果想要爬取未适配的网站，可以在yml文件中按照格式要求自行添加相关信息。

```yaml
Info:
  
  "www.wbsz.org":
    ASelector: ".chapter > ul > li > a"
    ContentSelector: ".readerCon"
    RemoveSelector: ['script']
    FrequencyLimit:
      Concurrent: 4
      Gap: 250ms


  "www.xbiqugeo.com":
    ASelector: ".section-box:nth-child(4) > ul > li > a"
    ContentSelector: "#content"
    RemoveSelector: ["a", "div"]
    FrequencyLimit:
      Concurrent: 4
      Gap: 250ms
    NextChapterList:
      MultiPageChapterList: true
      ChapterListNextSelector: ".listpage > .right > a"
      ChapterListNextStr: "下一页"
    NextContent:
      MultiPageContent: true
      ContentNextStr: "下一页"
      ContentNextSelector: "#next_url"
```