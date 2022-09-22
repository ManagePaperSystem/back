# back

## 项目相关信息

框架：Beego

数据库：sqlite3

本项目搭配 Vue 工程使用，代码为本项目的另一个仓库 https://github.com/ManagePaperSystem/front。




## 环境配置

本项目使用 Beego 框架 创建 API 工程，如果需要在您的本地计算机中运行，请提前下载 bee 并配置好相关的环境变量。

建议您使用 docker 等方法创建虚拟环境，或者直接访问我的网站调用相关 api，地址为[124.222.69.106:8055]()。



## 项目运行

运行本 API 工程时执行

```sh
bee run -gendoc=true -downdoc=true
```

即可在 8055 端口运行。

请注意，如果想要被前端项目访问那么前端也需要做相关的配置，详见`front/README.md` 。

另外，由于个人邮件和密钥的私密性，我将 `models/base.go` 中 55，56 行的相关配置删除，请自行填写个人的邮箱和密钥进行使用。

由于配置环境的过程比较复杂且充满艰辛，我们仍建议您直接访问我的网站，地址为[124.222.69.106:8080]()。

由于服务器配置较低，如果遇到访问问题，请 QQ 联系我们:281597094。提前为可能为您带来的麻烦表达歉意。



## 特别鸣谢

+ https://github.com/beego/beego ：一 个 伟 大 的 框 架。

+ https://github.com/beego/bee ：为本项目提供了环境配置和运行的能力。

+ https://github.com/mattn/go-sqlite3 ：为本项目提供了 sqlite3 处理工具。
+ https://github.com/swagger-api/swagger-ui ：为本项目提供 API 文档页面。
+ https://github.com/go-gomail/gomail ：为本项目提供邮件收发的功能。
+ 我的室友们：宵衣旰食，通力合作。
+ 我的父母：在我遇到瓶颈时为我提供~~金钱~~鼓励。
