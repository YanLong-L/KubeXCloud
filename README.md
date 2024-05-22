# KubeXClouds K8s运控平台

开发环境说明：

- go语言版本：go version go1.21.1 windows/amd64
- 编译环境：windows
- 开发工具：goland

## 项目的初始化

### web框架的选型

```bash
go get -u github.com/gin-gonic/gin@v1.8.1
```

### 把配置参数分离

```bash
go get github.com/spf13/viper@v1.13.0
```

> 参考文档：https://github.com/spf13/viper

### k8s集成

```bash
go get k8s.io/client-go@v0.20.4
```

## 项目接口开发

### kubeimooc 仪表盘功能 v1.11
- [x] 基础信息查看（k8s版本信息、集群初始化时间等）
- [x] 各资源的统计信息
- [x] 集群pod、cpu、内存耗用情况（瞬时）
    - [x] 安装metrics-server
    - [x] 调用metrics-server接口，计算集群的cpu和内存的耗用
- [x] 集群 cpu、内存变化趋势
    - [x] 安装prometheus
    - [x] 提供prometheus pull 数据的接口(exporter)
    - [x] 调用prometheus 查询指标统计数据

### kubeimooc 整合 Harbor v1.10
- [x] 集成HarborAPI
- [x] Projects 列表查询（分页，模糊查询）
- [x] Repositories 列表查询（分页，模糊查询）
- [x] Artifacts 列表查询（分页，模糊查询）
- [x] 镜像匹配的接口（用户Pod输入镜像信息的时候，自动匹配）

### k8s 认证与授权 v1.8.2

ServiceAccount
- [x] 创建
- [x] 删除
- [x] 查询（列表）

Role/ClusterRole
- [x] 创建/更新
- [x] 删除
- [x] 查询（详情/列表）

RoleBinding/ClusterRoleBinding
- [x] 创建/更新
- [x] 删除
- [x] 查询（详情/列表）

### k8s 认证与授权 v1.8.1
在集群内初始化，不需要指定.kube/config

### k8s工作负载 v1.7

StatefulSet
- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

Deployment
- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

DaemonSet
- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

Job
- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

CronJob
- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

### k8s服务发现 v1.6

Service

- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

Ingress

- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）

IngressRoute

- [x] 创建/更新
- [x] 删除
- [x] 查询（列表和详情）
- [x] Middleware的查询接口

### k8s卷管理接口 v1.5

PersistentVolume

- [x] 创建
- [x] 删除
- [x] 查询--列表

PersistentVolumeClaim

- [x] 创建
- [x] 删除
- [x] 查询--列表

StorageClass

- [x] 创建
- [x] 删除
- [x] 查询--列表

Pod管理

- [x] Pod管理（卷管理部分的逻辑）

优化点：

- [x] downward fileRefPath没有显示
- [x] PVC选择PV或Sc只能二选一
- [x] SC PVC PV 添加keyword搜素字段
- [x] PV显示StorageClassName

### 应用与配置分离接口 v1.4

ConfigMap

- [x] ConfigMap 新增/修改
- [x] 删除
- [x] 查询（列表和详情查询）

Secret

- [x] 新增/修改
- [x] 删除
- [x] 查询（列表和详情查询）

pod管理接口改动：

- [x] 新增ConfigMap和ConfigMepKey
- [x] 新增Secret和SecretKey

### NodeScheduling接口开发 1.3

- [x] node列表/详情(kubectl get nodes / kubectl describe node node-x)
- [x] node标签管理(kubectl label node node-x label-x=label-value-x)
    - 所有的标签上传
- [x] node污点(taint)管理
- [x] 查看node上所有的pod(kubectl get pod -n ns-x -o wide)

pod管理接口改动：

- [x] pod新增容忍(tolerations)参数
- [x] pod选择哪种方式调度：nodeName/nodeSelector/nodeAffinity

### Pod管理接口开发 1.2

- [x] 命名空间列表接口
- [x] Pod创建
- [x] Pod编辑（更新/升级）

---

- [x] Pod查看-详情 展示podrequest 数据 用于重新创建
- [x] Pod查看-列表
- [x] Pod删除

接口调优：

1. pod更新会多出一个卷挂载（ serviceAccount ）

- 计算哪些是emptydir volume mount 进行非emptydir过滤

2. 更新 pod 超时 -- pod删除等待时间不确定 改为强制删除
3. pod列表支持关键字搜索