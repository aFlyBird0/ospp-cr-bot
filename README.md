# ospp-cr-bot
开源之夏 - 一种通用的 Code Review 机器人

## 试运行/测试
### 实测
1. 将 `config.yml` 中的对应信息替换为自己的，并设置 `github_token` 环境变量。
2. 运行 `cmd/main.go` 中的主函数。
### mock
1. 将 `examples/mini-for-mock.yaml` 重命名为 `config.yaml` 并放置到根目录下。
2. 将 `internal/notify/main.go` 中的  
`import _ "github.com/devstream-io/devstream/ospp-cr-bot/internal/pkg/git/github"` 替换为，  
`import _ "github.com/devstream-io/devstream/ospp-cr-bot/internal/pkg/git/githubMock"`
3. 运行 `cmd/main.go` 中的主函数。

## 开发进度概要
备注：
* 先不实现 helm 部署
* 先实现定时任务主动拉取
  * webhook 的实现之后再商讨。webhook 感觉有两种。
    * 一种是暴露个拉取接口给其他项目（如机器人）调用
    * 一种是直接接 github 的 webhook
    * 我感觉前者更贴合业务逻辑？

### 已实现接口：
* 已实现支持任意代码托管平台、任意通讯平台的接口
* 任意代码托管平台的用户映射到任意通讯平台，一个仓库可配钉钉群、飞书群等不同群

### 未实现：
* 具体分析到底是谁阻塞了 PR（涉及业务和详细的 github api，后面慢慢完善接口和填充业务逻辑就好）
* 多次提醒、提醒时间自动变动 （主要是业务逻辑待商讨）
  * 怎么判断是否是长时间未响应
  * 如何判断 reviewer 或 committer 响应了，是点了链接还是要提交判断
  * 响应之后，下一次通知时间又是多少
* 代码托管平台、通讯平台的 client 实际实现，例如 auth 登录、实际接口调用。（github的已简单实现）
* 没有记录 reviewer、user 等人对通知的响应动作，同样主要是因为业务逻辑需要确认。

# 架构简单介绍
note: 
* 下文的「注册」指的是把信息注册到内存中
* 代码托管平台统称为 git platform
* 通讯平台（飞书、钉钉）统称为 community

pkg:
- pkg.community 定义了所有的通讯类平台的接口，抽象出了平台、用户等名词性接口，也包含了平台注册、消息发送、用户注册、群注册等动词性接口
- pkg.git 定义了所有的代码托管类平台的接口，抽象出了平台、Repo、Issue、PR、Comment等名词性接口，也包含了平台注册、仓库信息与动态拉取等动词性接口
- pkg.message 正在设计，目前是放了消息相关的东西
- pkg.union 联合了不同来源的信息与接口。如联合代码托管平台的用户和通讯平台的用户

internal:
- internal.community 各个通讯类平台的 client 实现，在 init() 中直接通过 pkg.community.RegisterCommunity() 完成注册。
- internal.git 各个代码托管类平台的 client 实现，在 init() 中直接通过 pkg.git.RegisterPlatform() 完成注册。
- internal.notify 暂时性的通知入口(main)
