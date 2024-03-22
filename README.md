<div align="center">

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ppdxzz/ChatBot?logo=go)
![GitHub License](https://img.shields.io/github/license/ppdxzz/ChatBot?label=License&color=red)
![GitHub repo size](https://img.shields.io/github/repo-size/ppdxzz/ChatBot?label=Size&color=green)

</div>

### 📖 项目介绍
ChatBot 是GoLang开发的聊天机器人项目，它基于<a href="https://github.com/eatmoreapple/openwechat"> OpenWechat </a>SDK包开发上层业务模块部分，旨在集成各种适合娱乐的接口服务，为机器人增加各式各样的功能。


### 🔥 快速开始
打开 GoLang IDE，克隆项目至本地，直接运行`main.go`即可启动ChatBot服务。注意如果你需要大模型聊天以及天气服务，你应该修改`config/config.yaml`中的配置。如需加入自己想要的功能，直接修改源码实现对应的逻辑即可。
```yaml
# 文心一言
qianfan:
  AccessKey: key
  SecretKey: secret
# 天行数据API
tianapi:
  key: key
# 高德开发平台
amap:
  key: key
```
***接口来源***
- <a href="https://developer.baidu.com/">文心一言</a>
- <a href="https://www.tianapi.com/">天行数据</a>
- <a href="https://lbs.amap.com/">高德开放平台</a>


### 🚀 功能
- [x] 智能问答
- [x] 进群提醒
- [x] 早安心语
- [x] 喝水提醒
- [ ] 天气预报（支持默认城市）
- [ ] 群签到
- [ ] 黑名单
- [ ] 热搜
- [ ] 历史今天
- [ ] 星座运势
- [ ] 发言排行榜
- [ ] AI 画图


### 💻 部署
#### 二进制文件
1. 设置编译配置
```shell
# 禁用CGO
set CGO_ENABLED=0
# 设置编译目标系统
set GOOS=linux
# 设置编译目标平台的64位 x86 架构
set GOARCH=amd64
```
2. 打包
```shell
# ChatBot可替换成你打包后的名称
go build -o ChatBot
```
3. 文件授权

`ChatBot`二进制可执行文件和`config/config.yaml`上传至目标服务器，放在统一目录下，`chmod +x ChatBot` 给文件设置为可执行权限，要不第一次没有权限无法执行。
4. 启动
```shell
# nohup 后台启动
nohup ./ChatBot &
```
5. 日志登录
正常来讲，这样你就已经启动了服务，这个时候你应该在日志中找到打印出的登录URL并成功登录，ChatBot启动完毕。
```shell
# 查看日志
tail -f nohup.out
```


### 📣 免责声明
- 本项目只供个人学习和研究目的，切勿使用项目做任何商业用途或牟利。微信平台不允许开发此类软件，因此使用本项目可能违反微信平台的使用条款，对于因使用本项目而导致的任何后果概不负责，包括但不限于微信账号被封禁或受到其他处罚。

- 本项目中的所有内容，包括但不限于代码、文档和其他材料，不允许任何公众号、自媒体进行任何形式的转载、发布。

- 本项目可能包含指向第三方网站或资源的链接。这些链接仅为方便用户而提供，作者对于这些第三方网站或资源的可用性或准确性不承担任何责任。

- 作者保留在任何时候修改本免责声明的权利。变更将在本页面上发布，并自发布之日起生效，您应定期查看本页面以了解任何变更。

- 如果您对本免责声明或本项目有任何疑问或意见，请通过<a href="mailto:ppdxzz@outlook.com">ppdxzz@outlook.com</a>与我们联系。

