package request

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Base struct {
	// pod 名字
	Name string `json:"name"`
	// 标签
	Labels []ListMapItem `json:"labels"`
	// 命名空间
	Namespace string `json:"namespace"`
	// 重启策略 Alway Never On-Failure
	RestartPolicy string `json:"restartPolicy"`
}

type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}

type Resources struct {
	// 是否配置容器的配额
	Enable bool `json:"enable"`
	// 内存 M
	MemRequest int32 `json:"memRequest"`
	MemLimit   int32 `json:"memLimit"`
	// cpu m
	CpuRequest int32 `json:"cpuRequest"`
	CpuLimit   int32 `json:"cpuLimit"`
}

type VolumeMount struct {
	//挂载卷名称
	MountName string `json:"mountName"`
	//挂载卷->对应的容器内的路径
	MountPath string `json:"mountPath"`
	//是否只读
	ReadOnly bool `json:"readOnly"`
}

type ProbeTime struct {
	//初始化时间 初始化若干秒之后才开始探针
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	//每隔若干秒之后 去探针
	PeriodSeconds int32 `json:"periodSeconds"`
	//探针等待时间 等待若干秒之后还没有返回 那么就是探测失败
	TimeoutSeconds int32 `json:"timeoutSeconds"`
	//探针若干次成功了 才认为这次探针成功
	SuccessThreshold int32 `json:"successThreshold"`
	//探测若干次 失败了 才认为这次探针失败
	FailureThreshold int32 `json:"failureThreshold"`
}

type ProbeHttpGet struct {
	//请求协议http / https
	Scheme string `json:"scheme"`
	//请求host 如果为空 那么就是Pod内请求
	Host string `json:"host"`
	//请求路径
	Path string `json:"path"`
	//请求端口
	Port int32 `json:"port"`
	//请求的header
	HttpHeaders []ListMapItem `json:"httpHeaders"`
}

type ProbeCommand struct {
	// cat /test/test.txt
	Command []string `json:"command"`
}

type ProbeTcpSocket struct {
	//请求host 如果为空 那么就是Pod内请求
	Host string `json:"host"`
	//探测端口
	Port int32 `json:"port"`
}

type ContainerProbe struct {
	// 是否打开探针
	Enable bool `json:"enable"`
	// 探针类型 tcp/http/exec
	HttpGet   ProbeHttpGet   `json:"httpGet"`
	Exec      ProbeCommand   `json:"exec"`
	TcpSocket ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}

type Container struct {
	// 容器名称
	Name string `json:"name"`
	// 容器镜像
	Image string `json:"image"`
	// 镜像拉取策略
	ImagePullPolicy string `json:"imagePullPolicy"`
	// 是否开启伪终端
	Tty bool `json:"tty"`
	// 容器端口
	Ports []ContainerPort `json:"ports"`
	// 工作目录
	WorkingDir string `json:"workingDir"`
	// 执行命令
	Command []string `json:"command"`
	// 参数
	Args []string `json:"args"`
	// 环境变量
	Envs []ListMapItem `json:"envs"`
	//是否开启模式
	Privileged bool `json:"privileged"`
	// 容器申请配额
	Resources Resources `json:"resources"`
	// 容器卷挂载
	VolumeMounts []VolumeMount `json:"volumeMounts"`
	// 启动探针
	StartupProbe ContainerProbe `json:"startupProbe"`
	// 存活探针
	LivenessProbe ContainerProbe `json:"livenessProbe"`
	// 就绪探针
	ReadinessProbe ContainerProbe `json:"readinessProbe"`
}

type Volume struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DnsConfig struct {
	Nameservers []string `json:"nameservers"`
}

type NetWorking struct {
	HostNetwork bool          `json:"hostNetwork"`
	HostName    string        `json:"hostName"`
	DnsPolicy   string        `json:"dnsPolicy"`
	DnsConfig   DnsConfig     `json:"dnsConfig"`
	HostAliases []ListMapItem `json:"hostAliases"`
}

type Pod struct {
	//基础定义信息
	Base Base `json:"base"`
	// 卷
	Volumes []Volume `json:"volumes"`
	//网络相关
	NetWorking NetWorking `json:"netWorking"`
	///init containers
	InitContainers []Container `json:"initContainers"`
	//containers
	Containers []Container `json:"containers"`
}

/*
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: my-container
    image: my-image
    ports:
    - containerPort: 80
    volumeMounts:
    - name: my-volume
      mountPath: /path/to/volume
    resources:
      limits:
        cpu: "1"
        memory: "1Gi"
      requests:
        cpu: "500m"
        memory: "500Mi"
    livenessProbe:
      exec:
        command:
        - cat
        - /path/to/healthy
      initialDelaySeconds: 30
      periodSeconds: 10
    readinessProbe:
      httpGet:
        path: /healthz
        port: 80
      initialDelaySeconds: 30
      periodSeconds: 10
  initContainers:
  - name: my-init-container
    image: my-init-image
    command:
    - sh
    - -c
    - echo "Initialization script executed" > /path/to/initialized
  volumes:
  - name: my-volume
    emptyDir: {}
  nodeSelector:
    disktype: ssd
  tolerations:
  - key: "node.kubernetes.io/not-ready"
    operator: "Exists"
    effect: "NoExecute"
    tolerationSeconds: 300
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disktype
            operator: In
            values:
            - ssd


这个 YAML 文件包含了以下 Kubernetes 特性：

定义了一个名为 my-pod 的 Pod。
Pod 中包含一个容器，使用了名为 my-image 的镜像。
容器暴露了端口 80。
容器挂载了一个名为 my-volume 的卷，挂载路径为 /path/to/volume。
为容器设置了资源限制和请求，包括 CPU 和内存。
定义了容器的存活探针（liveness probe）和就绪探针（readiness probe）。
添加了一个初始化容器（init container），使用名为 my-init-image 的镜像。
定义了一个名为 my-volume 的卷，使用 emptyDir 类型。
为 Pod 设置了节点选择器（node selector），要求节点的 disktype 标签为 ssd。
为 Pod 添加了容忍度（tolerations），允许在节点不可用时仍然调度到该节点。
为 Pod 设置了亲和性（affinity），要求在调度期间必须满足节点的 disktype 标签为 ssd
*/
