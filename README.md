# 简介
项目目前包含两个部分,一个是`kubectx`子命令用于kube context的管理,另一个为`kubectlx`子命令属于前期的`kubectx`的个性化扩展功能探索需要已安装`kubectl`工具,后面会考虑改为其他形式.
## `kubex`子命令
### 安装
通过git clone此项目在go环境下编译
### 使用
建议使用`alias`的方式简化子命令使用,例如`alias cx="kubex context"`,持久化到shell配置文件,后面的使用方式都是基于此alias的方式.
使用前建议使用`cx init`命令初始化保存现有的context信息,后续可以通过`cx restore`恢复.
#### 功能
通过help查看具体功能
```shell
cx -h
管理kubectl使用的context

Usage:
  crane context [flags]
  crane context [command]

Aliases:
  context, ctx

Available Commands:
  add         添加 [kubectl context] 资源
  delete      删除 [kubectl context] 资源
  info        查看指定 [kubectl context] 资源
  init        初始化 [kubectl context] 资源
  list        展示 [kubectl context] 资源
  prompt      获取当前 [kubectl context] 的prompt信息,用于配置shell提示
  restore     恢复 [kubectl context] 资源
  select      选择 [kubectl context] 资源
  set         设置 [kubectl context] 资源
  use         使用指定 [kubectl context] 资源

Flags:
  -h, --help   help for context

Use "crane context [command] --help" for more information about a command.
```
#### 添加`context`
1. ssh方式,此方式需要获取到主机密码
```shell
cx add root@127.0.0.1
```
2. acp方式,此方式需要获取到token
```shell
cx add --acp https://dev.me
√ input token: … ********
√ select context! … calicoca
"crane/app/kubectx.(*KubeCtx).AddMetadata" []interface {}{
  &kubectx.ContextMetadata{
    Host:      "192.168.176.113",
    Name:      "direct-connect",
    Namespace: "default",
    Path:      "/Users/aiden/.crane/kube/192.168.176.113",
    User:      "kubeconfig-user",
    Cluster:   "direct-calicoca",
  },
}
2023-06-16T11:57:16+08:00|INFO|kubectx.glob..func1| cmd/kubectx/add.go:36 |ok
```
#### 查看`context`
```shell
cx ls
+-------------------+----------------+-----------+-----------------+-----------------+
|       HOST        |      NAME      | NAMESPACE |     CLUSTER     |      USER       |
+-------------------+----------------+-----------+-----------------+-----------------+
|   192.168.132.183 | x86            | tsl-x     | x86             | admin           |
|   192.168.176.113 | direct-connect | default   | direct-calicoca | kubeconfig-user |
| * 192.168.18.130  | arm            | tsl-a     | arm             | admin           |
|   192.168.181.20  | global         | default   | global          | admin           |
+-------------------+----------------+-----------+-----------------+-----------------+
```
#### 选择`context`
使用`cx use`指定或`cx select`选择一个context
```shell
cx s
? select context! »
Filtering:
> 🔥 | 192.168.18.130 | arm | tsl-a | arm
     | 192.168.132.183 | x86 | tsl-x | x86
     | 192.168.176.113 | direct-connect | default | direct-calicoca
     | 192.168.181.20 | global | default | global
↑ move up • ↓ move down • tab/enter choice it • tab/enter finish selection • ^C kill program
```
#### 配置`prompt`
添加到shell配置文件中,例如`~/.zshrc`,可以添加参数修改具体样式
```shell
PROMPT='$(cx prompt)'$PROMPT
```
```shell
[⎈|arm:tsl-a]➜  workspace git:(master) ✗
```
#### 设置`context`
可以通过set设置当前context的默认命名空间与名称,也可以指定一个context设置其默认命名空间与名称
```shell
cx set -n tsl-a
```
#### 删除`context`
通过`cx delete`删除指定context,一般为host地址,也可以用name删除
```shell
cx delete 192.168.176.113
```
# 补充
kube context的管理功能在使用 https://github.com/ahmetb/kubectx 后根据个性化需求实现
prompt功能参考 https://github.com/jonmosco/kube-ps1 ,由于其不支持此项目切换context后更新提示,所以自己实现了一个
关于kubex部分优化后续参考 https://github.com/c-bata/kube-prompt 实现一个更好的支持交互式shell